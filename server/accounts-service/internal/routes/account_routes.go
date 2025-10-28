package routes

import (
	"accounts-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterAccountRoutes(router *gin.Engine, accountHandler *handlers.AccountHandler) {
	v1 := router.Group("/api/v1/accounts")
	{
		v1.POST("/", accountHandler.Create)
		v1.GET("/", accountHandler.GetAll)
		v1.GET("/:id", accountHandler.GetByID)
		v1.PUT("/:id", accountHandler.Update)
		v1.DELETE("/:id", accountHandler.Delete)
	}
}
