package router

import (
	mygin "W-chat/src/gin"
	"W-chat/src/httpserver/api/handler/web"

	"github.com/gin-gonic/gin"
)

// RegisterWebRoute 注册 Web 路由
func RegisterWebRoute(router *gin.Engine, handler *web.Handler) {

	// 授权验证中间件
	// authorize := middleware.Auth(secret, "api", session)

	// v1 接口
	v1 := router.Group("/api/v1")
	{
		// 授权相关分组
		auth := v1.Group("/auth")
		{
			auth.POST("/login", mygin.HandlerFunc(handler.V1.Auth.Login))                // 登录
			// auth.POST("/register", mygin.HandlerFunc(handler.V1.Auth.Register))          // 注册
			// auth.POST("/refresh", authorize, ichat.HandlerFunc(handler.V1.Auth.Refresh)) // 刷新 Token
			// auth.POST("/logout", authorize, ichat.HandlerFunc(handler.V1.Auth.Logout))   // 退出登录
			// auth.POST("/forget", mygin.HandlerFunc(handler.V1.Auth.Forget))              // 找回密码
		}

	}

}
