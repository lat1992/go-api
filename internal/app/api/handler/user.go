package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetUser(c *gin.Context) {
	var id int64
	data := struct {
		UserId int64 `uri:"userId"`
	}{}
	if err := c.BindUri(&data); err != nil {
		h.warnEndCall(c, http.StatusBadRequest, "BindError", "File: parse query data error: %v", err)
		return
	}
	if data.UserId > 0 {
		id = data.UserId
	} else {
		id = c.GetInt64("userId")
	}
	if id < 1 {
		h.warnEndCall(c, http.StatusForbidden, "IdShouldMoreThanZero", "User: id should more than zero")
		return
	}
	user, err := h.controller.GetUser(id)
	if err != nil {
		h.errorEndCall(c, http.StatusBadRequest, "InternalError", "User: GetUser: %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *Handler) ListUsers(c *gin.Context) {
	data := struct {
		Rows int `uri:"rows"`
		Page int `uri:"page"`
	}{}
	if err := c.BindUri(&data); err != nil {
		h.warnEndCall(c, http.StatusBadRequest, "BindError", "File: parse query data error: %v", err)
		return
	}
	if data.Rows < 1 {
		data.Rows = 10
	}
	if data.Page < 1 {
		data.Page = 1
	}
	users, err := h.controller.GetUsers(data.Rows, data.Page)
	if err != nil {
		h.errorEndCall(c, http.StatusBadRequest, "InternalError", "User: GetUsers: %v", err)
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	data := struct {
		Id int64 `json:"id"`
	}{}
	if err := c.BindJSON(&data); err != nil {
		h.warnEndCall(c, http.StatusBadRequest, "BindError", "File: parse query data error: %v", err)
		return
	}
	if data.Id < 1 {
		h.warnEndCall(c, http.StatusForbidden, "IdShouldMoreThanZero", "User: id should more than zero")
		return
	}
	if err := h.controller.DeleteUser(data.Id); err != nil {
		h.errorEndCall(c, http.StatusBadRequest, "InternalError", "User: DeleteUser: %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Deleted"})
}
