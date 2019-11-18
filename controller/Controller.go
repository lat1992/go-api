/*
 * File: X:\enigm-mvc\controller\Controller.go
 * Created At: 2019-11-03 20:18:34
 * Created By: Mauhoi WU
 * 
 * Modified At: 2019-11-18 20:49:48
 * Modified By: Mauhoi WU
 */

package controller

import (
	"github.com/gin-gonic/gin"
)

func ErrorEndCall(c *gin.Context, httpStatus int, message string) {
	c.JSON(httpStatus, gin.H{ "error":  message })
}
