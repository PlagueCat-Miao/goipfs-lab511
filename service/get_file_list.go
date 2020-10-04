package service

import (
	"fmt"
	"github.com/PlagueCat-Miao/goipfs-lab511/dal/db"
	"github.com/PlagueCat-Miao/goipfs-lab511/util"
	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
)

type GetFileListParams struct {
	FileHash     string `form:"filehash" json:"filehash"`
	Title        string `form:"title" json:"title"`
	Uploader     string ` form:"uploader" json:"uploader"`
	CreateTime   string `form:"createtime" json:"createtime"` // xx;yy
	UpdateTime   string `form:"updatetime" json:"updatetime"` // xx;yy
	Limit  int64 `form:"limit" json:"limit"` // 0代表没有上限制
	Offset int64 `form:"offset" json:"offset"`
}

func GetFileList(c *gin.Context) {
	var params GetFileListParams
	if err := c.ShouldBindJSON(&params); err != nil {
		util.ResponseBadRequest(c, err)
		return
	}
	fileDb := db.NewIPFSFileInfoDB()
	queryFactor :=map[string]string{}
	if params.FileHash !=""{
		queryFactor["filehash"] =  params.FileHash
	}
	if params.Title !="" {
		queryFactor["title"] = params.Title
	}
	if params.Uploader !=""{
		queryFactor["uploader"] =  params.Uploader
	}
	if params.CreateTime !=""{
		queryFactor["create_time"] =  params.CreateTime
	}
	if params.UpdateTime !=""{
		queryFactor["update_time"] =  params.UpdateTime
	}

	ansList,total,err:=fileDb.GetInfoList(params.Offset,params.Limit,queryFactor)
    if err!=nil{
		util.ResponseError(c, fmt.Errorf("[db-GetInfoList-err]:%v",err))
		return
	}
	ansByte,err:=json.Marshal(ansList)
	if err!=nil{
		util.ResponseError(c, fmt.Errorf("[db-Marshal-err]:%v",err))
		return
	}

	//返回给用户通信状态
	msg := map[string]interface{}{
		"ansList": string(ansByte),
		"total": total,
		"offset":params.Offset,
	}
	util.ResponseOK(c, msg)

}