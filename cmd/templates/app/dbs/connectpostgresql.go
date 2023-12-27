package dbs

import (
	"database/sql"
	"fmt"
)

var (
	connection *sql.DB
)

func Connect() error {
	var err error
	var connectionUri = fmt.Sprintf("")
	connection, err = sql.Open("postgres", connectionUri)
	if err != nil {
		return err
	}
	return connection.Ping()
}
