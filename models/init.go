package models

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
    gorm.Model
    Code  string
    Price uint
}

var DB *gorm.DB

func init() {
    var err error

    DB, err = gorm.Open("mysql", "root:password@tcp(go_docker_ec2-db)/test_db?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        panic(err)
    }

    // スキーマのマイグレーション
    DB.AutoMigrate(&Product{})

    // Create
    DB.Create(&Product{Code: "shop_price", Price: 4890})

    // Read
    var product Product
    DB.First(&product, 1)
}

