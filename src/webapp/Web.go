package main

import (
        "github.com/gorilla/mux"
        "net/http"
        "log"
        "fmt"
        "encoding/json"
        "strconv"
)

type Product struct {
        Id int
        Name string 
}

var products [] Product


func ProductsHandler( w http.ResponseWriter, r *http.Request) {
        //fmt.Fprintf(w, "<h1>Our Products </h1>")
        productsJson, _ := json.Marshal(products)
        fmt.Fprintf(w,"<h4>" +string(productsJson)+"</h4>")
        fmt.Fprintf(w,"<br/> <a href=\"/\">Home</a>")
}

func AddProductHandler(w http.ResponseWriter, r *http.Request) {
        log.Println(r.RequestURI, r.URL)
        r.ParseForm()
        log.Println(r.Form)
        var newProd Product
        newProd.Id,_=strconv.Atoi(r.FormValue("Id"))
        newProd.Name=r.FormValue("Name")
        products=append(products, newProd)
        fmt.Println(products)
        http.Redirect(w,r,"../",301)
}

func main() {
        r:=mux.NewRouter()
        r.HandleFunc("/products/", ProductsHandler).Methods("GET")
        r.HandleFunc("/products/add/", AddProductHandler).Methods("POST")
        r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
        http.Handle("/",r)
        http.ListenAndServe(":8080", nil)

}

