/*
 * File: X:\go-api\controller\Controller.go
 * Created At: 2019-11-03 20:18:34
 * Created By: Mauhoi WU
 * 
 * Modified At: 2019-11-19 17:13:58
 * Modified By: Mauhoi WU
 */

package controller

import (
	"github.com/gin-gonic/gin"
)

func ErrorEndCall(c *gin.Context, httpStatus int, message string) {
	c.JSON(httpStatus, gin.H{ "error":  message })
}
