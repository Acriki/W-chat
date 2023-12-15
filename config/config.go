package config

import (
	"fmt"
	"os"
	"time"

	"W-chat/pkg/encrypt"

	"gopkg.in/yaml.v2"
)

// Config 配置信息
type Config struct {
	sid    string  // 服务运行ID
	MySQL  *MySQL  `json:"mysql" yaml:"mysql"`
	Server *Server `json:"server" yaml:"server"`
	Cors   *Cors   `json:"cors" yaml:"cors"`
	Redis  *Redis  `json:"redis" yaml:"redis"`
	Jwt    *Jwt    `json:"jwt" yaml:"jwt"`
}

type Server struct {
	Http      int `json:"http" yaml:"http"`
	Websocket int `json:"websocket" yaml:"websocket"`
	Tcp       int `json:"tcp" yaml:"tcp"`
}

// Cors 跨域配置
type Cors struct {
	Origin      string `json:"origin" yaml:"origin"`
	Headers     string `json:"headers" yaml:"headers"`
	Methods     string `json:"methods" yaml:"methods"`
	Credentials string `json:"credentials" yaml:"credentials"`
	MaxAge      string `json:"max_age" yaml:"max_age"`
}

func New() *Config {
	filename := "config.yaml"
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var conf Config
	if yaml.Unmarshal(content, &conf) != nil {
		panic(fmt.Sprintf("解析 config.yaml 读取错误: %v", err))
	}

	// 生成服务运行ID
	conf.sid = encrypt.Md5(fmt.Sprintf("%d", time.Now().UnixNano()))

	return &conf
}

// ServerId 服务运行ID
func (c *Config) ServerId() string {
	return c.sid
}
