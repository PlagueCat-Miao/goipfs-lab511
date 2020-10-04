package service

import (
	"fmt"
	"github.com/PlagueCat-Miao/goipfs-lab511/constdef"
	"github.com/PlagueCat-Miao/goipfs-lab511/dal/httppack"
	"github.com/PlagueCat-Miao/goipfs-lab511/dal/ipfs"
	"github.com/PlagueCat-Miao/goipfs-lab511/operate"
	"github.com/PlagueCat-Miao/goipfs-lab511/util"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func IpfsSave(c *gin.Context) {
	if operate.MyInfo == nil {
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
	if addParams.FileSize > operate.MyInfo.Remain {
		util.ResponseError(c, fmt.Errorf("no spare space"))
		return
	}

	// ipfsget 此处交给子协程独立进行
	go func() {
		time.Sleep(1 * time.Second)
		err := ipfsCtrl.GetIPFSFile(addParams.FileHash, constdef.OutputFilePath+addParams.FileHash)
		if err != nil {
			log.Printf("[IpfsSave-GetIPFSFile-err]:%v", err)
		}
		var ReportParams = IPFSReportParams{
			OpType:       constdef.Add,
			OpInfo:       operate.MyInfo.MyClientInfo(),
			BackupNumber: addParams.BackupNumber,
			Fhash:        addParams.FileHash,
			Title:        addParams.Title,
			Size:         addParams.FileSize,
			Note:         "",
			Uploader:     addParams.Uploader,
		}
		url := fmt.Sprintf("http://%s:%v/ipfsreport",c.ClientIP(),constdef.GatewayPort)
		msg,err:= httppack.PostJson(url,ReportParams)
		_,msgErr:=util.ResponseParse(msg)
		if err != nil || msgErr!=nil{
			log.Printf("[IpfsSave-PostJson-err]:err=%v msgErr=%v ,url:%v", err,msgErr,url)
		}
	}()
	//应当每次调用时IpfsSave时 扫描一下存储空间更新一下，再相减文件送出
	operate.MyInfo.Remain = operate.MyInfo.Capacity - util.DirSize(constdef.IPFSPath)
	//返回给云端自己的状态
	msg := map[string]interface{}{
		"dhash":  operate.MyInfo.Dhash,
		"remain": operate.MyInfo.Remain-addParams.FileSize,
	}
	util.ResponseOK(c, msg)
}
