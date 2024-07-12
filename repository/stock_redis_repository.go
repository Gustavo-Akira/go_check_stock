package repository

import (
	"context"
	"encoding/json"
	m "stocks/model"
	"time"

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
	result := client.Set(cxt, stock.Name, p, 20*time.Minute)
	if result.Err() != nil {
		return "", result.Err()
	}
	return result.Result()
}

func Find(key string) (m.Stock, error) {
	var stock m.Stock
	cxt := context.Background()
	result := client.Get(cxt, key)
	if result.Err() != nil {
		return stock, result.Err()
	}
	resultByte, err := result.Bytes()
	if err != nil {
		return stock, err
	}

	e := json.Unmarshal(resultByte, &stock)

	return stock, e
}
