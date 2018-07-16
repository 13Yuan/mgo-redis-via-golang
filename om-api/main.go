package main

import (
	"github.com/gin-gonic/gin"
	routers "MA.Content.Services.OrgMapper/src/routers"
)

func main() {
	r := gin.Default()
	orgMapper := r.Group("organization")
	routers.InitRoutes(orgMapper);
	r.Run(":9093")
}