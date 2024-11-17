package main

import (
	"github.com/carlosclavijo/pizza/internal/database"
)

//var db *gorm.DB

func main() {
	go database.GetInstance()

}
