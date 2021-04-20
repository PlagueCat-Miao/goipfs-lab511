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

type Friendlist struct {
	Gateway []string `yaml:"gateway"`
}

func InitCloudServive() (int, error) {
	//ipfs shell连接
	ipfsClient, err := ipfs.InitIPFS()
	if err != nil || ipfsClient == nil {
		return -1, fmt.Errorf("[ipfs-init-err]: %v", err)
	}

	//本机状态更新
	operate.InitMyInfo(ipfsClient.DHash, constdef.CloudPort, constdef.CloudStatus, 4096)

	//主动连接 网关节点
	var fl Friendlist
	err = util.LoadYaml(constdef.CloudConfPath, &fl, constdef.CloudFriendlistKey)
	if err != nil {
		return -1, fmt.Errorf("[LoadYaml-err]: %v", err)
	}

	var peersLink []string
	for _, Friendip := range fl.Gateway {
		url := fmt.Sprintf("http://%v:%v/login", Friendip, constdef.GatewayPort)
		msg, err := httppack.PostJson(url, &operate.MyInfo)
		if err != nil {
			log.Printf("[httpPost-login-err]: %v, url:%v", err, url)
			continue
		}
		msgJson, err := util.ResponseParse(msg)
		if err != nil {
			log.Printf("[httpPost-login-err]: %v, url:%v", err, url)
			continue
		}
		gDhash := msgJson.Get("gatewaydhash").MustString()
		peersLink = append(peersLink, fmt.Sprintf(constdef.IPFSNodeUrlFormat, Friendip, gDhash))
	}
	if len(peersLink) == 0 {
		return -1, fmt.Errorf("[LoadYaml-err]: no friend accept login")
	}

	ipfsC := ipfsClient.NewClient()
	peers, err := ipfsC.BootstrapAdd(peersLink)
	if err != nil {
		return -1, fmt.Errorf("[ipfs-BootstrapAdd-err]: %v", err)
	}
	log.Printf("My ipfs peers: %+v", peers)
	return constdef.CloudPort, nil
}
