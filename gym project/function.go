package main

import (
	"fmt"
	"gym/models"
	"gym/ui"
	"log"
)

var ListRecords []models.Records

func readList() []models.Person {
	conn, err := getConnection()
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	var person []models.Person

	err = conn.Select(&person, "select name, weight, height from person")
	if err != nil {
		log.Fatal(err)
	}
	return person
}

func showList() {
	persons := readList()

	for _, person := range persons {
		nextLine()
		fmt.Println("##########")
		fmt.Println("Nome: ", person.GetName())
		fmt.Println("Peso: ", person.GetWeight())
		fmt.Println("Altura: ", person.GetHeight())
		fmt.Println("##########")
	}
}

func readName() {
	persons := readList()

	for _, person := range persons {
		nextLine()
		fmt.Println("##########")
		fmt.Println("Nome: ", person.GetName())
		fmt.Println("##########")
	}
}

func readWeight() {
	persons := readList()

	for _, person := range persons {
		nextLine()
		fmt.Println("##########")
		fmt.Println("Peso: ", person.GetWeight())
		fmt.Println("##########")
	}
}

func readHeight() {
	persons := readList()

	for _, person := range persons {
		nextLine()
		fmt.Println("##########")
		fmt.Println("Altura: ", person.GetHeight())
		fmt.Println("##########")
	}
}

func findPersonByName() {
	persons := readList()

	var finder string

	fmt.Print("Digite o nome a ser buscado: ")

	_, err := fmt.Scan(&finder)
	if err != nil {
		returnErr(err)
	}

	for _, person := range persons {
		if finder == person.GetName() {
			nextLine()
			fmt.Print("O nome: ", finder, " foi achado no banco de dados!")
			return
		} else {
			continue
		}
	}
	nextLine()
	fmt.Print("O nome pesquisado não foi encontrado!")
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
	_, err := fmt.Scan(&name)
	if err != nil {
		returnErr(err)
	}
	person.SetName(name)

	fmt.Println("Digite um peso:")
	_, err = fmt.Scan(&weight)
	if err != nil {
		returnErr(err)
	}
	person.SetWeight(weight)

	fmt.Println("Digite uma altura:")
	_, err = fmt.Scan(&height)
	if err != nil {
		returnErr(err)
	}
	person.SetHeight(height)

	conn, err := getConnection()
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	insertInto := `insert into person (name, weight, height) values ($1, $2, $3)`

	_, err = conn.Exec(insertInto, person.GetName(), person.GetWeight(), person.GetHeight())
	if err != nil {
		log.Fatal(err)
	}
}

func SwitchOption() {
	var opc int

	for {
		fmt.Print(ui.ShowLogo())
		fmt.Print(ui.ShowMenuFunctions())

		_, err := fmt.Scan(&opc)
		if err != nil {
			returnErr(err)
		}

		switch opc {
		case 1:
			nextLine()
			registerList()
			nextLine()
			continue
		case 2:
			showList()
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
			_, err := fmt.Scan(&user)
			if err != nil {
				returnErr(err)
			}

			fmt.Print("Digite a sua senha: ")
			_, err = fmt.Scan(&pass)
			if err != nil {
				returnErr(err)
			}

			if ListRecords[0].GetEmail() == user && ListRecords[0].GetPassword() == pass {
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
			_, err := fmt.Scan(&user)
			if err != nil {
				returnErr(err)
			}

			fmt.Print("Digite a sua senha: ")
			_, err = fmt.Scan(&pass)
			if err != nil {
				returnErr(err)
			}

			if ListRecords[0].GetEmail() == user && ListRecords[0].GetPassword() == pass {
				SwitchOption()
			} else if ListRecords[1].GetEmail() == user && ListRecords[1].GetPassword() == pass {
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
			_, err := fmt.Scan(&user)
			if err != nil {
				returnErr(err)
			}

			fmt.Print("Digite a sua senha: ")
			_, err = fmt.Scan(&pass)
			if err != nil {
				returnErr(err)
			}

			if ListRecords[0].GetEmail() == user && ListRecords[0].GetPassword() == pass {
				SwitchOption()
			} else if ListRecords[1].GetEmail() == user && ListRecords[1].GetPassword() == pass {
				SwitchOption()
			} else if ListRecords[2].GetEmail() == user && ListRecords[2].GetPassword() == pass {
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
			_, err := fmt.Scan(&user)
			if err != nil {
				returnErr(err)
			}

			fmt.Print("Digite a sua senha: ")
			_, err = fmt.Scan(&pass)
			if err != nil {
				returnErr(err)
			}

			if ListRecords[0].GetEmail() == user && ListRecords[0].GetPassword() == pass {
				SwitchOption()
			} else if ListRecords[1].GetEmail() == user && ListRecords[1].GetPassword() == pass {
				SwitchOption()
			} else if ListRecords[2].GetEmail() == user && ListRecords[2].GetPassword() == pass {
				SwitchOption()
			} else if ListRecords[3].GetEmail() == user && ListRecords[3].GetPassword() == pass {
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
			_, err := fmt.Scan(&user)
			if err != nil {
				returnErr(err)
			}

			fmt.Print("Digite a sua senha: ")
			_, err = fmt.Scan(&pass)
			if err != nil {
				returnErr(err)
			}

			if ListRecords[0].GetEmail() == user && ListRecords[0].GetPassword() == pass {
				SwitchOption()
			} else if ListRecords[1].GetEmail() == user && ListRecords[1].GetPassword() == pass {
				SwitchOption()
			} else if ListRecords[2].GetEmail() == user && ListRecords[2].GetPassword() == pass {
				SwitchOption()
			} else if ListRecords[3].GetEmail() == user && ListRecords[3].GetPassword() == pass {
				SwitchOption()
			} else if ListRecords[4].GetEmail() == user && ListRecords[4].GetPassword() == pass {
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
		_, err := fmt.Scan(&email)
		if err != nil {
			returnErr(err)
		}
		record.SetEmail(email)

		fmt.Print("Digite a sua senha: ")
		_, err = fmt.Scan(&password)
		if err != nil {
			returnErr(err)
		}
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
			fmt.Println("Email: " + ListRecords[i].GetEmail())
			nextLine()
		}
	} else {
		nextLine()
		fmt.Println("A lista está vazia!")
	}
}

func returnErr(err error) {
	log.Fatal(err)
}
