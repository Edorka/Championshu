package categories

import (
	"encoding/json"
	"net/http"
)

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
