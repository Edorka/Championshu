package main

import (
	"database/sql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"github.com/Edorka/championshu/server/categories"
)

// our main function
func main() {
	db, err := sql.Open("sqlite3", "file:championshu.db?_loc=auto")
	if err != nil {
		panic(err)
	}
	app := categories.CategoriesApp{}
	app.Source(db)
	app.Init()
	if err != nil {
		log.Fatal(err)
	}

	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	router := mux.NewRouter()
	router.HandleFunc("/api/categories", app.List).Methods("GET")
	router.HandleFunc("/api/categories/{id}", app.Get).Methods("GET")
	router.HandleFunc("/api/categories", app.Create).Methods("POST")
	router.HandleFunc("/api/categories/{id}", app.Update).Methods("PUT")
	router.HandleFunc("/api/categories/{id}", app.Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(router)))
}
