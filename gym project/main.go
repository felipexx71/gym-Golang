package main

import (
	"fmt"
	"gym/ui"
)

func main() {
	var choice int
	for {
		fmt.Print(ui.ShowMenuLogin())
		fmt.Scan(&choice)
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
