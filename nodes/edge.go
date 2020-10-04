package nodes

import (
	"encoding/json"
	"fmt"
	"github.com/PlagueCat-Miao/goipfs-lab511/constdef"
	"github.com/PlagueCat-Miao/goipfs-lab511/dal/httppack"
	"github.com/PlagueCat-Miao/goipfs-lab511/dal/ipfs"
	"github.com/PlagueCat-Miao/goipfs-lab511/model"
	"github.com/PlagueCat-Miao/goipfs-lab511/operate"
	"github.com/PlagueCat-Miao/goipfs-lab511/service"
	"github.com/PlagueCat-Miao/goipfs-lab511/util"
	"log"
)

func InitEdgeServive() (int, error) {
	//ipfs shell连接
	ipfsClient, err := ipfs.InitIPFS()
	if err != nil || ipfsClient == nil {
		return -1, fmt.Errorf("[ipfs-init-err]: %v", err)
	}

	//本机状态更新
	operate.InitMyInfo(ipfsClient.DHash, constdef.EdgePort, constdef.EdgeStatus, 4096)

	return constdef.EdgePort, nil
}

var GateWayIp string


func EdgeLogin(){
	fmt.Printf("Input ip:\n")

	fmt.Scanf("%s",&GateWayIp)
	ipfs:=ipfs.IPFSClient.NewClient()
	peers, err := ipfs.BootstrapAdd([]string{fmt.Sprintf(constdef.LocalTestNode)})
	if err != nil {
		log.Printf("[Login-BootstrapAdd-err]: %v", err)
		return
	}

	url := fmt.Sprintf("http://%v:%v/login", GateWayIp, constdef.GatewayPort)
	msg, err := httppack.PostJson(url, &operate.MyInfo)
	if err != nil {
		log.Printf("[Login-PostJson-err]: %v, url:%v", err, url)
		return
	}
	_, err = util.ResponseParse(msg)
	if err != nil {
		fmt.Errorf("[Login-PostJson-err]: %v, url:%v", err, url)
		return
	}
	log.Printf("My ipfs peers: %+v", peers)
	return
}

func EdgeAddFile(){
	fmt.Printf("Input file path:\n")
	var path string
	fmt.Scanf("%s",&path)
	size,name,err:=util.FileDetail(path)
	if err != nil {
		log.Printf("[AddFile-FileDetail-err]: %v, path:%v", err,path )
		return
	}
	ipfsCtrl:=ipfs.NewIPFSCtrl()
	hash,err :=ipfsCtrl.UploadIPFSFile(path)
	if err != nil {
		log.Printf("[AddFile-UploadIPFSFile-err]: %v, path:%v", err,path )
		return
	}
	if GateWayIp == ""{
		log.Printf("[AddFile-err]:GateWayIp is nil" )
		return
	}
	url := fmt.Sprintf("http://%v:%v/ipfsadd", GateWayIp, constdef.GatewayPort)
	body:=service.IPFSAddParams{
		FileHash:     hash,
		FileSize:     size,
		Title:        name,
		Uploader:     operate.MyInfo.Dhash,
		BackupAmount: 3,
		BackupNumber: 0,
	}
	msg,err:= httppack.PostJson(url,body)
	_,msgErr:=util.ResponseParse(msg)
	if err != nil || msgErr!=nil{
		log.Printf("[AddFile-PostJson-err]:err=%v msgErr=%v ,url:%v", err,msgErr,url)
	}
	log.Printf("Add Succ, hash:%v name:%v\n size :%v",hash,name,size)
	return
}

func EdgeReadFile(){
	fmt.Printf("Read File List:\n\n")

	ans :=""
	limit := int64(3)
	offset := int64(0)

	total,ans,err:=readReq(limit,offset)
	if err != nil {
		log.Printf("[ReadFile-readReq-err]: %v ", err)
		return
	}
	for{
		total,ans,err =readReq(limit,offset)
		if err != nil {
			log.Printf("[ReadFile-readReq-err]: %v", err)
			return
		}
		err:=printfReadList(offset,ans)
		if err != nil {
			log.Printf("[ReadFile-PrintfReadList-err]: %v, ans:%v", err, ans)
			return
		}
		var order string
		fmt.Scanf("%s",&order)
		switch order{
		case "0":
			break
		case "P":
			if offset - limit >=0{offset-=limit}
		case "N":
			if offset + limit <total{offset+=limit}
		default:
			fmt.Println("unknow order, try again")
		}


	}

}

func readReq(limit,offset int64)(int64,string,error){
	url := fmt.Sprintf("http://%v:%v/getfilelist", GateWayIp, constdef.GatewayPort)
	params:=service.GetFileListParams{
		Limit:      limit,
		Offset:     offset,
	}
	msg,err:=httppack.PostJson(url,params)
	if err != nil {
		return 0,"",fmt.Errorf("[ReadFile-httpPost-err]: %v, url:%v", err, url)
	}
	msgJson, err := util.ResponseParse(msg)
	if err != nil {
		return 0,"",fmt.Errorf("[ReadFile-httpPost-err]: %v, url:%v", err, url)
	}
	total:= msgJson.Get("total").MustInt64()
	ansList:= msgJson.Get("ansList").MustString()
	return total,ansList ,nil
}


func printfReadList(offset int64,ans string)error{
	if ans == ""{
		return nil
	}
	var ansList []*model.FileInfo
	err:=json.Unmarshal([]byte(ans),&ansList)
	if err !=nil{
		return fmt.Errorf("[PrintfReadList-err]:%v",err)
	}
	for i,ans :=range ansList{
		//var show model.FileInfo
		//show = &ans
		fmt.Printf("%v: %+v\n",offset+int64(i+1),ans)
	}
	fmt.Println("<=============================>")
	fmt.Println("N: next P:previous")
	fmt.Println("0. Exit")
	fmt.Println("<=============================>")
	fmt.Println("input :")
	return nil

}

func EdgeGetFile(){
	fmt.Printf("input Get File fhash:\n")
	var fhash string
	fmt.Scanf("%s",&fhash)
	ipfsCtrl := ipfs.NewIPFSCtrl()
	err := ipfsCtrl.GetIPFSFile(fhash, constdef.OutputFilePath+fhash)
	if err != nil {
		log.Printf("[IpfsSave-GetIPFSFile-err]:%v", err)
	}
}