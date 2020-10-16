package main

import (
	"context"
	"flag"
	"github.com/PlagueCat-Miao/goipfs-lab511/constdef"
	"github.com/PlagueCat-Miao/goipfs-lab511/dal/db"
	"github.com/PlagueCat-Miao/goipfs-lab511/operate"

	"github.com/PlagueCat-Miao/goipfs-lab511/nodes"
	"github.com/PlagueCat-Miao/goipfs-lab511/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func serverListen(router *gin.Engine, port int) {
	server := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: router,
	}
	log.Printf("ListenAndServePort: %+v", port)
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen:%s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
	//得到在当前上下文，并设定的了死亡时间
	//等待中断信号以优雅地关闭服务器（设置 10 秒的超时时间）
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//首先停止接受所有新请求，并一个个处理旧请求，
	// server.Shutdown()通过case <-ctx.Done(): 实现超时时及时退出
	// 于min( 上下文的死亡时间，全部请求处理结束) 时间 停止堵塞。
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown: ", err)
	}
}

func main() {
	//<================================入参解析=================================>
	var status int
	var help bool
	flag.IntVar(&status, "s", int(constdef.GatewayStatus), "身份")
	flag.BoolVar(&help, "h", false, "帮助")
	go func() {
		http.ListenAndServe("127.0.0.1:8848", nil)
	}()
	//解析命令行参数
	flag.Parse()
	if help {
		log.Println("USAGE \n use Makefile ")
		return
	}
	//<================================初始化==================================>
	var port int
	var err error
	switch status {
	case int(constdef.GatewayStatus):
		port, err = nodes.InitGatewayServive()
	case int(constdef.CloudStatus):
		port,err=nodes.InitCloudServive()
	default:
		log.Printf("[status-err]: invail status,status:%+v",status)
		return
	}

	if err != nil {
		log.Printf("[initServive-err]:%v", err)
		return
	}

	router := gin.Default()
	//<================================功能注册================================>
	router.POST("/login", service.Login)
	router.POST("/ipfsadd", service.IpfsAdd)
	router.POST("/ipfssave", service.IpfsSave)
	router.POST("/ipfsreport", service.IpfsReport)

	router.POST("/getfilelist", service.GetFileList)
	//<================================开启服务================================>
	serverListen(router, port)

	//<================================结束程序================================>
	log.Println("server exiting...")
	if db.FileInfoDB != nil {
		db.FileInfoDB.Close()
	}
	if operate.MyInfo.Status == constdef.GatewayStatus {
		operate.ClientsMgr.SaveUserCSV()
	}

}
