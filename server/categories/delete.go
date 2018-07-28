package categories

import (
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
        "fmt"
)

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
