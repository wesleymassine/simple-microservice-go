package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (

	// URL Microservice
	ProductURL = ""

	CheckoutURL = ""

	// StringConnection é a string de coneão com o BD
	StringConnection = ""

	//Porta onde a API vai estar rodando
	Porta = 0

	//SecretKey é a chave que vai ser usada para assinar o token
	SecretKey []byte
)

//Carregar vai inicializar as variaveis de ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 8080
	}

	ProductURL = os.Getenv("PRODUCT_URL")
	CheckoutURL = os.Getenv("CHECKOUT_URL")

	// StringConnection = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_NAME"),
	// )

	// SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
