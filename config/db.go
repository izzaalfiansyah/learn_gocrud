package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func DB() (*sql.DB, error) {
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
	return db, nil
}
