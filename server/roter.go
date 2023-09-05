package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kar-pev/task-checker/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", controllers.Ping)

	return router
}
