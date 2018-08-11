package categories

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
        "fmt"
)

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
        if err != nil {
		http.Error(response, fmt.Sprintf("Cannot retrieve information: ", err), 500)
		return
	}
	json.NewEncoder(response).Encode(result)
}
