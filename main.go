package main

import (
	"github.com/gin-gonic/gin"
	routers "MA.Content.Services.OrgMapper/src/routers"
)

func main() {
	r := gin.Default()
	orgMapper := r.Group("orgmapper")
	routers.InitRoutes(orgMapper);
	r.Run(":9092")
}