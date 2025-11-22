package lite

import (
	"database/sql"
	"fmt"

	"github.com/attendeee/url-shortener/storage/db"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB
var Query *db.Queries

// Inits sqlite (Db) and queries (Query) variables
func Init() {
	var err error
	Db, err = sql.Open("sqlite3", "./data.sqlite")
	if err != nil {
		panic(fmt.Sprintf("Sqlite database init error: %s", err))
	}

	Query = db.New(Db)
}
