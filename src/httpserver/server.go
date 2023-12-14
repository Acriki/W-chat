package httpserver

import (
	"W-chat/config"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

type Basic struct {
	Config *config.Config
	Gin    *gin.Engine
}

func Run(info *Basic) error {
	// if !server.Config.Debug() {
	// 	gin.SetMode(gin.ReleaseMode)
	// }

	ctx := context.Background()
	eg, groupCtx := errgroup.WithContext(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	log.Printf("HTTP Listen Port :%d", info.Config.Server.Http)
	log.Printf("HTTP Server Pid  :%d", os.Getpid())

	return run(c, eg, groupCtx, info)
}

func run(c chan os.Signal, eg *errgroup.Group, ctx context.Context, info *Basic) error {

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", info.Config.Server.Http),
		Handler: info.Gin,
	}

	// 启动 http 服务
	eg.Go(func() error {
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}

		return nil
	})

	eg.Go(func() error {
		defer func() {
			log.Println("Shutting down server...")

			// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
			timeCtx, timeCancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer timeCancel()

			if err := server.Shutdown(timeCtx); err != nil {
				log.Fatalf("HTTP Server Shutdown Err: %s", err)
			}
		}()

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-c:
			return nil
		}
	})

	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		log.Fatalf("HTTP Server forced to shutdown: %s", err)
	}

	log.Println("Server exiting")

	return nil
}
