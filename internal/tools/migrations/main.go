//go:build migrate
// +build migrate

package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	database := os.Getenv("MYSQL_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}
	defer db.Close()

	// 最大10回リトライして接続を確認する
	for i := 0; i < 10; i++ {
		err = db.Ping()
		if err == nil {
			log.Println("Database connection successful.")
			break
		}
		log.Printf("Database connection failed. %v", err)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		log.Fatalf("Failed to connect to database after multiple attempts: %v", err)
	}

	m, err := migrate.New(
		"file://database/migrations",
		"mysql://"+dsn,
	)
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 2 {
		log.Fatalf("No command provided. Use up, down, or force <version>.")
	}

	command := os.Args[1]

	switch command {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		log.Println("Migration up completed successfully.")
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		log.Println("Migration down completed successfully.")
	case "force":
		if len(os.Args) < 3 {
			log.Fatalf("No version provided.")
		}
		version, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("Invalid version number: %v", err)
		}
		if err := m.Force(version); err != nil {
			log.Fatal(err)
		}
		log.Println("Migration force completed successfully.")
	case "step":
		if len(os.Args) < 3 {
			log.Fatalf("No step number provided.")
		}
		num, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("Invalid from number: %v", err)
		}
		if err := m.Steps(num); err != nil {
			log.Fatal(err)
		}
		log.Println("Migration step completed successfully.")
	default:
		log.Fatalf("Invalid command provided. Use up, down, or force <version>.", command)
	}
}
