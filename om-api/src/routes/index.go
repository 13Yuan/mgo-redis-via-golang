package routes

import (
	handler "om-api/src/handlers"
	"github.com/gin-gonic/gin"
)

/*InitRoutes initial the routes for rating*/
func InitRoutes(router *gin.RouterGroup) {
	router.GET("/:type/:id", handler.GetIdentifiers)
}
