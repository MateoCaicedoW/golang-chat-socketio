package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var Tx *sql.DB

func init() {
	psqlInfo := "host=localhost port=5432 user=postgres password=postgres dbname=socket sslmode=disable"
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("running without database connection")
	}

	Tx = db
	fmt.Println("Successfully connected!")
}
