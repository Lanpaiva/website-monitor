package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "Alan"
	var versao float32 = 1.1
	fmt.Println("olá senhor", nome, "sua idade é:")
	fmt.Println("versão atual:", versao)

	fmt.Println("O tipo da variavel nome é", reflect.TypeOf(nome))

	fmt.Println("1- Iniciar Monitoramento dos Sites")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do programa")

	var comando int
	fmt.Scan(&comando)

	fmt.Println("o comando escolhido foi", comando)
}
