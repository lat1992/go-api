/*
 * File: X:\enigm-mvc\controller\AuthController.go
 * Created At: 2019-11-07 18:11:23
 * Created By: Mauhoi WU
 * 
 * Modified At: 2019-11-18 20:49:45
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
	m := model.NewModel()
	defer m.Destroy()
	token := uniuri.NewLen(20)
	m.DeleteOldToken(user_id)
	m.CreateToken(user_id, token, time.Now().AddDate(0, 1, 0))
	return token
}

func CheckToken(c *gin.Context) int {
	token := c.Param("token")
	m := model.NewModel()
	defer m.Destroy()
	if m.VerifyToken(token) == 0 {
		return 0
	}
	user_id := m.GetUserIdByToken(token)
	return user_id
}
