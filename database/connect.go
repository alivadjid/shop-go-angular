package database

import (
	"database/sql"
	"fmt"
)

import (
	_ "github.com/lib/pq"
)

func Connect() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "pguser"
		password = "pguser"
		dbname   = "angular-go-database"
	)

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	// defer db.Close()
	fmt.Println(db)

	n1, n2 := two()
	fmt.Println(n1, n2)

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to the database!")
}

func two() (int, int) {
	return 3, 5
}
