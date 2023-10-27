package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/natanaelrusli/hexagonal-arch/internals/core/domain"
)

type PostgresAccountRepository struct {
	db *sql.DB
}

func NewPostgresAccountRepository() (*PostgresAccountRepository, error) {
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")

	connStr := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=5432 sslmode=disable", dbHost, dbUser, dbName, dbPassword)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("postgres database connected")

	return &PostgresAccountRepository{
		db: db,
	}, nil
}

func (r *UserRepository) CreateAccount(*domain.Account) error {
	return nil
}
