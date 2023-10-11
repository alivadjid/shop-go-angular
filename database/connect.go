package database

import (
	"awesomeProject/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "pguser"
		password = "pguser"
		dbname   = "angular-go-database"
	)

	//connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	//	host, port, user, password, dbname)
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%d", user, password, dbname, host, port)
	fmt.Println(dsn)
	//database, err := sql.Open("postgres", connectionString)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	fmt.Println(database)
	database.AutoMigrate(&models.User{})
}

//func two() (int, int) {
//	return 3, 5
//}
