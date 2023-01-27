package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gym/ui"
)

func getConnection() (*sqlx.DB, error) {
	conn, err := sqlx.Open("postgres", "postgres://postgres:12345678@localhost:5432/exemplo_db?sslmode=disable")
	return conn, err
}

func main() {
	var choice int
	for {
		fmt.Print(ui.ShowMenuLogin())
		_, err := fmt.Scan(&choice)
		if err != nil {
			returnErr(err)
		}

		switch choice {
		case 1:
			RegisterInSystem()
			continue
		case 2:
			LoginInSystem()
			continue
		case 3:
			ShowRecords()
			continue
		default:
			nextLine()
			fmt.Println("Você não escolheu nenhuma opção acima!")
			nextLine()
			continue
		}
	}
}
