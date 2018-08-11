package categories

import (
	"encoding/json"
	"net/http"
	"strconv"
	"fmt"
	"github.com/gorilla/mux"
)

func (app *CategoriesApp) Update(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	targetId, err := strconv.ParseInt(params["id"], 10, 32)
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
	update.Exec(&target.Style, &target.Division, targetId)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(target)
}
