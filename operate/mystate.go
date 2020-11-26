package operate

import (
	"fmt"
	"github.com/PlagueCat-Miao/goipfs-lab511/constdef"
	"github.com/PlagueCat-Miao/goipfs-lab511/model"
	"github.com/PlagueCat-Miao/goipfs-lab511/util"
	"sync"
	"time"
)

var MyInfo *MyInfoStruct

type MyInfoStruct struct {
	model.ClientInfo
	Lock sync.Mutex
}

func InitMyInfo(dhash string, port int, status constdef.UserStatus, capacity int64) *MyInfoStruct {
	mypath,_:=util.ShowMyHomePath()
	myInfo := &MyInfoStruct{
		ClientInfo: model.ClientInfo{
			Dhash:    dhash,
			Ip:       constdef.LocalIP, //此IP没有太大意义,请服务端自行判断IP
			Port:     port,
			Status:   status,
			Capacity: capacity,
			Remain:  capacity - util.DirSize(fmt.Sprintf(constdef.IPFSPath,mypath)),
		},
		Lock: sync.Mutex{},
	}
	MyInfo = myInfo
	return myInfo
}
func (m *MyInfoStruct) MyClientInfo() model.ClientInfo{
	return model.ClientInfo{
		Dhash:            m.Dhash,
		Status:           m.Status,
		Ip:               m.Ip,
		Port:             m.Port,
		Capacity:         m.Capacity,
		Remain:           m.Remain,
		LastPingPongTime: time.Time{},
	}
}
