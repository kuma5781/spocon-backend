package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	db, err := sql.Open("mysql", "root:mysql@tcp(localhost:3306)/spocon")
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
		"mysql://root:mysql@tcp(localhost:3306)/spocon",
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	log.Println("Migration completed successfully.")
}
