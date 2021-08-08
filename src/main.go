package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	t "amazonpricetracker/src/trackprice"
)

type Product struct {
	Url string `json:"url"`
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
	mux := http.NewServeMux()
	mux.HandleFunc("/getprice", getPriceHandler)

	http.ListenAndServe(":8080", mux)
}
