package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"encoding/json"
)

type Product struct {
	Id int
	Name string 
}

var products [10] Product

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	    fmt.Fprintf(w, "<h1>Hello Mux Rest !</h1>")
}

func ProductsHandler( w http.ResponseWriter, r *http.Request) {
	productsJson, _ := json.Marshal(products)
	fmt.Fprintf(w,string(productsJson))
}


func main() {
	r:=mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products", ProductsHandler).Methods("GET").Headers("Accept", "application/json")
	http.Handle("/",r)
	http.ListenAndServe(":8080", nil)

}
