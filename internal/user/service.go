package user

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	db *sql.DB
}

func NewService() (*service, error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return &service{db: db}, nil
}

type User struct {
	Name     string `json:"username"`
	Password string `json:"password"`
}

func (s *service) AddUser(u User) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	var id string

	q := "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id"

	err = s.db.QueryRow(q, u.Name, string(hashedPassword)).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("failed to insert: %w", err)
	}

	return id, nil
}