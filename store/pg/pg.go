package pg

import (
	"database/sql"
	"fmt"
)

// Conn contains the postgres connection pool
type Conn struct {
	db *sql.DB
}

// New returns a new instance of Conn
func New() (*Conn, error) {
	// todo - put these vars in a config that gets passed to this function
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s, dbname=%s sslmode=disable", "localhost", "5432", "todouser", "todopass", "todoname"))
}
