package main

import (
	"encoding/json"
	"net/http"
)

type Person struct {
	Name string
}

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	u := Person{Name: "US123"}

	json.NewEncoder(w).Encode(u)

}
