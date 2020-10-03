package main

import (
	"context"
	"fmt"
	"github.com/PlagueCat-Miao/goipfs-lab511/constdef"
	"github.com/PlagueCat-Miao/goipfs-lab511/dal/ipfs"
	"github.com/PlagueCat-Miao/goipfs-lab511/operate"
	"github.com/PlagueCat-Miao/goipfs-lab511/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func serverListen(router *gin.Engine) {
	server := &http.Server{
		Addr:    ":8888",
		Handler: router,
	}
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
func initGatewayServive() error{
	//ipfs shell连接
	ipfsClient,err:=ipfs.InitIPFS()
	if err!=nil ||ipfsClient == nil {
		return fmt.Errorf("[ipfs-err]: %v",err)
	}
	//本机状态更新
	operate.InitMyInfo(ipfsClient.DHash,8888,constdef.GatewayStatus,100,100)
	return nil
}
func initCloudServive() error{
	//ipfs shell连接
	ipfsClient,err:=ipfs.InitIPFS()
	if err!=nil ||ipfsClient == nil {
		return fmt.Errorf("[ipfs-init-err]: %v",err)
	}
	ipfs := ipfsClient.NewClient()
	peers ,err:=ipfs.BootstrapAdd([]string{constdef.MyCloudNode})
	if err!=nil{
		return fmt.Errorf("[ipfs-BootstrapAdd-err]: %v",err)
	}
	log.Printf("My ipfs peers: %+v",peers)
	//本机状态更新
	operate.InitMyInfo(ipfsClient.DHash,8888,constdef.CloudStatus,100,100)
	return nil
}

func main() {
	//<================================初始化==================================>
	err:=initCloudServive()


	if err!=nil{
		log.Printf("[initServive-err]:%v",err)
		return
	}
	router := gin.Default()
	//<================================功能注册================================>
	router.POST("/login",service.Login)
	router.POST("/ipfsadd",service.IpfsAdd)
	router.POST("/ipfssave",service.IpfsSave)

	//<================================开启服务================================>
	serverListen(router)

	//service.UManagement.SaveUserCSV()
	log.Println("server exiting...")

}
