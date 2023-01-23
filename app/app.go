package app

import (
	"github.com/andres15alvarez/ticket_manager/config"
	"github.com/andres15alvarez/ticket_manager/handlers"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type App struct {
	DB     *bun.DB
	Server *gin.Engine
}

func InitializeServer() *gin.Engine {
	return gin.Default()
}

func CreateApp() *App {
	conf := config.InitializeConfig()
	conf.InitializeDB()
	app := &App{
		DB:     config.GetDB(),
		Server: InitializeServer(),
	}
	v1 := app.Server.Group("/api")
	{
		homeGroup := v1.Group("/")
		{
			homeGroup.GET("", handlers.HomeHandler)
		}
		userTypeGroup := v1.Group("/usertype")
		{
			userTypeGroup.GET("", handlers.ListUserTypeHandler)
		}
	}
	return app
}
