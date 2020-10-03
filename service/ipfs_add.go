package service

import (
	"fmt"
	"github.com/PlagueCat-Miao/goipfs-lab511/constdef"
	"github.com/PlagueCat-Miao/goipfs-lab511/dal/httppack"
	"github.com/PlagueCat-Miao/goipfs-lab511/model"
	"github.com/PlagueCat-Miao/goipfs-lab511/operate"
	"github.com/PlagueCat-Miao/goipfs-lab511/util"
	"github.com/gin-gonic/gin"
	"log"

	"sort"
)



type IPFSAddParams struct {
	FileHash    string `form:"filehash" json:"filehash"`
	FileSize    int64 `form:"filesize" json:"filesize"`
	BackupAmount  int `form:"backupamount" json:"backupamount"`
	BackupNumber int `form:"backupnumber" json:"backupnumber"`
}

func IpfsAdd(c *gin.Context) {
	var addParams IPFSAddParams
	if err := c.ShouldBindJSON(&addParams); err != nil {
		util.ResponseBadRequest(c, err)
		return
	}

	sort.Slice(operate.ClientsMgr.CloudClientList,func (i,j int)bool{ //降序 （大在前）
		return operate.ClientsMgr.CloudClientList[i].Remain > operate.ClientsMgr.CloudClientList[j].Remain
	})
	var CloudSaveList []*model.ClientInfo
	for i:=0;i<constdef.SaveListLesslength && i< addParams.BackupAmount;i++{
		if  operate.ClientsMgr.CloudClientList[i].Remain > addParams.FileSize {
			CloudSaveList = append(CloudSaveList,operate.ClientsMgr.CloudClientList[i])
		}else{
			break
		}
	}
	//命令云节点存储
	var failList []*model.ClientInfo
    for i, targetCloud :=range CloudSaveList{
		url := fmt.Sprintf("http://%v:%v/ipfssave", targetCloud.Ip, targetCloud.Port)
		addParams.BackupNumber = i
		_,err:= httppack.PostJson(url,addParams)
		if err != nil{
			failList =append(failList,targetCloud)
			log.Printf("[PostJson-err]:%v,url:%v targetCloud:%v addParams:%+v",err,url,targetCloud,addParams)
		}
	}
	//返回给用户通信状态
	msg :=map[string]interface{}{
		"SaveList" : CloudSaveList,
		"failList" :  failList,
		"SaveNum" : len(CloudSaveList)-len(failList),
	}
	util.ResponseOK(c, msg)
}

