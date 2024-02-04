package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	gorm.Model // permite que trabalhe com soft delete...
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	// // create
	// db.Create(&Product{
	// 	Name:  "Laptop",
	// 	Price: 1000.0,
	// })

	// create batch
	// products := []Product{
	// 	{Name: "Product 1", Price: 1000.00},
	// 	{Name: "Product 2", Price: 50.00},
	// 	{Name: "Product 3", Price: 100.00},
	// }
	// db.Create(&products)

	// select one
	// var product Product
	// db.First(&product, 1)
	// fmt.Println(product)

	// db.First(&product, "name = ?", "Product 3")
	// fmt.Println(product)

	// select all
	// var products []Product
	// db.Limit(2).Offset(0).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// where
	// var products []Product
	// db.Where("price > ?", 80).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// LIKE
	// var products []Product
	// db.Where("name LIKE ?", "%Mouse%").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// Update
	// var product Product
	// db.First(&product, 1)
	// product.Name = "New Mouse"
	// db.Save(&product)

	// var product2 Product
	// db.First(&product2, 1)
	// fmt.Println(product2.Name)

	// Delete
	// var product Product
	// db.First(&product, 1)
	// db.Delete(&product)
}
