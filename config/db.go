package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() (*sql.DB, error) {
	connectionString := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", Env.DBUsername, Env.DBPassword, Env.DBHost, Env.DBPort, Env.DBName)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Database connected")
	DB = db
	return db, nil
}
