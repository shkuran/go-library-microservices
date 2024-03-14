package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shkuran/go-library-microservices/user-service/user"
)

func RegisterRoutes(server *gin.Engine, user user.Handler) {
	server.POST("/signup", user.CreateUser)
	server.POST("/login", user.Login)
	server.GET("/users", user.GetUsers)
}
