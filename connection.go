package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// getConnection obtiene una conexión a la BD
func GetConnection() *sql.DB {
	dsn := "postgres://postgres:password@localhost:5432/postgres?sslmode=disable"
	//postgres://usuario:clave@servidor:puerto/nombreBD?ssl=disable
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
