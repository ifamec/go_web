package utils

import "database/sql"

var (
	Db *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", "root:00000000@tcp(localhost:3306)/bookstore?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
}
