package models

type Book struct {
	ID       string `gorm:"primary_key"`
	Title    string `gorm:"not null"`
	Author   string `gorm:"not null"`
	Quantity int    `gorm:"not null"`
}

var books = []Book{
	{ID: "2324", Title: "Introduction to Golang", Author: "Satya", Quantity: 500},
	{ID: "2325", Title: "Introduction to Docker", Author: "Vamsi", Quantity: 600},
	{ID: "2326", Title: "Introduction to ReactJs", Author: "Raghava", Quantity: 700},
	{ID: "2327", Title: "Introduction to Database", Author: "Ashok", Quantity: 800},
}
