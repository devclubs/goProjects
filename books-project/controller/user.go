package controller

import (
	"books_project/dao/mysql"
	"books_project/model"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RegisterHandler(c *gin.Context) {
	var usr model.User
	err := c.BindJSON(&usr)
	fmt.Println(err, "++++")
	if err != nil {
		resData := struct {
			Code int         `json:"code"`
			Data interface{} `json:"data"`
			Msg  string      `json:"msg"`
		}{400, "", err.Error()}
		c.JSON(200, resData)

		fmt.Println(err)
		return
	}

	mysql.DB.Create(&usr)
	resData := struct {
		Code int         `json:"code"`
		Data interface{} `json:"data"`
		Msg  string      `json:"msg"`
	}{400, usr, ""}
	c.JSON(200, resData)

}

func LoginHandler(c *gin.Context) {
	var usr model.User
	err := c.BindJSON(&usr)
	if err != nil {
		resData := struct {
			Code int         `json:"code"`
			Data interface{} `json:"data"`
			Msg  string      `json:"msg"`
		}{400, "", err.Error()}
		c.JSON(200, resData)
		return
	}
	result := mysql.DB.Where("username=? AND password=?", usr.Username, usr.Password).First(&usr)
	if result.Error != nil {
		resData := struct {
			Code int         `json:"code"`
			Data interface{} `json:"data"`
			Msg  string      `json:"msg"`
		}{400, "", result.Error.Error()}
		c.JSON(400, resData)
		return
	}

	token := uuid.New().String()
	mysql.DB.Model(usr).Update("token", token)

	resData := struct {
		Code  int         `json:"code"`
		Token interface{} `json:"token"`
		Msg   string      `json:"msg"`
	}{200, token, "succes"}
	c.JSON(200, resData)

}
