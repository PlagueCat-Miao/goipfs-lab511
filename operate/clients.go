package operate

import (
	"encoding/csv"
	"fmt"
	"github.com/PlagueCat-Miao/goipfs-lab511/constdef"
	"github.com/PlagueCat-Miao/goipfs-lab511/model"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

var ClientsMgr *ClientsManager

func init() {
	ClientsMgr = NewClientsManager()
}

type ClientsManager struct {
	ClientList        map[string]*model.ClientInfo
	EdgeClientList    []*model.ClientInfo
	GatewayClientList []*model.ClientInfo
	CloudClientList   []*model.ClientInfo

	Lock sync.Mutex
}

func NewClientsManager() *ClientsManager {
	return &ClientsManager{
		ClientList: make(map[string]*model.ClientInfo),
		Lock:       sync.Mutex{},
	}
}

//新用户登记
func (c *ClientsManager) AddUser(ip string, info *model.ClientInfo) error {
	if c.ClientList == nil {
		return fmt.Errorf("UserCheck Forget initialization")
	}
	if info == nil || info.Ip == "" {
		return fmt.Errorf("ClientInfo is empty")
	}
	c.Lock.Lock()
	c.ClientList[fmt.Sprintf("%v:%v", ip, info.Port)] = info
	Status := info.Status
	if constdef.UserStatus(Status) == constdef.EdgeStatus {
		c.EdgeClientList = append(c.EdgeClientList, info)
	} else if constdef.UserStatus(Status) == constdef.CloudStatus {
		c.CloudClientList = append(c.CloudClientList, info)
	} else if constdef.UserStatus(Status) == constdef.GatewayStatus {
		c.GatewayClientList = append(c.GatewayClientList, info)
	}
	c.Lock.Unlock()
	return nil
}

//网关下线时，保存数据
func (c *ClientsManager) SaveUserCSV() {
	fileName := constdef.UserCSVName
	nfs, err := os.OpenFile(fileName, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("can not create file, err is %+v", err)
	}
	defer nfs.Close()
	nfs.Seek(0, io.SeekEnd)
	w := csv.NewWriter(nfs)
	//设置属性
	w.Comma = ','
	w.UseCRLF = true
	for key, client := range c.ClientList {
		row := []string{key, client.Dhash, string(client.Status), client.Ip, strconv.Itoa(client.Port), strconv.FormatInt(client.Capacity, 10), strconv.FormatInt(client.Remain, 10)}
		err = w.Write(row)
	}
	if err != nil {
		log.Fatalf("can not write, err is %+v", err)
	}
	//这里必须刷新，才能将数据写入文件。
	w.Flush()
}

//网关上线时，读取上次数据
func (c *ClientsManager) LoadUserCSV() {
	fileName := constdef.UserCSVName
	fs, err := os.Open(fileName)
	if err != nil {
		log.Printf("can not open the file, err is %+v", err)
	}
	defer fs.Close()

	fs1, _ := os.Open(fileName)
	r1 := csv.NewReader(fs1)
	content, err := r1.ReadAll()
	if err != nil {
		log.Fatalf("can not readall, err is %+v", err)
	}
	c.Lock.Lock()

	for _, row := range content {
		key := row[0]
		Dhash := row[1]
		Status, _ := strconv.Atoi(row[2])
		Ip := row[3]
		Port, _ := strconv.Atoi(row[4])
		Capacity, _ := strconv.ParseInt(row[5], 10, 64)
		Remain, _ := strconv.ParseInt(row[6], 10, 64)
		info := &model.ClientInfo{
			Dhash:            Dhash,
			Status:           constdef.UserStatus(Status),
			Ip:               Ip,
			Port:             Port,
			Capacity:         Capacity,
			Remain:           Remain,
			LastPingPongTime: time.Now(),
		}
		c.ClientList[key] = info
		if constdef.UserStatus(Status) == constdef.EdgeStatus {
			c.EdgeClientList = append(c.EdgeClientList, info)
		} else if constdef.UserStatus(Status) == constdef.CloudStatus {
			c.CloudClientList = append(c.CloudClientList, info)
		} else if constdef.UserStatus(Status) == constdef.GatewayStatus {
			c.GatewayClientList = append(c.GatewayClientList, info)
		}
	}
	log.Printf("Gateway ClientList %+v", c.ClientList)

	c.Lock.Unlock()

}
