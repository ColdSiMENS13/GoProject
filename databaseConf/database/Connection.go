package config

import "fmt"

func GetConnectionString() string {
	database := NewConfig()
	conString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		database.DataBase.Host, database.DataBase.Port, database.DataBase.UserName, database.DataBase.Password, database.DataBase.DataBaseName)

	return conString
}
