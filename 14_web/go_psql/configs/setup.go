package configs

import (
	"database/sql"
	"fmt"
	"log"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=employee_database sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("DB Connected...")
	}

	return db
}
