package ipfs

import (
	"fmt"
	"github.com/PlagueCat-Miao/goipfs-lab511/constdef"
	shell "github.com/ipfs/go-ipfs-api"
	"log"
)

type IPFS_Shell struct {
	url  string
	DHash string
}

var IPFSClient *IPFS_Shell

func InitIPFS() (*IPFS_Shell,error) {
	IPFSClient = initIPFSClient()
	err := IPFSClient.Conn()
	if err != nil {
		return  nil,fmt.Errorf("[Conn-err]: %v",err)
	}

	return IPFSClient,nil
}

func initIPFSClient() *IPFS_Shell {
	return &IPFS_Shell{
		url:  constdef.IPFSHost,
		DHash: constdef.IPFSMyId,
	}
}
func (sh *IPFS_Shell) Conn() error {
	client := shell.NewShell(sh.url)
	myInfo, err := client.ID(sh.DHash)
	if err != nil {
		return fmt.Errorf("[NewShell-err] url:%+v id:%+v err= %+v ", sh.url, sh.DHash, err)
	}
	log.Printf("ipfs: %+v 初始化完成	...", myInfo.ID)
	return nil
}
func (sh *IPFS_Shell) NewClient() *shell.Shell {
	return shell.NewShell(sh.url)
}
