package repository

import (
	"context"
	"encoding/json"
	m "stocks/model"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func Save(stock m.Stock) (string, error) {
	cxt := context.Background()
	p, err := json.Marshal(stock)
	if err != nil {
		return "", err
	}
	result := client.Set(cxt, stock.Name, p, 20000)
	return result.Result()
}
