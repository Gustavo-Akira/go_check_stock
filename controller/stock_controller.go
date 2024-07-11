package controller

import (
	m "stocks/model"

	"stocks/repository"

	"github.com/gin-gonic/gin"
)

func AddStock(c *gin.Context) {
	var stock m.Stock
	e := c.ShouldBind(&stock)
	if e == nil {

		stockStruct, erro := m.MakeStock(stock.Price, stock.Name, stock.TargetPrice, stock.Links)
		if erro == nil {
			result, e := repository.Save(stockStruct)
			if result == "" {
				c.JSON(200, stockStruct)
			} else {
				c.JSON(400, e)
			}

		} else {

			c.JSON(400, erro)
		}
	} else {
		c.JSON(500, e.Error())
	}
}
