package util

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseOK(c *gin.Context, msg map[string]interface{}) {
	msg["success"] = true
	c.JSON(http.StatusOK, gin.H(msg))
}

func ResponseError(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{
		"err":     fmt.Sprintf("%v", err),
		"success": false,
	})
}

func ResponseBadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"err":     fmt.Sprintf("%v", err),
		"success": false,
	})
}

func ResponseParse(msgBytes []byte) (*simplejson.Json, error) {
	msgJson, err := simplejson.NewJson(msgBytes)
	if err != nil {
		return nil, fmt.Errorf("[ResponseParse-simplejson-err]: %v", err)
	}
	isSucc, _ := msgJson.Get("success").Bool()
	if !isSucc {
		return nil, fmt.Errorf("[ResponseParse-err]:Response Err ,msg:%+v", string(msgBytes))
	}
	return msgJson, nil
}
