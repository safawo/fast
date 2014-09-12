package ds

import (
	"database/sql"
)

func DB() *sql.DB {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=threemanonroad dbname=think sslmode=disable")
	verifyErr(err)
	return db
}

func verifyErr(err error) {
	if err != nil {
		panic(err)
	}

}
