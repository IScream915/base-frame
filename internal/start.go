package internal

import (
	"base_frame/internal/middleware"
	"base_frame/internal/repo/models"
	"base_frame/pkg/common/config"
	"base_frame/pkg/db/mysqlutil"
	"base_frame/pkg/db/redisutil"
	"base_frame/pkg/validation"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// NewGinEngine 构造一个新的Gin引擎，用于生成一个新的经过初始化和配置的Gin引擎
func NewGinEngine() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	// 启用自定义的引擎
	engine := gin.New()
	// 注册全局中间件
	// gin.Recovery是内置的回复中间件，用于捕获程序中的panic，防止服务崩溃，并返回500错误
	engine.Use(gin.Recovery(), middleware.RequestID(), middleware.GinLogger())
	// 设置自定义参数验证器
	// TODO 弄清楚这个设置自定义参数验证器的作用
	// 注册自定义参数验证器
	validation.RegisterCustomValidation()
	// 设置路由
	SetRoute()
	return engine

}

func SetRoute() {

}

func Start(ctx context.Context, config *config.Config) error {
	app := fx.New(
		injectConfig(ctx, config),
		injectComponent(),
		injectRepo(),
		fx.Invoke(func(db *gorm.DB) {
			//CREATE DATABASE IF NOT EXISTS {{DB}};
			_ = db.AutoMigrate(models.User{})
		}),
		injectService(),
		injectApi(),
		fx.Provide(NewGinEngine),
		fx.Invoke(func(lifecycle fx.Lifecycle, engine *gin.Engine, db *gorm.DB, rdb redis.UniversalClient) {
			lifecycle.Append(
				fx.Hook{
					OnStart: func(ctx context.Context) error {
						go func() {
							if err := engine.Run(fmt.Sprintf(":%d", config.Port)); err != nil {
								panic(err)
							}
						}()
						return nil
					},
					OnStop: func(ctx context.Context) error {
						// 在这里添加关闭逻辑，例如关闭数据库连接等
						return nil
					},
				},
			)
		}),
	)
	app.Run()
	fmt.Println("1111")
	return nil
}

// 注入配置
func injectConfig(ctx context.Context, cfg *config.Config) fx.Option {
	return fx.Provide(
		func() *config.Config { return cfg },
		func() *mysqlutil.Config { return cfg.Mysql },
		func() *redisutil.Config { return cfg.Redis },
		func() context.Context { return ctx },
	)
}

// 注入公共组件实例
func injectComponent() fx.Option {
	return fx.Provide(
		// Mysql客户端
		mysqlutil.NewMysqlClient,
		// Redis客户端
		redisutil.NewRedisClient,
	)
}

// 注入仓储实现
func injectRepo() fx.Option {
	return fx.Provide()
}

// 注入service
func injectService() fx.Option {
	return fx.Provide()
}

// 注入API
func injectApi() fx.Option {
	return fx.Provide()
}
