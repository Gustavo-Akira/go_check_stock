package model

import (
	exception "stocks/exception"
	"time"
)

type Stock struct {
	Price       float64   `json:"price"`
	Name        string    `json:"name"`
	TargetPrice float64   `json:"target_price"`
	LastUpdate  time.Time `json:"last_update"`
	Links       []string  `json:"links"`
}

func MakeStock(price float64, name string, targetPrice float64, links []string) (Stock, error) {
	listOfError := []string{}
	if price == 0.0 {
		listOfError = append(listOfError, "Price cannot be 0.0")
	}
	if len(name) == 0 {
		listOfError = append(listOfError, "Name cannot be empty")
	}

	if targetPrice == 0.0 {
		listOfError = append(listOfError, "Target price cannot be 0.0")
	}

	if len(links) < 1 {
		listOfError = append(listOfError, "It have to has 1 link")
	}
	if len(listOfError) > 0 {
		validationError := exception.ValidationError{
			ListOfError: listOfError,
		}

		return Stock{}, validationError
	}

	return Stock{
		Price:       price,
		Name:        name,
		TargetPrice: targetPrice,
		LastUpdate:  time.Now(),
		Links:       links,
	}, nil
}

func (Stock *Stock) updateTargetPrice(targetPrice float64) error {
	if targetPrice == 0.0 {
		return exception.ValidationError{ListOfError: []string{"Target price cannot be 0.0"}}
	}
	Stock.TargetPrice = targetPrice
	return nil
}
