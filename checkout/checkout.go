package main

import (
	"checkout/config"
	"checkout/queue"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	Uuid    string  `json:"uuid`
	Product string  `json:"product"`
	Price   float32 `json:"price,string"`
}

type Order struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	ProductId string `json:"product_id"`
}

func init() {
	config.Carregar()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/finish", finish)
	r.HandleFunc("/checkout/{id}", displayCheckout)

	fmt.Println("Application Started listening on port:", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}

func displayCheckout(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	response, err := http.Get(config.ProductURL + "/products/" + vars["id"])
	if err != nil {
		fmt.Println("The HTTP request failed with erroe %s\n", err)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Err read data %s\n", err)
	}

	var product Product
	json.Unmarshal(data, &product)

	t := template.Must(template.ParseFiles("templates/checkout.html"))
	t.Execute(w, product)
}

func finish(w http.ResponseWriter, r *http.Request) {
	var order Order

	order.Name = r.FormValue("name")
	order.Email = r.FormValue("email")
	order.Phone = r.FormValue("phone")
	order.ProductId = r.FormValue("product_id")

	data, _ := json.Marshal(order)
	fmt.Println(string(data))

	connection := queue.Connect()
	queue.Notify(data, "checkout_ex", "", connection)

	w.Write([]byte("Processou!"))
}
