package controller

import (
	m "stocks/model"

	"github.com/gin-gonic/gin"
)

func AddStock(c *gin.Context) {
	var stock m.Stock
	e := c.ShouldBind(&stock)
	if e == nil {

		stockStruct, erro := m.MakeStock(stock.Price, stock.Name, stock.TargetPrice, stock.Links)
		if erro == nil {
			c.JSON(200, stockStruct)
		} else {

			c.JSON(400, erro)
		}
	} else {
		c.JSON(500, e.Error())
	}

}
