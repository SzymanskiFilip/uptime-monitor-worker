package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitializeDatabase(){
	fmt.Println("Initializing database...")

	connStr := "postgres://user:password@localhost:5430/pinger_database?sslmode=disable"
	database, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	db = database

	_, err = db.Exec(`CREATE EXTENSION IF NOT EXISTS "pgcrypto";`)
	if err != nil{
		log.Fatal(err)
	}

	//DOMAINS TABLE
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS urls (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		url TEXT NOT NULL UNIQUE
	);
	`)
	if err != nil {
		log.Fatal(err)
	}

	//STATISTICS TABLE
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS statistics (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    url_id UUID NOT NULL,
    headers TEXT NOT NULL,
    status INT NOT NULL CHECK (status >= 100 AND status < 600),
    success BOOLEAN NOT NULL,
    response_time INT NOT NULL CHECK (response_time >= 0),
	FOREIGN KEY (url_id) REFERENCES urls (id)
	);
	`)
	if err != nil{
		log.Fatal(err)
	}



	log.Println("Database initialized successfully.")
}

func GetDB() *sql.DB{
	return db
}