package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model // permite que trabalhe com soft delete...
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	// create category
	// category := Category{Name: "Eletronicos"}
	// db.Create(&category)

	// create product
	// db.Create(&Product{
	// 	Name:       "Mouse",
	// 	Price:      1000.00,
	// 	CategoryID: category.ID,
	// })

	var products []Product
	db.Preload("Category").Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}
}
