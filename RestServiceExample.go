package main

import (
	"fmt"
	"log"
	"net/http"
)

func createEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "add  Call")
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "delete  Call")
}

func main() {
	http.HandleFunc("/create", createEmployee)
	http.HandleFunc("/delete", deleteEmployee)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
