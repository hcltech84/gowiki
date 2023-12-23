package main

import (
	"os"
	"log"
	"fmt"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"net/http"
	"github.com/gorilla/mux"
	"html/template"
	"strconv"
)

type Page struct {
	ID int64
	Title string
	Body string
	Author string
}

// Declare a `db` variable of type *sql.DB
var db *sql.DB

func loadPage(id int64) (*Page, error) {
	p := &Page{}
	row := db.QueryRow("SELECT id, title, body, author FROM page WHERE id = ?", id)
	err := row.Scan(&p.ID, &p.Title, &p.Body, &p.Author)
	switch {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("loadPage: No page with id %d", id)
	case err != nil:
		return nil, fmt.Errorf("loadPage: id %d: %v", id, err)
	default:
		return p, nil
	}
}

func renderTemplate(tmpl string, w http.ResponseWriter, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// Convert string to int
	id, _ := strconv.Atoi(vars["id"])
	log.Printf("viewHandler: Received view request for id %v\n", id)
	p, err := loadPage(int64(id))
	if err != nil {
		log.Fatal(err)
	}
	renderTemplate("view", w, p)
}

func main() {
	// Database

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
	
	// Web Server
	r := mux.NewRouter()
	r.HandleFunc("/view/{id:[0-9]+}", viewHandler)
	log.Fatal(http.ListenAndServe(":8083", r))
}