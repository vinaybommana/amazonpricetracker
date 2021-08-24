package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	productdetails "amazonpricetracker/src/ProductDetails"
	t "amazonpricetracker/src/trackprice"
)

type Product struct {
	Url string `json:"url"`
}

type GamePrice struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}

func getGamePriceHandler(w http.ResponseWriter, r *http.Request) {
	var gameprices []GamePrice
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	keyword := params["keyword"]
	output := productdetails.GetProductDetails(keyword)
	for name, price := range output {
		gameprices = append(gameprices, GamePrice{Name: name, Price: price})
	}
	json.NewEncoder(w).Encode(gameprices)
}

func getPriceHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	// body, _ := ioutil.ReadAll(r.Body)
	// fmt.Println(string(body))

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&product)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(t.GetProductPrice(product.Url))
	fmt.Fprintln(w, t.GetProductPrice(product.Url))
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/getprice", getPriceHandler)

	mux.HandleFunc("/gameprice/{keyword}", getGamePriceHandler)

	http.ListenAndServe(":8080", mux)
}
