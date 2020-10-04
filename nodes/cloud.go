package nodes

import (
	"fmt"
	"github.com/PlagueCat-Miao/goipfs-lab511/constdef"
	"github.com/PlagueCat-Miao/goipfs-lab511/dal/httppack"
	"github.com/PlagueCat-Miao/goipfs-lab511/dal/ipfs"
	"github.com/PlagueCat-Miao/goipfs-lab511/operate"
	"github.com/PlagueCat-Miao/goipfs-lab511/util"
	"log"
)

func InitCloudServive() (int, error) {
	//ipfs shell连接
	ipfsClient, err := ipfs.InitIPFS()
	if err != nil || ipfsClient == nil {
		return -1, fmt.Errorf("[ipfs-init-err]: %v", err)
	}

	//本机状态更新
	operate.InitMyInfo(ipfsClient.DHash, constdef.CloudPort, constdef.CloudStatus, 4096)

	//主动连接 网关节点
	ipfs := ipfsClient.NewClient()
	//TODO M 云节点的网关List 可以配置成csv
	peers, err := ipfs.BootstrapAdd([]string{fmt.Sprintf(constdef.LocalTestNode)})
	if err != nil {
		return -1, fmt.Errorf("[ipfs-BootstrapAdd-err]: %v", err)
	}

	url := fmt.Sprintf("http://%v:%v/login", constdef.LocalIP, constdef.GatewayPort)
	msg, err := httppack.PostJson(url, &operate.MyInfo)
	if err != nil {
		return -1, fmt.Errorf("[httpPost-login-err]: %v, url:%v", err, url)
	}
	_, err = util.ResponseParse(msg)
	if err != nil {
		return -1, fmt.Errorf("[httpPost-login-err]: %v, url:%v", err, url)
	}
	log.Printf("My ipfs peers: %+v", peers)
	return constdef.CloudPort, nil
}
