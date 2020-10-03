package service

import (
	"fmt"
	"github.com/PlagueCat-Miao/goipfs-lab511/constdef"
	"github.com/PlagueCat-Miao/goipfs-lab511/dal/ipfs"
	"github.com/PlagueCat-Miao/goipfs-lab511/operate"
	"github.com/PlagueCat-Miao/goipfs-lab511/util"
	"github.com/gin-gonic/gin"
)


func IpfsSave(c *gin.Context) {
	if operate.MyInfo == nil{
		util.ResponseBadRequest(c, fmt.Errorf("[MyInfo-err]:MyInfo is nil"))
		return
	}

	var addParams IPFSAddParams
	if err := c.ShouldBindJSON(&addParams); err != nil {
		util.ResponseBadRequest(c, err)
		return
	}
	ipfsCtrl := ipfs.NewIPFSCtrl()

    // 剩余空间判断
	if addParams.FileSize > operate.MyInfo.Remain{
		util.ResponseBadRequest(c, fmt.Errorf("no spare space"))
		return
	}
	// ipfsget
	err := ipfsCtrl.GetIPFSFile(addParams.FileHash,constdef.OutputFilePath + addParams.FileHash)
	if err !=nil{
		util.ResponseBadRequest(c, err)
		return
	}

	//返回给云端自己的状态
	msg :=map[string]interface{}{
		"dhash" : operate.MyInfo.Dhash,
		"remain" : operate.MyInfo.Remain,
	}
	util.ResponseOK(c, msg)
}

