package util

import (
	"fmt"
	"github.com/PlagueCat-Miao/goipfs-lab511/constdef"
	"log"
	"testing"
)

func TestDirSize(t *testing.T) {
	t.Log("[DirSize]: TestDirSize 测试")
	mypath,_:=ShowMyHomePath()
	sizeMB:=DirSize(fmt.Sprintf(constdef.IPFSPath,mypath))
    log.Printf("sizeMB:%v",sizeMB)
}