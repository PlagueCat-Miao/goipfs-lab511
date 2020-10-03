package main

import (
	"context"

	"github.com/PlagueCat-Miao/goipfs-lab511/nodes"
	"github.com/PlagueCat-Miao/goipfs-lab511/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func serverListen(router *gin.Engine,port int) {
	server := &http.Server{
		Addr:    ":"+strconv.Itoa(port),
		Handler: router,
	}
	log.Printf("ListenAndServePort: %+v",port)
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
	//<================================初始化==================================>
	port,err:=nodes.InitCloudServive()
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
	serverListen(router,port)

	//service.UManagement.SaveUserCSV()
	log.Println("server exiting...")

}
