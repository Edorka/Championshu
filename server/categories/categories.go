package categories

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"strconv"
)

type Category struct {
	ID       int    `json:"id"`
	Style    string `json:"style,omitempty"`
	Division string `json:"division,omitempty"`
}

type Categories struct {
	Items []Category `json:"categories,omitempty"`
}
type CategoriesApp struct {
	Access *sql.DB
}

func (app *CategoriesApp) Source(source *sql.DB) {
	app.Access = source
}

func (app *CategoriesApp) Init() {
	statement := "create table IF NOT EXISTS categories (id integer PRIMARY KEY, style varchar(255) NOT null, division varchar(255) NOT null)"
	_, err := app.Access.Exec(statement)
	if err != nil {
		log.Fatal(err)
	}

}
func (app *CategoriesApp) List(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	const query = `SELECT id, style, division from categories`
	selection, err := app.Access.Query(query)
	if err != nil {
		panic(err.Error())
	}
	var categories []Category
	for selection.Next() {
		var current Category = Category{}
		err = selection.Scan(&current.ID, &current.Style, &current.Division)
		if err != nil {
			panic(err.Error())
		}
		categories = append(categories, current)
	}
	result := Categories{Items: categories}
	json.NewEncoder(response).Encode(result)
}

func (app *CategoriesApp) Get(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	expectedId, err := strconv.ParseInt(params["id"], 10, 32)
	if err != nil {
		http.Error(response, fmt.Sprintf("malformed id: %v", params["id"]), 400)
		return
	}
	const query = `SELECT id, style, division from categories where id = $1 `
	var result Category = Category{}
	err = app.Access.QueryRow(query, expectedId).Scan(&result.ID, &result.Style, &result.Division)
	json.NewEncoder(response).Encode(result)

}

func (app *CategoriesApp) Create(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var newCategory Category
	_ = json.NewDecoder(request.Body).Decode(&newCategory)
	const query = "INSERT INTO categories(style, division) VALUES(?,?)"
	insertion, err := app.Access.Prepare(query)
	if err != nil {
		http.Error(response, fmt.Sprintf("Cannot prepare statement: ", err), 500)
		return
	}
	res, err := insertion.Exec(&newCategory.Style, &newCategory.Division)
	if err != nil {
		http.Error(response, fmt.Sprintf("Cannot run insert statement: ", err), 500)
		return
	}
	insertionID, _ := res.LastInsertId()
	newCategory.ID = int(insertionID)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(newCategory)
}

func (app *CategoriesApp) Update(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	expectedId, err := strconv.ParseInt(params["id"], 10, 32)
	if err != nil {
		http.Error(response, fmt.Sprintf("malformed id: %v", params["id"]), 400)
		return
	}
	const query = "UPDATE categories SET style=?, division=?  where id = ?"
	update, err := app.Access.Prepare(query)
	if err != nil {
		panic(err.Error())
	}
	var target Category
	_ = json.NewDecoder(request.Body).Decode(&target)
	update.Exec(&target.Style, &target.Division, expectedId)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(target)
}

func (app *CategoriesApp) Delete(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	expectedId, err := strconv.ParseInt(params["id"], 10, 32)
	if err != nil {
		http.Error(response, fmt.Sprintf("malformed id: %v", params["id"]), 400)
		return
	}
	const query = "DELETE FROM categories where id = ?"
	update, err := app.Access.Prepare(query)
	if err != nil {
		panic(err.Error())
	}
	update.Exec(expectedId)
	response.WriteHeader(http.StatusOK)
}
