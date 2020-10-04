package service

import (
	"fmt"
	"github.com/PlagueCat-Miao/goipfs-lab511/constdef"
	"github.com/PlagueCat-Miao/goipfs-lab511/dal/db"
	"github.com/PlagueCat-Miao/goipfs-lab511/model"
	"github.com/PlagueCat-Miao/goipfs-lab511/util"
	"github.com/gin-gonic/gin"
	"log"
)

type IPFSReportParams struct {
	OpType       constdef.OpTypeEvent `form:"optype" json:"optype"`
	OpInfo       model.ClientInfo     `form:"opinfo" json:"opinfo"`
	BackupNumber int                  `form:"backupnumber" json:"backupnumber"`
	Fhash        string               `form:"fhash" json:"fhash"`
	Title        string               `form:"title" json:"title"`
	Size         int64                `form:"size" json:"size"`
	Note         string               `form:"note" json:"note"`
	Uploader     string               ` form:"uploader" json:"uploader"`
}

//目前只做存储的上报
func IpfsReport(c *gin.Context) {
	var reportParams IPFSReportParams
	if err := c.ShouldBindJSON(&reportParams); err != nil {
		util.ResponseBadRequest(c, err)
		return
	}
	reportParams.OpInfo.Ip=c.ClientIP()
	var err error
	switch reportParams.OpType {
	case constdef.Add:
		err = DBAdd(reportParams)
	default:
		log.Printf("[warning] OpType:%+v", reportParams.OpType)
	}
	if err != nil {
		util.ResponseError(c, err)
		return
	}
	//返回给用户通信状态
	msg := map[string]interface{}{}
	util.ResponseOK(c, msg)
}

func DBAdd(reportParams IPFSReportParams) error {
	fileDb := db.NewIPFSFileInfoDB()
	fileInfo := &model.FileInfo{
		Fhash:         reportParams.Fhash,
		Title:         reportParams.Title,
		Uploader:      reportParams.Uploader,
		Size:          reportParams.Size,
		AuthorityCode: 7,
		Note:          reportParams.Note,
	}
	Ownermap := map[string] model.ClientInfo{
		reportParams.OpInfo.Dhash : reportParams.OpInfo,
	}
	fileInfo.OwnersMarshal(Ownermap)
	err := fileDb.OwnerIncrByFhash(reportParams.Fhash, fileInfo)
	if err != nil {
		return fmt.Errorf("[OwnerIncrByFhash-err]:%v", err)
	}
	return nil
}
