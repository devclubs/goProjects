package controller

import (
	"books_project/dao/mysql"
	"books_project/model"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateBookHandler(c *gin.Context) {
	var bk model.Book

	err := c.ShouldBind(&bk)

	if err != nil {
		resData := struct {
			Code int         `json:"code"`
			Data interface{} `json:"data"`
			Msg  string      `json:"msg"`
		}{400, "", err.Error()}
		c.JSON(400, resData)
		return
	}

	mysql.DB.Create(&bk)

	resData := struct {
		Code int         `json:"code"`
		Data interface{} `json:"data"`
		Msg  string      `json:"msg"`
	}{200, bk, "书籍添加成功"}
	c.JSON(200, resData)

}

func GetBookListHandler(c *gin.Context) {
	books := []model.Book{}
	mysql.DB.Find(&books)
	resData := struct {
		Code int         `json:"code"`
		Data interface{} `json:"data"`
		Msg  string      `json:"msg"`
	}{200, books, "获取书籍列表成功"}
	c.JSON(200, resData)
}

func GetBookDetailHandler(c *gin.Context) {
	// id := c.Query("id") //http://127.0.0.1:8000/v1/book/detial?id=5
	id := c.Param("id") // http://127.0.0.1:8000/v1/book/detial/4
	fmt.Printf("%T, %[1]v+++++++++++++++++\n", id)
	if id == "" {
		resData := struct {
			Code int         `json:"code"`
			Data interface{} `json:"data"`
			Msg  string      `json:"msg"`
		}{400, "", "请输入书籍ID"}
		c.JSON(400, resData)
	}

	bid, _ := strconv.ParseInt(id, 10, 64)

	var book = model.Book{Id: bid}
	mysql.DB.Find(&book)
	resData := struct {
		Code int         `json:"code"`
		Data interface{} `json:"data"`
		Msg  string      `json:"msg"`
	}{200, book, "success"}
	c.JSON(200, resData)
}

func UpdateBookDetailHandler(c *gin.Context) {
	var bk model.Book
	err := c.ShouldBind(&bk)

	if err != nil {
		resData := struct {
			Code int         `json:"code"`
			Data interface{} `json:"data"`
			Msg  string      `json:"msg"`
		}{400, "", "信息修改错误"}
		c.JSON(400, resData)
		return
	}

	
	mysql.DB.Model(&model.Book{}).Where("id=?", bk.Id).Updates(&bk)

	resData := struct {
		Code int         `json:"code"`
		Data interface{} `json:"data"`
		Msg  string      `json:"msg"`
	}{200, bk, "success"}
	c.JSON(200, resData)

}
