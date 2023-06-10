package main

import (
	"log"
	"net/http"

	"github.com/Athul-k-ak/crud/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/managers", handlers.CreateManager).Methods("POST")
	router.HandleFunc("/managers", handlers.GetManagers).Methods("GET")
	router.HandleFunc("/managers/{email}", handlers.DeleteManager).Methods("DELETE")
	router.HandleFunc("/employees", handlers.CreateEmployee).Methods("POST")
	router.HandleFunc("/employees", handlers.GetEmployees).Methods("GET")
	router.HandleFunc("/employees/{email}", handlers.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/employees", handlers.DeleteEmployee).Methods("DELETE")

	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
