package ipfs

import (
	"fmt"
	"github.com/PlagueCat-Miao/GOIPFS-gateway/constdef"
	"github.com/PlagueCat-Miao/GOIPFS-gateway/mmlog"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("[ipfs单测 开始]")
	err := InitIPFS()
	if err != nil {
		mmlog.ErrLog(err, constdef.Show, constdef.ParamNil)
		return
	}
	result := m.Run()
	fmt.Println("[ipfs单测 结束]")
	os.Exit(result)
}
