package internal

import (
	"base_frame/internal/middleware"
	"base_frame/pkg/config"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
)

// NewGinEngine 构造一个新的Gin引擎，用于生成一个新的经过初始化和配置的Gin引擎
func NewGinEngine() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	// 启用自定义的引擎
	engine := gin.New()
	// 注册全局中间件
	// gin.Recovery是内置的回复中间件，用于捕获程序中的panic，防止服务崩溃，并返回500错误
	engine.Use(gin.Recovery(), middleware.RequestID(), middleware.GinLogger())
	return engine
}

func Start(ctx context.Context, config *config.Config) error {
	fmt.Println("1111")
	return nil
}
