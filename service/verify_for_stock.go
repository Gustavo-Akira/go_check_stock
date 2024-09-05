package service

import (
	"fmt"
	"stocks/repository"
)

func VerifyAndSearchForStocks() {
	stocks, err := repository.FindAll()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for _, stock := range stocks {
			stock, err = repository.Find(stock.Name)
			if err != nil {
				fmt.Println("Http Request")
			} else {
				fmt.Println("Stock up to date")
			}
		}
	}

}
