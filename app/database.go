package app

import (
	"database/sql"
	"eisenhower-todo-api/helper"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

// Function NewDB
// Use for initialize the database connection
// Return *sql.DB
func NewDB() *sql.DB {
	// Load the environment variables file
	err := godotenv.Load(".env")
	helper.PanicIfError(err)

	// Set the database connection
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	helper.PanicIfError(err)

	// Set the database connection pool
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	// SetConnMaxIdleTime sets the maximum amount of time a connection may be reused.
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(60 * time.Minute)
	db.SetConnMaxLifetime(10 * time.Minute)

	return db
}
