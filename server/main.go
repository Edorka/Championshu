package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/handlers"
    "github.com/gorilla/mux"
)


type Category struct {
    ID        int      `json:"id"`
    Style     string   `json:"style,omitempty"`
    Division  string   `json:"division,omitempty"`
}

type Categories struct {
    Items  []Category   `json:"categories,omitempty"`
}
var categories []Category;

func GetCategories(w http.ResponseWriter, r *http.Request) {
    var result Categories = Categories{Items: categories}
    json.NewEncoder(w).Encode(result); 
}

func GetCategory(w http.ResponseWriter, r *http.Request) {}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
    var category Category
    _ = json.NewDecoder(r.Body).Decode(&category)
    categories = append(categories, category)
    json.NewEncoder(w).Encode(category)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {}

// our main function
func main() {
    categories = append(categories, Category{ID: 1, Style: "Chang Quan", Division: "Male"})
    categories = append(categories, Category{ID: 2, Style: "Chang Quan", Division: "Female"})
    categories = append(categories, Category{ID: 3, Style: "Chang Quan", Division: "Jr Mixed"})

    allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
    allowedOrigins := handlers.AllowedOrigins([]string{"*"})
    allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

    router := mux.NewRouter()
    router.HandleFunc("/categories", GetCategories).Methods("GET")
    router.HandleFunc("/categories", CreateCategory).Methods("POST")
    router.HandleFunc("/categories/{id}", GetCategory).Methods("GET")
    router.HandleFunc("/categories/{id}", DeleteCategory).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8000",handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(router)))
}
