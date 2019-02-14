package dbHelper

import (
	"database/sql"
)

const (
	DEFAULT_ROW_COUNT_PER_QUERY = 20
)

var db *sql.DB

func Init(theDb *sql.DB) {
	db = theDb
}
