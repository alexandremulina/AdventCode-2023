package main

import (
	"go-advent/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", api.HealthCheck)
	r.POST("/dayOne", api.DayOne)

	r.Run(":9009")
}
