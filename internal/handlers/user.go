package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lat1992/go-api/internal"
)

type GetUserRequest struct {
	UserId int64 `uri:"userId"`
}

type GetUserResponse struct {
	User internal.User `json:"user"`
}

func GetUser(us internal.UserService) func(c *gin.Context) {
	return func(c *gin.Context) {
		var id int64
		var data GetUserRequest
		if err := c.BindUri(&data); err != nil {
			warnEndCall(c, http.StatusBadRequest, "BindError", "File: parse query data error", "error", err.Error())
			return
		}
		if data.UserId > 0 {
			id = data.UserId
		} else {
			id = c.GetInt64("userId")
		}
		if id < 1 {
			warnEndCall(c, http.StatusForbidden, "IdShouldMoreThanZero", "User: id should more than zero")
			return
		}
		user, err := us.GetUser(id)
		if err != nil {
			errorEndCall(c, http.StatusBadRequest, "InternalError", "User: GetUser", err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"user": user})
	}
}

type GetUsersRequest struct {
	Rows int `uri:"rows"`
	Page int `uri:"page"`
}

type GetUsersResponse struct {
	Users []internal.User `json:"users"`
}

func GetUsers(us internal.UserService) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data GetUsersRequest
		if err := c.BindUri(&data); err != nil {
			warnEndCall(c, http.StatusBadRequest, "BindError", "File: parse query data error", "error", err.Error())
			return
		}
		if data.Rows < 1 {
			data.Rows = 10
		}
		if data.Page < 1 {
			data.Page = 1
		}
		users, err := us.GetUsers(data.Rows, data.Page)
		if err != nil {
			errorEndCall(c, http.StatusBadRequest, "InternalError", "User: GetUsers: %v", err)
			return
		}
		c.JSON(http.StatusOK, GetUsersResponse{
			Users: users,
		})
	}
}

type RegisterRequest struct {
	Id int64 `json:"id"`
}

func DeleteUser(us internal.UserService) func(c *gin.Context) {
	return func(c *gin.Context) {
		data := struct {
			Id int64 `json:"id"`
		}{}
		if err := c.BindJSON(&data); err != nil {
			warnEndCall(c, http.StatusBadRequest, "BindError", "File: parse query data error: %v", err)
			return
		}
		if data.Id < 1 {
			infoEndCall(c, http.StatusForbidden, "IdShouldMoreThanZero", "User: id should more than zero")
			return
		}
		if err := us.DeleteUser(data.Id); err != nil {
			errorEndCall(c, http.StatusBadRequest, "InternalError", "User: DeleteUser: %v", err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "Deleted"})
	}
}
