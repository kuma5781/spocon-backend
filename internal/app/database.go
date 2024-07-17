package app

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

func InitDB() (*sql.DB, error) {
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	database := os.Getenv("MYSQL_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

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
		return nil, err
	}

	return db, nil
}
