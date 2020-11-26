package service

import (
	"fmt"
	"github.com/PlagueCat-Miao/goipfs-lab511/constdef"
	"github.com/PlagueCat-Miao/goipfs-lab511/dal/httppack"
	"github.com/PlagueCat-Miao/goipfs-lab511/util"
	"github.com/gin-gonic/gin"
	"log"
)

type RtmpCtrlParams struct {
	//Dhash string `form:"dhash" json:"dhash"`
	TagIp string `form:"ip" json:"ip"`
	Port int `form:"port" json:"port"`
	//Room string `form:"room" json:"room"`
	Enable bool  `form:"enable" json:"enable"`
}

func RtmpCtrl(c *gin.Context) {
	var rtmpCtrlParams RtmpCtrlParams
	if err := c.ShouldBindJSON(&rtmpCtrlParams); err != nil {
		util.ResponseBadRequest(c, err)
		return
	}

	var RtmpPushParams = RtmpPushParams{
		TagIp:  "127.0.0.1",
		Port:   constdef.RtmpDefaultPort,
		Room:   constdef.LiveRoom,
		Enable: rtmpCtrlParams.Enable,
	}
	url := fmt.Sprintf("http://%v:%v/rtmppush", rtmpCtrlParams.TagIp, rtmpCtrlParams.Port)
	msg,err:= httppack.PostJson(url,RtmpPushParams)
	_,msgErr:=util.ResponseParse(msg)
	if err != nil || msgErr!=nil{
		log.Printf("[RtmpStart-PostJson-err]:err=%v msgErr=%v ,url:%v", err,msgErr,url)
		util.ResponseBadRequest(c, err)
		return
	}
	state := "start"
	if !rtmpCtrlParams.Enable{
		state = "stop"
	}
	//返回给用户通信状态
	succMsg := map[string]interface{}{
		"State": state,
		"Ip" : rtmpCtrlParams.TagIp,
	}
	util.ResponseOK(c, succMsg)
}