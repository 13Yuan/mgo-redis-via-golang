package routes

import (
	handler "MA.Content.Services.OrgMapper/src/handlers"
	"github.com/gin-gonic/gin"
)

/*InitRoutes initial the routes for rating*/
func InitRoutes(router *gin.RouterGroup) {
	router.GET("identifiers/:type/:id", handler.GetIdentifiers)
	router.GET("identifier", handler.Test)
}
