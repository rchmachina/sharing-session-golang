package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rchmachina/sharing-session-golang/handler"
	"github.com/rchmachina/sharing-session-golang/repositories"
	"github.com/rchmachina/sharing-session-golang/utils/database"
	"github.com/rchmachina/sharing-session-golang/utils/middleware"
)

func UserRoutes(r *gin.RouterGroup) {
	userRepository := repositories.RepositoryUser(database.DB)
	h := handlers.HandlerUser(userRepository)

	r.POST("/user", h.CreateUser)
	r.DELETE("/user/:id", middleware.Auth(h.DeleteUser))
	r.POST("/loginUser", h.LoginUser)
	r.PUT("/user", middleware.Auth(h.UpdateUser))
	r.GET("/user", middleware.Auth(h.GetAllUser))
	// e.GET("/user")
	// e.POST("/deleteUser/:id", middleware.Auth(h.FindUsersPeer))
}
