package models

type User struct {
	Id        string
	FirstName string
	LastName  string
	Email     string `gorm:"unique"`
	Password  string
}
