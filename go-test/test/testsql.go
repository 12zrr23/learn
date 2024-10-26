package main

import (
  "test/models"

  "gorm.io/gorm"
  "gorm.io/driver/mysql"
)



func main() {
  db, err := gorm.Open(mysql.Open("root:my-secret-pw@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  // Migrate the schema
  db.AutoMigrate(&models.UserBasic{})

  //db.Migrator().DropTable(&models.UserBasic{})
  user := models.UserBasic{}
  user.Username = "test"
  user.Password = "1234"
  // Create
  db.Create(&user)

  // Read
  //var product Product
  //db.First(&product, 1) // find product with integer primary key
  //db.First(&product, "code = ?", "D42") // find product with code D42

  // Update - update product's price to 200
  //db.Model(&product).Update("Price", 200)
  // Update - update multiple fields
  //db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
  //db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

  // Delete - delete product
  //db.Delete(&product, 1)
}