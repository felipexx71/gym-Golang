package main

import (
	"fmt"
	"gym/models"
	"gym/ui"
)

var ListPerson []models.Person

var ListRecords []models.Records

func readList() {
	if len(ListPerson) != 0 {
		for i := range ListPerson {
			nextLine()
			fmt.Println("Nome(s): " + ListPerson[i].Name())
			fmt.Println("Pesos(s): " + ListPerson[i].Weight())
			fmt.Println("Altura(s): " + ListPerson[i].Height())
		}
	} else {
		nextLine()
		fmt.Print("Não há registros encontrados!")
	}
}

func readName() {
	if len(ListPerson) != 0 {
		for i := range ListPerson {
			fmt.Print("Nome(s): " + ListPerson[i].Name())
			nextLine()
		}
	} else {
		fmt.Print("Não há nomes cadastrados!")
	}

}

func readWeight() {
	if len(ListPerson) != 0 {
		for i := range ListPerson {
			fmt.Print("Pesos(s): " + ListPerson[i].Weight())
			nextLine()
		}
	} else {
		fmt.Print("Não há pesos cadastrados!")
	}
}

func readHeight() {
	if len(ListPerson) != 0 {
		for i := range ListPerson {
			fmt.Print("Altura(s): " + ListPerson[i].Height())
			nextLine()
		}
	} else {
		fmt.Print("Não há alturas cadastradas!")
	}
}

func findPersonByName() {
	var finder string

	if len(ListPerson) != 0 {
		for i := range ListPerson {
			fmt.Print("Digite o nome a ser buscado: ")
			fmt.Scan(&finder)
			if ListPerson[i].Name() == finder {
				nextLine()
				fmt.Print("o nome " + finder + " foi achado!")
			} else {
				fmt.Print("Não é possível buscar o nome: " + finder + "\n")
			}
		}
	} else {
		fmt.Print("Não há cadastros encontrados!")
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
			println("Opção inválida!")
			nextLine()
			continue
		}
	}
}

func LoginInSystem() {
	var user string
	var pass string

	if len(ListRecords) == 0 {
		nextLine()
		fmt.Println("Não há cadastros no sistema!")
	} else if len(ListRecords) == 1 {
		for i := 0; i < 5; i++ {
			nextLine()
			fmt.Print("Digite o seu email: ")
			fmt.Scan(&user)
			fmt.Print("Digite a sua senha: ")
			fmt.Scan(&pass)

			if ListRecords[0].Email() == user && ListRecords[0].Password() == pass {
				SwitchOption()
			} else {
				nextLine()
				fmt.Println("Não há cadastros com esse usuário e senha! Aperte 1 para poder registrar")
				return
			}
		}
	} else if len(ListRecords) == 2 {
		for i := 0; i < 5; i++ {
			nextLine()
			fmt.Print("Digite o seu email: ")
			fmt.Scan(&user)
			fmt.Print("Digite a sua senha: ")
			fmt.Scan(&pass)

			if ListRecords[0].Email() == user && ListRecords[0].Password() == pass {
				SwitchOption()
			} else if ListRecords[1].Email() == user && ListRecords[1].Password() == pass {
				SwitchOption()
			} else {
				nextLine()
				fmt.Println("Não há cadastros com esse usuário e senha! Aperte 1 para poder registrar")
				return
			}
		}
	} else if len(ListRecords) == 3 {
		for i := 0; i < 5; i++ {
			nextLine()
			fmt.Print("Digite o seu email: ")
			fmt.Scan(&user)
			fmt.Print("Digite a sua senha: ")
			fmt.Scan(&pass)

			if ListRecords[0].Email() == user && ListRecords[0].Password() == pass {
				SwitchOption()
			} else if ListRecords[1].Email() == user && ListRecords[1].Password() == pass {
				SwitchOption()
			} else if ListRecords[2].Email() == user && ListRecords[2].Password() == pass {
				SwitchOption()
			} else {
				nextLine()
				fmt.Println("Não há cadastros com esse usuário e senha! Aperte 1 para poder registrar")
				return
			}
		}
	} else if len(ListRecords) == 4 {
		for i := 0; i < 5; i++ {
			nextLine()
			fmt.Print("Digite o seu email: ")
			fmt.Scan(&user)
			fmt.Print("Digite a sua senha: ")
			fmt.Scan(&pass)

			if ListRecords[0].Email() == user && ListRecords[0].Password() == pass {
				SwitchOption()
			} else if ListRecords[1].Email() == user && ListRecords[1].Password() == pass {
				SwitchOption()
			} else if ListRecords[2].Email() == user && ListRecords[2].Password() == pass {
				SwitchOption()
			} else if ListRecords[3].Email() == user && ListRecords[3].Password() == pass {
				SwitchOption()
			} else {
				nextLine()
				fmt.Println("Não há cadastros com esse usuário e senha! Aperte 1 para poder registrar")
				return
			}
		}
	} else if len(ListRecords) == 5 {
		for i := 0; i < 5; i++ {
			nextLine()
			fmt.Print("Digite o seu email: ")
			fmt.Scan(&user)
			fmt.Print("Digite a sua senha: ")
			fmt.Scan(&pass)

			if ListRecords[0].Email() == user && ListRecords[0].Password() == pass {
				SwitchOption()
			} else if ListRecords[1].Email() == user && ListRecords[1].Password() == pass {
				SwitchOption()
			} else if ListRecords[2].Email() == user && ListRecords[2].Password() == pass {
				SwitchOption()
			} else if ListRecords[3].Email() == user && ListRecords[3].Password() == pass {
				SwitchOption()
			} else if ListRecords[4].Email() == user && ListRecords[4].Password() == pass {
				SwitchOption()
			} else {
				nextLine()
				fmt.Println("Não há cadastros com esse usuário e senha! Aperte 1 para poder registrar")
				return
			}
		}
	}
}

func RegisterInSystem() {
	record := models.Records{}

	var email string
	var password string

	if len(ListRecords) < 5 {
		fmt.Print("Digite o seu email: ")
		fmt.Scan(&email)
		record.SetEmail(email)
		fmt.Print("Digite a sua senha: ")
		fmt.Scan(&password)
		record.SetPassword(password)

		ListRecords = append(ListRecords, record)
	} else {
		nextLine()
		fmt.Println("A quantidade de usuários foi excedida!")
	}
}

func ShowRecords() {

	if len(ListRecords) != 0 {
		for i := range ListRecords {
			fmt.Println("Email: " + ListRecords[i].Email())
			nextLine()
		}
	} else {
		nextLine()
		fmt.Println("A lista está vazia!")
	}
}
