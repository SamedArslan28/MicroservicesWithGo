package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	// Import PostgreSQL drivers
	_ "github.com/jackc/pgconn"        // Provides PostgreSQL connection functionalities.
	_ "github.com/jackc/pgx/v4"        // A PostgreSQL driver and toolkit.
	_ "github.com/jackc/pgx/v4/stdlib" // Provides a standard library `database/sql` driver for PostgreSQL.
)

var retryCount int64 // Counter to track the number of retry attempts for database connection.

/**
 * openDB opens a connection to the database using the provided DSN (Data Source Name)
 * and driver. It validates the connection by attempting a ping to the database.
 *
 * @param dsn The Data Source Name string used for connecting to the database.
 * @param driver The name of the database driver (e.g., "pgx").
 * @return A pointer to the sql.DB instance or an error if the connection fails.
 */
func openDB(dsn string, driver string) (*sql.DB, error) {
	db, err := sql.Open(driver, dsn) // Open a new database connection.
	if err != nil {
		log.Fatal(err) // Log and exit if the connection cannot be opened.
	}

	err = db.Ping() // Ping the database to ensure the connection is alive.
	if err != nil {
		return nil, err // Return an error if the ping fails.
	}

	return db, nil // Return the established database connection.
}

// ConnectToDB
/*
 * ConnectToDB attempts to establish a connection to the database with retry logic.
 * It reads the DSN from the environment variable "DSN" and uses the specified driver
 * to connect. If the connection fails, it retries up to 10 times with a delay between attempts.
 *
 * @param driver The name of the database driver (e.g., "pgx").
 * @return A pointer to the sql.DB instance if the connection succeeds, or nil if it fails.
 */
func ConnectToDB(driver string) *sql.DB {
	dsn := os.Getenv("DSN") // Get the DSN from the environment variables.

	for {
		// Attempt to open a connection to the database.
		db, err := openDB(dsn, driver)
		if err != nil {
			log.Println("Error connecting to database: ", err) // Log the connection error.
			retryCount++                                       // Increment the retry counter.
		} else {
			log.Println("Connected to database") // Log success.
			return db                            // Return the successful connection.
		}

		// If retry attempts to exceed the limit, log and stop retrying.
		if retryCount > 10 {
			log.Println("Could not connect to database") // Log failure after exceeding retries.
			return nil
		}

		// Log retry attempt and wait for 2 seconds before retrying.
		log.Println("Retrying connection to database")
		time.Sleep(2 * time.Second)
	}
}
