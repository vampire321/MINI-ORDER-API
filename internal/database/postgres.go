package database

import (
	"context"
	"time"
	"database/sql"
	"fmt"
	"mini-order-api/internal/config"
	_ "github.com/lib/pq"
)

func Connect(cfg *config.Config) (*sql.DB, error) {  //sql.DB is used in conjuction with the database/sql package to manage database connections and perform operations on the database. It provides a high-level interface for interacting with databases, allowing you to execute queries, manage transactions, and handle connection pooling.
	conStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,cfg.DBPort,cfg.DBUser,cfg.DBPassword,cfg.DBName,
	)

	db, err := sql.Open("postgres", conStr) //sql.open intialize a connection pool for postgres
	if err != nil {
		return nil, fmt.Errorf("failed to open DB: %w", err)
}
	//connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	ctx , cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() //ensures cleanup after existing the function

	if err := db.PingContext(ctx); err != nil{
		return nil, fmt.Errorf("failed to ping DB: %w", err)
	}
	return db, nil
}