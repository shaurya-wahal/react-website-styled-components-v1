package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Car struct {
	gorm.Model
	Brand string
	Type  string
	Price float64
	MPG   int
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&Car{})
	channels := []Car{
		{Brand: "string", Type: "Sedan", Price: 5000, MPG: 25},
	}
	for _, c := range channels {
		db.Create(&c)
	}
	var cars []Car
	db.Find(&cars)
	for _, c := range cars {
		fmt.Println(c.Brand)
	}
}
