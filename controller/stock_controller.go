package controller

import (
	m "stocks/model"
	"strconv"

	"stocks/repository"

	"github.com/gin-gonic/gin"
)

func AddStock(c *gin.Context) {
	var stock m.Stock
	e := c.ShouldBind(&stock)
	if e == nil {

		stockStruct, erro := m.MakeStock(stock.Price, stock.Name, stock.TargetPrice, stock.Links)
		if erro == nil {
			_, e := repository.AddStock(stockStruct)
			if e == nil {
				c.JSON(200, stockStruct)
			} else {
				c.JSON(400, gin.H{
					"message": e.Error(),
				})
			}

		} else {

			c.JSON(400, gin.H{
				"message": erro.Error(),
			})
		}
	} else {
		c.JSON(500, gin.H{
			"message": e.Error(),
		})
	}
}

func RetrieveStock(c *gin.Context) {

	key, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "key cannot be null",
		})
		return
	}
	stock, err := repository.FindById(key)
	if err != nil {
		c.JSON(404, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(200, stock)
}

func RetrieveAllStocks(c *gin.Context) {
	stocks, err := repository.FindAll()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	}

	if stocks == nil {
		stocks = []m.Stock{}
	}

	c.JSON(200, stocks)
}
