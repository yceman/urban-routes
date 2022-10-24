package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Cliente struct {
	Nome  string
	Email string
	CPF   int
}

/*Método com argumento, esteja atrelado no Cliente usando abreviação "c" para cliente*/

func (c Cliente) Exibe() {
	fmt.Println("Exibindo cliente pelo método", c.Nome)
}

/*Cliente brasileiro com cpf, nome e email mas mora em outro país*/
/*Para formatar o texto do json uma opção é `json:"pais"`*/
type ClienteInternacional struct {
	Cliente
	/*Nome  string
	Email string
	CPF   int*/
	Pais string `json:"pais"`
}

func main() {

	cliente := Cliente{
		Nome:  "Marcos",
		Email: "marcosmorais.v10@gmail.com",
		CPF:   12345678900,
	}

	fmt.Println("Hello!!, Marcos")
	fmt.Println(cliente)

	/*cliente2 - Outro tipo de declarar os dados*/

	cliente2 := Cliente{"Mari", "mari@gmail.com", 98765432111}
	fmt.Printf("Nome: %s, Email: %s, CPF: %d\n", cliente2.Nome, cliente2.Email, cliente2.CPF)

	cliente3 := ClienteInternacional{
		Cliente: Cliente{
			Nome:  "Davi",
			Email: "davi@gmail.com",
			CPF:   12312312300,
		},
		Pais: "EUA",
		/*Nome:  "Davi",
		Email: "davi@gmail.com",
		CPF:   12312312300,
		Pais:  "EUA",*/
	}
	fmt.Printf("Nome: %s, Email: %s, Pais: %s\n", cliente3.Nome, cliente3.Email, cliente3.Pais)
	cliente.Exibe()
	cliente2.Exibe()
	cliente3.Exibe()

	/*Transformando uma struct em Json*/

	clienteJson, err := json.Marshal(cliente3)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string(clienteJson))

	/*Transformar Json para struct*/
	jsonCliente4 := `{"Nome":"Elias","Email":"elias@gmail.com","CPF":12332112300,"Pais":"EUA"}`
	cliente4 := ClienteInternacional{}

	json.Unmarshal([]byte(jsonCliente4), &cliente4)
	fmt.Println(cliente4)
}
