package routes

import (
	handler "MA.Content.Services.OrgMapper/src/handlers"
	"github.com/gin-gonic/gin"
)

/*InitRoutes initial the routes for rating*/
func InitRoutes(router *gin.RouterGroup) {
	router.GET("/:type/:id", handler.GetIdentifiers)
}
