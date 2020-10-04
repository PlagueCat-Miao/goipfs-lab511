package service

import (
	"github.com/PlagueCat-Miao/goipfs-lab511/model"
	"github.com/PlagueCat-Miao/goipfs-lab511/operate"
	"github.com/PlagueCat-Miao/goipfs-lab511/util"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var userLoginInfo model.ClientInfo
	if err := c.ShouldBindJSON(&userLoginInfo); err != nil {
		util.ResponseBadRequest(c, err)
		return
	}
	clientIP := c.ClientIP()

	err := operate.ClientsMgr.AddUser(clientIP, &userLoginInfo)
	if err != nil {
		util.ResponseError(c, err)
	}
	msg := map[string]interface{}{"dhash": userLoginInfo.Dhash}
	util.ResponseOK(c, msg)
}
