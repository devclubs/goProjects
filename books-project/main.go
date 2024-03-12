package main

import (
	"books_project/dao/mysql"
	"books_project/model"
	"books_project/route"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateTable() {
	db := mysql.MysqlInit()
	if err := db.AutoMigrate(model.User{}, model.Book{}); err != nil {
		fmt.Println("表创建失败", err)
	}
	fmt.Println("表创建创建成功")
}

func main() {
	r := gin.Default()
	r = route.CollectRoute(r)

	CreateTable()

	r.Run(":8000")
}
