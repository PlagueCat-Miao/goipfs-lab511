package ipfs

import (
	"fmt"
	"github.com/PlagueCat-Miao/goipfs-lab511/constdef"
	"github.com/PlagueCat-Miao/goipfs-lab511/util"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("[ipfs单测 开始]")
	_, err := InitIPFS()
	if err != nil {
		util.ErrLog(err, constdef.Show, constdef.ParamNil)
		return
	}
	result := m.Run()
	fmt.Println("[ipfs单测 结束]")
	os.Exit(result)
}
