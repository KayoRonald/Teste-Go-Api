package main

import (
  "net/http"
  "log"
  "fmt"
  "encoding/json"
)

func main(){
  http.HandleFunc("/", getProduto)
  fmt.Println("Ronando na porta 3030")
  log.Fatal(http.ListenAndServe(":3030", nil))
}

type Produto struct{
  ID int `json:"id"`
  Name string `json:"nome"`
  Valor float32 `json:"preço"`
}

func getProduto(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode([]Produto{{
    ID: 1,
    Name: "Maça",
    Valor: 10.50,
  }})
}