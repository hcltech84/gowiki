package main

import (
	"os"
	"log"
	"fmt"
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

type Page struct {
	ID int64
	Title string
	Body []byte
	Author string
}

// Declare a `db` variable of type *sql.DB
var db *sql.DB

func main() {
	// Configure a database connection
	cfg := mysql.Config{
		User: os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net: "tcp",
		Addr: "127.0.0.1:3306",
		DBName: "gowiki",
	}

	// Open a database handle
	var err error;
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	// Test the database connection
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected to MySQL")
}