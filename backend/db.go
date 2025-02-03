package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn

func connectDB() {
	databaseURL := "postgres://postgres:123456789@localhost:5432/crud_golang_db"

	var err error

	conn, err = pgx.Connect(context.Background(), databaseURL)

	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Connected to DB successfully !!!")
}

func CloseDB() {
	conn.Close(context.Background())
	fmt.Println("DB connection close")
}
