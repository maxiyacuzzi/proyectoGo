package config

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
)

var DB *sql.DB

func Connect() {
    var err error
    connStr := "user=usuario dbname=db-pruebas sslmode=disable password=66276"
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }

    err = DB.Ping()
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Database connected successfully")
}
