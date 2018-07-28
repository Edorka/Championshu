package categories

import (
	"fmt"
	"encoding/json"
	"net/http"
)

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
