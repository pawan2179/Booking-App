package config

import (
	"database/sql"
	"fmt"

	env "hotel-service/config/env"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func SetupDB() (*sql.DB, error) {
	if DB != nil {
		return DB, nil
	}

	user := env.GetString("DB_USER", "root")
	password := env.GetString("DB_PASSWORD", "")
	addr := env.GetString("DB_ADDRESS", "127.0.0.1:3306")
	dbName := env.GetString("DB_NAME", "hotel")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, addr, dbName)
	fmt.Println("Connecting DB with DSN: ", dsn)

	db, err := sql.Open("mysql", dsn)

	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println("Failed to ping database: ", pingErr)
		return nil, pingErr
	}

	if err != nil {
		fmt.Println("Failed to connect to DB: ", err)
		return nil, err
	}

	fmt.Println("Connected to DB successfully")
	DB = db
	return db, nil
}
