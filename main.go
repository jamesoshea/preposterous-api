package main

import (
	"fmt"
	"net/http"
	"os"

	"preposterous/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/verb/new", controllers.CreateVerb).Methods("POST")
	router.HandleFunc("/api/verb", controllers.GetVerbs).Methods("GET")
	router.HandleFunc("/api/verb/{id}", controllers.GetVerb).Methods("GET")

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
