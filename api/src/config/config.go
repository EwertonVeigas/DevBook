package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StingConexaoBanco = ""
	Porta             = 0
)

// inicializar as variaveis de ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 9000
	}

	StingConexaoBanco = fmt.Sprintf("golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local")
	//		os.Getenv("DB_USUARIO"),
	//		os.Getenv("DB_SENHA"),
	//		os.Getenv("DB_NOME"),

}
