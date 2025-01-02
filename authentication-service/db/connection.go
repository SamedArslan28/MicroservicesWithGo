package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var retryCount int64

// openDB opens a connection to the database using the provided DSN.
func openDB(dsn string, driver string) (*sql.DB, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// ConnectToDB attempts to establish a connection to the database with retry logic.
func ConnectToDB(driver string) *sql.DB {
	dsn := os.Getenv("DSN")

	for {
		db, err := openDB(dsn, driver)
		if err != nil {
			log.Println("Error connecting to database: ", err)
			retryCount++
		} else {
			log.Println("Connected to database")
			return db
		}

		if retryCount > 10 {
			log.Println("Could not connect to database")
			return nil
		}

		log.Println("Retrying connection to database")
		time.Sleep(2 * time.Second)
	}
}
