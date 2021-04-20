package ipfs

import (
	"encoding/json"
	"fmt"
	"github.com/PlagueCat-Miao/goipfs-lab511/constdef"
	"github.com/PlagueCat-Miao/goipfs-lab511/util"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

var GOPath string

const (
	testFilePath   = "/src/github.com/PlagueCat-Miao/GOIPFS-gateway/test/test.json"
	outputFilePath = "/src/github.com/PlagueCat-Miao/GOIPFS-gateway/test/output.json"
)

func init() {
	GOPath = os.Getenv("GOPATH")
	fmt.Println("[IpfsAPI-单测] 初始化完成") //将执行于 TestMain()之前
}

//交易结构体(未来的通道)
type MockTransaction struct {
	Person1      string `json:"person1,omitempty" xml:"person1"`
	Person2      string `json:"person2,omitempty" xml:"person2"`
	Person1money string `json:"person1Money,omitempty" xml:"person1Money"`
	Person2money string `json:"person2Money,omitempty" xml:"person2Money"`
}

//通道序列化
func marshalStruct(transaction MockTransaction) []byte {

	data, err := json.Marshal(&transaction)
	if err != nil {
		fmt.Println("序列化err=", err)
	}
	return data
}

//数据反序列化为通道
func unmarshalStruct(str []byte) MockTransaction {
	var transaction MockTransaction
	err := json.Unmarshal(str, &transaction)
	if err != nil {
		fmt.Printf("unmarshal err=%v", err)
	}
	return transaction
}

func TestIpfsAPI(t *testing.T) {
	ipfs := NewIPFSCtrl()
	//生成一个交易结构体(未来的通道)
	transaction := MockTransaction{
		Person1:      "Aaron",
		Person2:      "Bob",
		Person1money: "100",
		Person2money: "200",
	}
	//结构体序列化
	data := marshalStruct(transaction)
	//上传到ipfs
	hash, err := ipfs.UploadIPFS(string(data))
	if err != nil {
		util.ErrLog(err, constdef.Show, constdef.ParamNil)

		t.Errorf("errLog为json文件 ，可以使用json 格式化工具 jq")
		return
	}
	fmt.Println("文件hash是", hash)
	//从ipfs下载数据
	str2, err := ipfs.CatIPFS(hash)
	if err != nil {
		//mmlog.ErrLog(err, constdef.Show, constdef.ParamNil)
		t.Errorf("errLog为json文件 ，可以使用json 格式化工具 jq")
		return
	}
	//数据反序列化
	transaction2 := unmarshalStruct([]byte(str2))
	//验证下数据
	fmt.Println(transaction2)
	t.Log("IpfsAPI  测试结束")
}

func TestIpfsFileSys(t *testing.T) {
	//单测全局变量，上下convey传递的部分
	var ReadHash = "???"

	Convey("[TestIpfsFileSys] ipfs文件上传/下载测试", t, func() {
		//这个部分属于子Convey的公共初始化部分
		//子Convey按程序顺序串行运行，但每个子Convey都将重新执行"公共初始化部分"
		ipfs := NewIPFSCtrl()
		fmt.Println("\nipfs succ")

		Convey("[上传] success", func() {
			hash, err := ipfs.UploadIPFSFile(GOPath + testFilePath)
			So(err, ShouldBeNil)
			So(hash, ShouldNotBeNil)
			ReadHash = hash
		})
		Convey("[下载] success", func() {
			So(ReadHash, ShouldNotEqual, "???")
			err := ipfs.GetIPFSFile(ReadHash, GOPath+outputFilePath)
			So(err, ShouldBeNil)
		})
	})
}
