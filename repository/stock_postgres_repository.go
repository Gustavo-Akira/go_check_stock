package repository

import (
	"database/sql"
	"fmt"
	"stocks/model"
	"strings"

	"github.com/lib/pq"
)

func openConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=postgres port=5432 user=postgres password=postgres dbname=stock sslmode=disable")
	if err != nil {
		panic(err)
	}

	err = db.Ping()

	return db, err
}

func AddStock(stock model.Stock) (id int64, err error) {
	conn, err := openConnection()
	if err != nil {
		return id, err
	}

	query := `INSERT INTO stock(price, name, target_price, links) VALUES($1,$2,$3,$4) RETURNING id`
	err = conn.QueryRow(query, stock.Price, stock.Name, stock.TargetPrice, "{"+strings.Join(stock.Links, ";")+"}").Scan(&id)
	return id, err
}

func FindById(id int) (stock model.Stock, e error) {
	conn, err := openConnection()
	if err != nil {
		return stock, err
	}

	query := `SELECT id, price, name, target_price, links FROM stock WHERE id=$1`

	e = conn.QueryRow(query, id).Scan(&stock.Id, &stock.Price, &stock.Name, &stock.TargetPrice, pq.Array(&stock.Links))
	return stock, e
}

func FindAll() (stocks []model.Stock, e error) {
	conn, err := openConnection()
	var rows *sql.Rows
	if err != nil {
		return stocks, err
	}

	query := `SELECT id, price, name, target_price, links FROM stock`
	rows, e = conn.Query(query)
	var stock model.Stock
	for rows.Next() {
		var x []uint8
		e = rows.Scan(&stock.Id, &stock.Price, &stock.Name, &stock.TargetPrice, pq.Array(&stock.Links))
		if e != nil {
			break
		}
		fmt.Print(x)
		stocks = append(stocks, stock)
	}

	return stocks, e
}
