package main

import (
    "encoding/json"
    "net/http"
)

type Product struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
    Paid bool `json:"paid"`
}

var products = []Product{
	{ID: 1, Name: "Wiertarka udarowa", Price: 100, Paid: false},
	{ID: 2, Name: "Młot pneumatyczny", Price: 200, Paid: false},
	{ID: 3, Name: "Gwordziarka", Price: 300, Paid: false},
	{ID: 4, Name: "Imadło", Price: 400, Paid: true},
	{ID: 5, Name: "Nitownica", Price: 500, Paid: false},
}

func main() {
    http.HandleFunc("/products", ProductsHandler)
    http.HandleFunc("/products/pay", PaymentHandler)
    http.ListenAndServe(":8080", nil)
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
    jData, err := json.Marshal(products)
    if err != nil {
        // handle error
    }
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Write(jData)
}

func PaymentHandler(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		ID int `json:"id"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	productID := requestData.ID
	products[productID - 1].Paid = true
}


