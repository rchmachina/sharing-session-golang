package routes

import (
	"github.com/gin-gonic/gin"
)

func RouteInit(r *gin.RouterGroup) {
	UserRoutes(r)
}
