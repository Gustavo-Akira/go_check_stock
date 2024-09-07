package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"stocks/dto"
	"stocks/model"
	"stocks/repository"
	"strconv"
)

func VerifyAndSearchForStocks() {
	stocks, err := repository.FindAll()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if len(stocks) == 0 {
		fmt.Println("No stocks found")
		return
	}
	for _, stock := range stocks {
		_, err = repository.Find(stock.Name)
		if err != nil {
			fmt.Println("Http Request" + stock.Name)
			go getStockInformation(&stock)

		} else {
			fmt.Println("Stock up to date")
		}
	}
}

func getStockInformation(stock *model.Stock) {
	if stock == nil {
		fmt.Println("Stock not found")
		return
	}
	resp, err := http.Get("https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=" + stock.Name + "&interval=5min&apikey=U8XNX0F5VQTGE77V")
	if err != nil {
		fmt.Println("Error when updating stock " + stock.Name + " : " + err.Error())
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error when updating stock " + stock.Name + " : " + err.Error())
		return
	}
	var apiResponse dto.StockData
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Println("Error when updating stock " + stock.Name + " : " + err.Error())
		return
	}
	fmt.Println("Http request " + string(body))
	for _, data := range apiResponse.TimeSeries {
		stock.Price, err = strconv.ParseFloat(data.Close, 64)
		if err != nil {
			fmt.Println("Error when updating stock " + stock.Name + " : " + err.Error())
			return
		}
		repository.UpdateStockValue(stock)
		repository.Save(*stock)
		break
	}
	fmt.Print("End")
}
