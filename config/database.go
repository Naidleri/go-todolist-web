package config

import (
	"database/sql"
	"log"
	_ "time"

	_ "github.com/go-sql-driver/mysql"
)

	var DB *sql.DB

func ConnectDb () {
	db, err := sql.Open("mysql", "root:root@/todolist_go")
	if err != nil {
	panic(err)
	}

	log.Println("Db connected")
	DB = db
}