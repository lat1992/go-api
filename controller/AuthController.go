/*
 * File: X:\go-api\controller\AuthController.go
 * Created At: 2019-11-07 18:11:23
 * Created By: Mauhoi WU
 * 
 * Modified At: 2019-11-21 19:20:50
 * Modified By: Mauhoi WU
 */

package controller

import (
	"time"
	"github.com/gin-gonic/gin"
	"github.com/dchest/uniuri"
	"../model"
)

func GenerateToken(user_id int) string {
	token := uniuri.NewLen(20)
	model.DeleteOldToken(user_id)
	model.CreateToken(user_id, token, time.Now().AddDate(0, 1, 0))
	return token
}

func CheckToken(c *gin.Context) int {
	token := c.Param("token")
	if model.VerifyToken(token) == 0 {
		return 0
	}
	user_id := model.GetUserIdByToken(token)
	return user_id
}
