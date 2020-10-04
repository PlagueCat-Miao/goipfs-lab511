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
	"sync"
)

type IPFSAddParams struct {
	FileHash     string `form:"filehash" json:"filehash"`
	FileSize     int64  `form:"filesize" json:"filesize"`
	Title        string `form:"title" json:"title"`
	Uploader     string ` form:"uploader" json:"uploader"`
	BackupAmount int    `form:"backupamount" json:"backupamount"`
	BackupNumber int    `form:"backupnumber" json:"backupnumber"`
}

func IpfsAdd(c *gin.Context) {
	var addParams IPFSAddParams
	if err := c.ShouldBindJSON(&addParams); err != nil {
		util.ResponseBadRequest(c, err)
		return
	}

	operate.ClientsMgr.Lock.Lock()
	sort.Slice(operate.ClientsMgr.CloudClientList, func(i, j int) bool { //降序 （大在前）
		return operate.ClientsMgr.CloudClientList[i].Remain > operate.ClientsMgr.CloudClientList[j].Remain
	})
	operate.ClientsMgr.Lock.Unlock()
	var CloudSaveList []*model.ClientInfo
	ListLen := len(operate.ClientsMgr.CloudClientList)
	for i := 0; i < constdef.SaveListLesslength && i < addParams.BackupAmount && i < ListLen; i++ {
		if operate.ClientsMgr.CloudClientList[i].Remain > addParams.FileSize {
			CloudSaveList = append(CloudSaveList, operate.ClientsMgr.CloudClientList[i])
		} else {
			break
		}
	}
	if len(CloudSaveList) == 0{  //很可能文件太大
		err:= fmt.Errorf("invail file, addParams:%+v",addParams)
		log.Printf("[IpfsAdd-err]:%v",err)
		util.ResponseBadRequest(c, err)
		return
	}
	//命令云节点存储
	failList := parallelCloudSave(CloudSaveList, addParams)

	//返回给用户通信状态
	msg := map[string]interface{}{
		"SaveList": CloudSaveList,
		"failList": failList,
		"SaveNum":  len(CloudSaveList) - len(failList),
	}
	util.ResponseOK(c, msg)
}

func parallelCloudSave(CloudSaveList []*model.ClientInfo, tempParams IPFSAddParams) []*model.ClientInfo {
	var failList []*model.ClientInfo

	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}
	for i, targetCloud := range CloudSaveList {
		targetParams := tempParams
		wg.Add(1)
		go func(targetCloud *model.ClientInfo, targetParams IPFSAddParams) {
			defer wg.Done()
			url := fmt.Sprintf("http://%v:%v/ipfssave", targetCloud.Ip, targetCloud.Port)
			targetParams.BackupNumber = i
			cloudinfoByte, err := httppack.PostJson(url, targetParams)
			if err != nil {
				mutex.Lock()
				failList = append(failList, targetCloud)
				mutex.Unlock()
				log.Printf("[PostJson-err]:%v,url:%v targetCloud:%v addParams:%+v", err, url, targetCloud, targetParams)
				return
			}
			// 存储成功 解析返回参数 更新状态
			cloudinfoJson, err := util.ResponseParse(cloudinfoByte)
			if err != nil {
				log.Printf("[ResponseParse-err]:%v,url:%v targetCloud:%v addParams:%+v", err, url, targetCloud, targetParams)
				return
			}
			targetCloud.Remain = cloudinfoJson.Get("remain").MustInt64()
			if targetCloud.Remain < 50 {
				log.Printf("[Cloud.Remain-warning]: no spare space, url:%v targetCloud:%+v Remain:%v", url, targetCloud, targetCloud.Remain)
			}

		}(targetCloud, targetParams)

	}
	wg.Wait()

	return failList
}
