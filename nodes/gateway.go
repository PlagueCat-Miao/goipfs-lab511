package nodes

import (
	"fmt"
	"github.com/PlagueCat-Miao/goipfs-lab511/constdef"
	"github.com/PlagueCat-Miao/goipfs-lab511/dal/ipfs"
	"github.com/PlagueCat-Miao/goipfs-lab511/operate"
)

func InitGatewayServive() (int,error){
	//ipfs shell连接
	ipfsClient,err:=ipfs.InitIPFS()
	if err!=nil ||ipfsClient == nil {
		return -1,fmt.Errorf("[ipfs-err]: %v",err)
	}
	//本机状态更新
	operate.InitMyInfo(ipfsClient.DHash,constdef.GatewayPort,constdef.GatewayStatus,100,100)
	return constdef.GatewayPort,nil
}