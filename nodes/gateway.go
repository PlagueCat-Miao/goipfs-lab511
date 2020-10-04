package nodes

import (
	"fmt"
	"github.com/PlagueCat-Miao/goipfs-lab511/constdef"
	"github.com/PlagueCat-Miao/goipfs-lab511/dal/db"
	"github.com/PlagueCat-Miao/goipfs-lab511/dal/ipfs"
	"github.com/PlagueCat-Miao/goipfs-lab511/operate"
)

func InitGatewayServive() (int, error) {
	//ipfs shell连接
	ipfsClient, err := ipfs.InitIPFS()
	if err != nil || ipfsClient == nil {
		return -1, fmt.Errorf("[ipfs-err]: %v", err)
	}
	err = db.InitDataBase()
	if err != nil {
		return -1, fmt.Errorf("[db-err]: %v", err)
	}
	//本机状态更新
	operate.InitMyInfo(ipfsClient.DHash, constdef.GatewayPort, constdef.GatewayStatus, 4096)
	//加载上一次关闭时的连接用户
	operate.ClientsMgr.LoadUserCSV()
	return constdef.GatewayPort, nil
}
