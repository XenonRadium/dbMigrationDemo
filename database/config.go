package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Open() *sql.DB {
	db, err := sql.Open("mysql", "root:bernard@/migrationDB")
	if err != nil {
		fmt.Printf("error opening db: %v \n", err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("error pinging db: %v \n", err)
	}

	fmt.Println("Connected to DB!")

	return db
}
