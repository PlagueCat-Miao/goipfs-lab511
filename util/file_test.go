package util

import (
	"github.com/PlagueCat-Miao/goipfs-lab511/constdef"
	"log"
	"testing"
)

func TestDirSize(t *testing.T) {
	t.Log("[DirSize]: TestDirSize 测试")
	sizeMB:=DirSize(constdef.IPFSPath)
    log.Printf("sizeMB:%v",sizeMB)
}