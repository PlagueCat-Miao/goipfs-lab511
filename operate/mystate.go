package operate

import (
	"github.com/PlagueCat-Miao/goipfs-lab511/constdef"
	"github.com/PlagueCat-Miao/goipfs-lab511/model"
	"sync"
)

var MyInfo *MyInfoStruct

type MyInfoStruct struct{
	model.ClientInfo
	Lock       sync.Mutex
}

func InitMyInfo(dhash string,port int, status constdef.UserStatus ,capacity,remain int64) *MyInfoStruct {
	myInfo :=& MyInfoStruct{
		ClientInfo: model.ClientInfo{
			Dhash : dhash,
			Port: port,
			Status: status,
			Capacity  : capacity,
			Remain    :  remain,
		},
		Lock:sync.Mutex{},
	}
	MyInfo = myInfo
	return myInfo
}
