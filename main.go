package main

import (
	controller "stocks/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/stock", controller.AddStock)
	r.Run()
}
