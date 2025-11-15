package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func initDB() *sql.DB {
	// load env file (database.env)
	_ = godotenv.Load("database.env")

	connstr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"), os.Getenv("SSL_MODE"),
	)

	db, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}

func createUserInfoTable(db *sql.DB) {
	_, err := db.Exec(`CREATE EXTENSION IF NOT EXISTS pgcrypto;`)
	if err != nil {
		log.Fatal(err)
	}

	query := `CREATE TABLE IF NOT EXISTS user_account (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name TEXT NOT NULL,
    lastname TEXT NOT NULL,
    age INT NOT NULL,
    height INT NOT NULL,
    weight DECIMAL(5,2) NOT NULL,
    gender TEXT NOT NULL CHECK (gender IN ('Male','Female')),
    plan TEXT NOT NULL CHECK (plan IN ('normal','semi-interactive','fully-interactive')),
    created TIMESTAMP DEFAULT now(),
    skilllevel TEXT NOT NULL CHECK (skilllevel IN ('Beginner','Intermediate','Master')),
    email TEXT NOT NULL
	);`

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

func createCoachInfoTable(db *sql.DB) {
	_, err := db.Exec(`CREATE EXTENSION IF NOT EXISTS pgcrypto;`)
	if err != nil {
		log.Fatal(err)
	}

	// use lowercase table name to avoid quoting issues; cost DECIMAL to match float64
	query := `CREATE TABLE IF NOT EXISTS coach_account (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name TEXT NOT NULL,
    lastname TEXT NOT NULL,
    age INT NOT NULL,
    gender TEXT NOT NULL CHECK (gender IN ('Male','Female')),
    created TIMESTAMP DEFAULT now(),
    email TEXT NOT NULL,
    cost DECIMAL(10,2) NOT NULL,
    verified BOOLEAN NOT NULL DEFAULT false,
    profile_pic TEXT,
	bio VARCHAR(5000)
	);`

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}
