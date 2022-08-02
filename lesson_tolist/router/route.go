package router

import (
	"example.com/to_list/api"
	"example.com/to_list/conf"
	"example.com/to_list/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)



func NewRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	router:=gin.Default()
	store:=conf.NewStore()
	//middleware
	router.Use(sessions.Sessions("Access",store))
	router.GET("/ping", api.PingService)
	prefix:=router.Group("/api/v1")
	{
		prefix.POST("/signup",api.SignUp)
		prefix.POST("/login",api.Login)
		taskGroup:=prefix.Group("/")
		taskGroup.Use(middleware.JwtMiddleware())
		{
			taskGroup.GET("tasks",api.ListTask)
			taskGroup.POST("task",api.CreateTask)
			taskGroup.GET("show/:tid",api.ShowTask)
			taskGroup.GET("delete/:tid",api.DeleteTask)
			taskGroup.POST("update/:tid",api.UpdateTask)
		}
	}
	return router
}