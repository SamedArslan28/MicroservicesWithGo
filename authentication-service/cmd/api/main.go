package main

import (
	"authentication/data"
	"authentication/db"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"net/http"
)

const webPort = "80"

var counts int64

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	fmt.Printf("Starting server on port %s\n", webPort)
	conn := db.ConnectToDB("pgx")

	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}

	srv := http.Server{
		Addr:    ":" + webPort,
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		return
	}
}
