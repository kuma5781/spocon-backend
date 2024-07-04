package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
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

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)
	db, err := sql.Open("mysql", dns)
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
		"mysql://"+dns,
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	log.Println("Migration completed successfully.")
}
