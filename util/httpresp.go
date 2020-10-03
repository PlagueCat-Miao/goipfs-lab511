package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseOK(c *gin.Context,msg map[string]interface{}) {
	msg["success"] = true
	c.JSON(http.StatusOK, gin.H(msg))
}


func ResponseError(c *gin.Context,err error) {
	c.JSON(http.StatusOK, gin.H{
		"err":     fmt.Sprintf("%v",err),
		"success": false,
	})
}

func ResponseBadRequest(c *gin.Context,err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"err":     fmt.Sprintf("%v",err),
		"success": false,
	})
}