package main

import (
	controller "stocks/controller"
	"stocks/job"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/stock", controller.AddStock)
	r.GET("/stock", controller.RetrieveStock)
	r.PUT("/stock/:id", controller.AddStock)
	r.GET("/stocks", controller.RetrieveAllStocks)
	job.CheckAgain()
	r.Run()
}
