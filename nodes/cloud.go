package nodes

import (
	"fmt"
	"github.com/PlagueCat-Miao/goipfs-lab511/constdef"
	"github.com/PlagueCat-Miao/goipfs-lab511/dal/httppack"
	"github.com/PlagueCat-Miao/goipfs-lab511/dal/ipfs"
	"github.com/PlagueCat-Miao/goipfs-lab511/operate"
	"log"
)

func InitCloudServive() (int,error){
	//ipfs shell连接
	ipfsClient,err:=ipfs.InitIPFS()
	if err!=nil ||ipfsClient == nil {
		return -1,fmt.Errorf("[ipfs-init-err]: %v",err)
	}

	//本机状态更新
	operate.InitMyInfo(ipfsClient.DHash,constdef.CloudPort,constdef.CloudStatus,100,100)

	//主动连接 网关节点
	ipfs := ipfsClient.NewClient()
	peers ,err:=ipfs.BootstrapAdd([]string{fmt.Sprintf(constdef.LocalTestNode)})
	if err!=nil{
		return -1,fmt.Errorf("[ipfs-BootstrapAdd-err]: %v",err)
	}
	log.Printf("My ipfs peers: %+v",peers)

	url := fmt.Sprintf("http://%v:%v/login",constdef.LocalIP,constdef.GatewayPort)
	_,err =httppack.PostJson(url,operate.MyInfo)
	if err!=nil{
		return -1,fmt.Errorf("[httpPost-login-err]: %v, url:%v",err,url)
	}
	return constdef.CloudPort,nil
}

