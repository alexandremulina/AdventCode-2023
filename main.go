package main

import (
	"go-advent/api"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	r := gin.Default()
	r.GET("/", api.HealthCheck)
	r.POST("/dayOne", api.DayOne)
	r.POST("/dayTwo", api.DayTwo)
	r.POST("/dayThree", api.DayThreePart2)
	r.POST("/dayFour", api.DayFour)
	r.POST("/dayFive", api.DayFive)
	r.POST("/daySix", api.DaySix)

	r.Run(":9002")
}
