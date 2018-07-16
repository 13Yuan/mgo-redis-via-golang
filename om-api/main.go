package main

import (
	routers "om-api/src/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	orgMapper := r.Group("organization")
	routers.InitRoutes(orgMapper);
	r.Run(":9093")
}