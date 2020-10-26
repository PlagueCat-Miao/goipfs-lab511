package service

import (
	"github.com/PlagueCat-Miao/goipfs-lab511/dal/rtmpffmpeg"
	"github.com/PlagueCat-Miao/goipfs-lab511/util"
	"github.com/gin-gonic/gin"
	"log"
)

type RtmpPushParams struct {
	TagIp string `form:"ip" json:"ip"`
	Port int `form:"port" json:"port"`
	Room string `form:"room" json:"room"`
	Enable bool  `form:"enable" json:"enable"`
}

func RtmpPush(c *gin.Context) {
	var rtmpPushParams RtmpPushParams
	if err := c.ShouldBindJSON(&rtmpPushParams); err != nil {
		util.ResponseBadRequest(c, err)
		return
	}
	RtmpFfmpegCtrl := rtmpffmpeg.NewRmptFfmpeg()
	var err error
	var url string
	if rtmpPushParams.Enable {
		if rtmpPushParams.TagIp == "127.0.0.1" || rtmpPushParams.TagIp == "localhost" {
			rtmpPushParams.TagIp =  c.ClientIP()
		}
		url, err =RtmpFfmpegCtrl.RtmpPushExec(rtmpPushParams.TagIp, rtmpPushParams.Port, rtmpPushParams.Room)
		if err != nil{
			log.Printf("[RtmpPushExec] err = %v",err)
		}else{
			log.Printf("[RtmpPushExec] Push %s Start",url)
		}
	}else{
		err =RtmpFfmpegCtrl.Stop()
		if err != nil{
			log.Printf("[RtmpPushExec] err = %v",err)
		}else{
			log.Printf("[RtmpPushExec] Push Stop")
		}
	}

	if err != nil{
		util.ResponseError(c, err)
		return
	}
	//返回给用户通信状态
	msg := map[string]interface{}{
		"url" : url,
		"enable": rtmpPushParams.Enable,
	}
	util.ResponseOK(c, msg)
}