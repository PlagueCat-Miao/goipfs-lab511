package ipfs

import (
	"bytes"
	"fmt"
	//"github.com/PlagueCat-Miao/GOIPFS-gateway/mmlog"

	"io/ioutil"
	"os"
)

type IIPFSCtrl interface {
	UploadIPFS(dataStr string) (string, error)
	CatIPFS(hash string) (string, error)
	UploadIPFSFile(filePath string) (string, error)
	GetIPFSFile(hash, filePath string) error
}

func NewIPFSCtrl() IIPFSCtrl {
	return &IPFSCtrl{}
}

type IPFSCtrl struct{}

//数据上传到ipfs
func (d *IPFSCtrl) UploadIPFS(dataStr string) (string, error) {
	ipfs := IPFSClient.NewClient()
	hash, err := ipfs.Add(bytes.NewBufferString(dataStr))
	//ipfs.BootstrapAdd("")
	if err != nil {
		return "", fmt.Errorf("[func Add] 上传ipfs时错误 err= %+v ", err)
	}
	return hash, nil
}
func (d *IPFSCtrl) CatIPFS(hash string) (string, error) {
	ipfs := IPFSClient.NewClient()
	read, err := ipfs.Cat(hash)
	if err != nil {
		return "", fmt.Errorf("[func Cat] 上传ipfs时错误 err= %+v ", err)
	}
	body, err := ioutil.ReadAll(read)
	return string(body), nil
}

func (d *IPFSCtrl) UploadIPFSFile(filePath string) (string, error) {
	ipfs := IPFSClient.NewClient()
	fileReader, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("[func os.Open] 读取文件出错 filePath:%v err= %+v ", filePath, err)
	}
	hash, err := ipfs.Add(fileReader)
	if err != nil {
		return "", fmt.Errorf("[func Add] 上传ipfs时错误 err= %+v ", err)
	}
	return hash, nil
}
func (d *IPFSCtrl) GetIPFSFile(hash, filePath string) error {
	ipfs := IPFSClient.NewClient()
	err := ipfs.Get(hash, filePath)
	if err != nil {
		return fmt.Errorf("[func Get] 下载ipfs时错误 err= %+v ", err)
	}
	return nil
}
