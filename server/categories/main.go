package categories

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
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
