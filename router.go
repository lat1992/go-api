/*
 * File: X:\enigm-mvc\router.go
 * Created At: 2019-11-07 18:29:31
 * Created By: Mauhoi WU
 * 
 * Modified At: 2019-11-18 20:51:45
 * Modified By: Mauhoi WU
 */

package main

import (
	"github.com/gin-gonic/gin"
	"./controller"
)

func GetRouter(router *gin.Engine) {
	router.POST("/user/create", controller.CreateUser)
	router.POST("/user/login", controller.LoginUser)
}
