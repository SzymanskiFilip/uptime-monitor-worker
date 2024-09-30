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

	//STATISTICS TABLE
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS statistics (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		url TEXT NOT NULL	
	)
	`)
	if err != nil{
		log.Fatal("error")
	}



	log.Println("Database initialized successfully.")
}

func GetDB() *sql.DB{
	return db
}