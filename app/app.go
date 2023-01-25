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
		helpTopicGroup := v1.Group("/helptopic")
		{
			helpTopicGroup.GET("", handlers.ListHelpTopicHandler)
		}
		departmentGroup := v1.Group("/department")
		{
			departmentGroup.GET("", handlers.ListDepartmentHandler)
		}
		stateGroup := v1.Group("/state")
		{
			stateGroup.GET("", handlers.ListStateHandler)
		}
		userGroup := v1.Group("/user")
		{
			userGroup.GET("", handlers.ListUserHandler)
			userGroup.GET("/:id", handlers.GetUserHandler)
		}
	}
	return app
}
