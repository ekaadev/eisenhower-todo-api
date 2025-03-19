package app

import (
	"database/sql"
	"eisenhower-todo-api/helper"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func NewDB() *sql.DB {
	err := godotenv.Load("../.env")
	helper.PanicIfError(err)

	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(60 * time.Minute)
	db.SetConnMaxLifetime(10 * time.Minute)

	return db
}
