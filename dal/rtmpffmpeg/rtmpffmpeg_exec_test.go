package rtmpffmpeg

import (
	"fmt"
	"github.com/PlagueCat-Miao/goipfs-lab511/constdef"
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"os"
	"os/signal"
	"testing"
)

func TestRmptFfmpegCheck(t *testing.T) {
	Convey("[TestRmptFfmpegCheck] ffmpeg环境测试", t, func() {
		rmptCtrl:=NewRmptFfmpeg()
		err:=rmptCtrl.FfmpegExecCheck()
		So(err,ShouldBeNil)
	})
}

func TestRmptFfmpegPush(t *testing.T) {
	t.Logf("[TestRmptFfmpegPush] ffmpeg推流测试")
	rmptCtrl:=NewRmptFfmpeg()
	err:=rmptCtrl.FfmpegExecCheck()
	if err!=nil{
		log.Fatal(fmt.Errorf("[FfmpegExecCheck]: %v",err))
	}
	_,err =rmptCtrl.RtmpPushExec("127.0.0.1",constdef.RtmpDefaultPort,constdef.LiveRoom)
	if err!=nil{
		log.Fatal(fmt.Errorf("[RtmpPushExec]: %v",err))
	}

	quit := make(chan os.Signal) // ctrl+c
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Rtmp Push ...")

	err =rmptCtrl.Stop()
	if err!=nil{
		log.Fatal(fmt.Errorf("[RtmpPushExec]: %v",err))
	}
	log.Printf("[TestRmptFfmpegPush] ffmpeg推流测试ok")
}


func TestRmptFfmpegPull(t *testing.T) {
	t.Logf("[TestRmptFfmpegPull] ffmpeg推流测试")
	rmptCtrl:=NewRmptFfmpeg()
	err:=rmptCtrl.FfmpegExecCheck()
	if err!=nil{
		log.Fatal(fmt.Errorf("[FfmpegExecCheck]: %v",err))
	}
	_,err =rmptCtrl.RtmpPullExec("127.0.0.1",constdef.RtmpDefaultPort,constdef.LiveRoom)
	if err!=nil{
		log.Fatal(fmt.Errorf("[RtmpPushExec]: %v",err))
	}

	quit := make(chan os.Signal) // ctrl+c
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Rtmp Push ...")

	err =rmptCtrl.Stop()
	if err!=nil{
		log.Fatal(fmt.Errorf("[RtmpPushExec]: %v",err))
	}
	log.Printf("[TestRmptFfmpegPull] ffmpeg推流测试ok")
}


