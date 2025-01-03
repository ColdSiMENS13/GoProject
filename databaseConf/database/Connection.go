package config

import (
	"database/sql"
	"fmt"
	"log"
)

func OpenConnection() *sql.DB {
	conString := getConnectionString()
	db, err := sql.Open("postgres", conString)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func getConnectionString() string {
	database := NewConfig()
	conString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		database.DataBase.Host, database.DataBase.Port, database.DataBase.UserName, database.DataBase.Password, database.DataBase.DataBaseName)

	return conString
}
