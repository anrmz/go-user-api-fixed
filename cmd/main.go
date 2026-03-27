package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"github.com/matthewjamesboyle/golang-interview-prep/internal/user"
)

func main() {

	godotenv.Load()

	runMigrations()

	svc, err := user.NewService()
	if err != nil {
		log.Fatal(err)
	}

	h := user.Handler{Svc: svc,}

	http.HandleFunc("/user", h.AddUser)

	log.Println("starting http server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func runMigrations() {
	// Database connection string
	dbURL := fmt.Sprintf(
	"postgres://%s:%s@%s:%s/%s?sslmode=disable",
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_HOST"),
	os.Getenv("DB_PORT"),
	os.Getenv("DB_NAME"),
)

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a new instance of the PostgreSQL driver for migrate
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://internal/migrations", "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	fmt.Println("Database migration complete.")
}
