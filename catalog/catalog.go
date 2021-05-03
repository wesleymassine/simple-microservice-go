package main

import (
	"catalog/config"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	Uuid    string  `json:"uuid"`
	Product string  `json:"product"`
	Price   float64 `json:"price,string"`
}

type Products struct {
	Products []Product
}

func init() {
	config.Carregar()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", ListProducts)
	r.HandleFunc("/products/{id}", ShowProduct)
	fmt.Println("Application Started listening on port:", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}

func loadProducts() []Product {
	response, err := http.Get(config.ProductURL + "/products")
	if err != nil {
		fmt.Println("Erro de HTTP")
	}

	data, _ := ioutil.ReadAll(response.Body)

	var products Products
	json.Unmarshal(data, &products)

	return products.Products
}

func ListProducts(w http.ResponseWriter, r *http.Request) {
	products := loadProducts()
	t := template.Must(template.ParseFiles("templates/catalog.html"))
	t.Execute(w, products)
}

func ShowProduct(w http.ResponseWriter, r *http.Request) {
	productID := mux.Vars(r)

	response, err := http.Get(config.ProductURL + "/products/" + productID["id"])
	if err != nil {
		fmt.Println("The HTTP request failed with error %s\n", err)
	}

	data, _ := ioutil.ReadAll(response.Body)

	var product Product
	json.Unmarshal(data, &product)

	t := template.Must(template.ParseFiles("templates/view.html"))
	t.Execute(w, product)
}
