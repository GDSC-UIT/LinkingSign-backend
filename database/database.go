package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" // postgres golang driver
)

func CreateConnection() *sql.DB {
	// load .env file
	// fmt.Println("Loading .env file...")

	// err := godotenv.Load()

	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }

	// Open the connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected database!")
	// return the connection
	return db
}
