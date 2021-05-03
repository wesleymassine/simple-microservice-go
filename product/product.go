package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"product/config"

	"github.com/gorilla/mux"
)

type Product struct {
	Uiid    string  `json:"uuid"`
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

	r.HandleFunc("/products", ListProducts)
	r.HandleFunc("/products/{id}", ṔroductById)

	fmt.Println("Application Started listening on port:", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}

func loadData() []byte {
	jsonFile, err := os.Open("products.json")

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)

	return data
}

func ListProducts(w http.ResponseWriter, r *http.Request) {
	products := loadData()

	w.Write([]byte(products))
}

func ProductById(w http.ResponseWriter, r *http.Request) {
	productID := mux.Vars(r)
	data := loadData()

	var products Products
	json.Unmarshal(data, &products)

	for _, product := range products.Products {
		if product.Uiid == productID["id"] {
			value, _ := json.Marshal(product)
			w.Write([]byte(value))
		}
	}
}
