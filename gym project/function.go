package main

import (
	"fmt"
	"gym/models"
	"gym/ui"
)

var ListPerson []models.Person

var ListRecords []models.Records

func readList() {
	for i := range ListPerson {
		nextLine()
		fmt.Println("Nome(s): " + ListPerson[i].Name())
		fmt.Println("Pesos(s): " + ListPerson[i].Weight())
		fmt.Println("Altura(s): " + ListPerson[i].Height())
	}
}

func readName() {
	for i := range ListPerson {
		fmt.Println("Nome(s): " + ListPerson[i].Name())
	}
}

func readWeight() {
	for i := range ListPerson {
		fmt.Println("Pesos(s): " + ListPerson[i].Weight())
	}
}

func readHeight() {
	for i := range ListPerson {
		fmt.Println("Altura(s): " + ListPerson[i].Height())
	}
}

func findPersonByName() {
	var finder string

	for i := range ListPerson {
		fmt.Print("Digite o nome a ser buscado: ")
		fmt.Scan(&finder)
		if ListPerson[i].Name() == finder {
			fmt.Print("o nome " + finder + " foi achado!")
		} else {
			fmt.Print("Não é possivel buscar o nome: " + finder)
		}
	}
}

func nextLine() {
	fmt.Println("")
}

func registerList() {
	person := models.Person{}

	var name string
	var weight string
	var height string

	fmt.Println("Digite um nome:")
	fmt.Scan(&name)
	person.SetName(name)
	fmt.Println("Digite um peso:")
	fmt.Scan(&weight)
	person.SetWeight(weight)
	fmt.Println("Digite uma altura:")
	fmt.Scan(&height)
	person.SetHeight(height)

	ListPerson = append(ListPerson, person)
}

func SwitchOption() {
	var opc int

	for {
		fmt.Print(ui.ShowLogo())
		fmt.Print(ui.ShowMenuFunctions())
		fmt.Scan(&opc)
		switch opc {
		case 1:
			nextLine()
			registerList()
			nextLine()
			continue
		case 2:
			readList()
			nextLine()
			continue
		case 3:
			nextLine()
			readName()
			nextLine()
			continue
		case 4:
			nextLine()
			readWeight()
			nextLine()
			continue
		case 5:
			nextLine()
			readHeight()
			nextLine()
			continue
		case 6:
			nextLine()
			findPersonByName()
			nextLine()
			continue
		default:
			nextLine()
			println("Opção inválida")
			continue
		}
	}
}

func LoginInSystem() {
	var user string
	var pass string

	for i := 0; i < 5; i++ {
		nextLine()
		fmt.Print("Digite o seu email: ")
		fmt.Scan(&user)
		fmt.Print("Digite a sua senha: ")
		fmt.Scan(&pass)
		if ListRecords[0].Email() == user && ListRecords[0].Password() == pass || ListRecords[1].Email() == user && ListRecords[1].Password() == pass || ListRecords[2].Email() == user && ListRecords[2].Password() == pass || ListRecords[3].Email() == user && ListRecords[3].Password() == pass || ListRecords[4].Email() == user && ListRecords[4].Password() == pass {
			SwitchOption()
		} else {
			fmt.Print("Não foi possível buscar esse cadastro. Tente novamente!\n")
		}
	}
}

func RegisterInSystem() {
	record := models.Records{}

	var email string
	var password string

	fmt.Print("Digite o seu email: ")
	fmt.Scan(&email)
	record.SetEmail(email)
	fmt.Print("Digite a sua senha: ")
	fmt.Scan(&password)
	record.SetPassword(password)

	ListRecords = append(ListRecords, record)
}

func ShowRecords() {
	for i := range ListRecords {
		nextLine()
		fmt.Println("Email: " + ListRecords[i].Email())
	}
}
