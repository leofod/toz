package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func ConnectDB(cfg Config) (*sql.DB, error) {

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	performMigrations(db)

	return db, nil
}

func performMigrations(db *sql.DB) error {
	_, err := db.Exec(`
	 	CREATE TABLE IF NOT EXISTS urls (
	 		short_url VARCHAR(10) UNIQUE,
	 		full_url VARCHAR(255)
	 	)
	`)
	fmt.Println(err)
	return err
}
