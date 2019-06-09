package main

import (
	"fmt"
	"net/http"
	"os"
	"preposterous/controllers"

	"github.com/gorilla/mux"
)

func main() {

	// filename := "verben.csv"

	// f, err := os.Open(filename)
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// lines, err := csv.NewReader(f).ReadAll()
	// if err != nil {
	// 	panic(err)
	// }

	// for _, line := range lines {
	// 	verb := &models.Verb{
	// 		German:          line[0],
	// 		Preposition:     line[1],
	// 		GrammaticalCase: line[2],
	// 		English:         line[3],
	// 	}
	// 	verb.Create()
	// }

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
