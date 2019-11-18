/*
 * File: X:\enigm-mvc\controller\UserController.go
 * Created At: 2019-11-12 18:10:25
 * Created By: Mauhoi WU
 * 
 * Modified At: 2019-11-15 22:40:09
 * Modified By: Mauhoi WU
 */

package controller

import (
	"strconv"
	"golang.org/x/crypto/bcrypt"
	"github.com/gin-gonic/gin"
	"../model"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes), err
}

func comparePassword(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}

func CreateUser(c *gin.Context) {
	data := struct {
		Username string `json:"username"`
		Email string `json:"email"`
		Password string `json:"password"`
		FullName string `json:"full_name"`
		Country string `json:"country"`
		Telephone string `json:"telephone"`
	}{}
	err := c.BindJSON(&data)
	if err != nil {
		ErrorEndCall(c, 400, "Server error")
		return
	}
	if data.Username == "" || data.Email == "" || data.Password == "" || data.FullName == "" {
		ErrorEndCall(c, 403, "Username, Email, Password, FullName parameter not found")
		return
	}
	m := model.NewModel()
	defer m.Destroy()
	if m.VerifyUsername(data.Username) != 0 || m.VerifyEmail(data.Email) != 0 {
		c.JSON(203, gin.H{ "message": "Username or Email used" })
		return
	}
	hash, _ := hashPassword(data.Password)
	user_id := m.AddUser(data.Username, data.Email, hash, data.FullName, data.Country, data.Telephone)
	token := GenerateToken(user_id)
	c.JSON(201, gin.H{ "token": token })
}

func LoginUser(c *gin.Context) {
	data := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	err := c.BindJSON(&data)
	if err != nil {
		ErrorEndCall(c, 400, "Server error")
		return
	}
	if data.Username == "" || data.Password == "" {
		ErrorEndCall(c, 403, "username or password parameter not found")
		return
	}
	m := model.NewModel()
	defer m.Destroy()
	if m.VerifyUsername(data.Username) == 0 {
		c.JSON(203, gin.H{ "message": "Username is incorrect" })
		return
	}
	user_id, password := m.GetUserIdAndPasswordByUsername(data.Username)
	if comparePassword(password, data.Password) != nil {
		c.JSON(203, gin.H{ "message": "Password is incorrect" })
		return
	}
	token := GenerateToken(user_id)
	c.JSON(200, gin.H{ "token": token })
}

func GetUserList(c *gin.Context) {
	rows_str := c.Param("rows")
	page_str := c.Param("page")
	if rows_str == "" || page_str == "" {
		ErrorEndCall(c, 400, "rows or page parameter not found")
		return
	}
	rows, err := strconv.Atoi(rows_str)
	if err != nil || rows < 1 {
		rows = 10
	}
	page, err := strconv.Atoi(page_str)
	if err != nil || rows < 1 {
		page = 1
	}
	m := model.NewModel()
	defer m.Destroy()
	users := m.GetUsers(rows, page)
	c.JSON(200, users)
}

func GetUser(c *gin.Context) {
	id_str := c.Param("id")
	if id_str == "" {
		ErrorEndCall(c, 400, "id parameter not found")
		return
	}
	id, err := strconv.Atoi(id_str)
	if err != nil || id < 1 {
		ErrorEndCall(c, 400, "id parameter should be numeric and bigger than 0")
		return
	}
	m := model.NewModel()
	defer m.Destroy()
	m.GetUserById(id)
}

func ChangeUserPassword(c *gin.Context) {

}

func ChangeUserInformation(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
