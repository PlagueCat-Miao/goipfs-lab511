package rtmpffmpeg

import (
	"context"
	"fmt"
	"github.com/PlagueCat-Miao/goipfs-lab511/constdef"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"

	//"time"
)

var RtmpCtx *context.Context
var RtmpCancelFunc *context.CancelFunc
var LogFileStdOut  *os.File

type IRtmpFfmpeg interface {
	FfmpegExecCheck() error
	RtmpPushExec(ip string,port int,room string) (string, error)
	RtmpPullExec(ip string,port int,room string) (string, error)
	Stop() error
}

type RtmpFfmpeg struct{}

func NewRmptFfmpeg() IRtmpFfmpeg {
	return &RtmpFfmpeg{}
}

func (r *RtmpFfmpeg) RtmpPullExec(ip string,port int,room string)(string, error){
	if RtmpCtx !=nil || RtmpCancelFunc!=nil{
		return "",fmt.Errorf("[func RtmpPullExec]  err= other exec is running")
	}
	//正在处理事件注册
	cancelCtx, cancelFun := context.WithCancel(context.Background()) //使用ctx保证控制
	RtmpCtx = &cancelCtx
	RtmpCancelFunc = &cancelFun
	url :=fmt.Sprintf(constdef.RtmpFormat,ip,port,room)
	cmdArgs := strings.Split(fmt.Sprintf(constdef.FfplayArgFormat,url)," ")
	cmd := exec.CommandContext(cancelCtx,constdef.FfplayCmd,cmdArgs...)
	// 执行命令
	if err := cmd.Start(); err != nil {
		return "",fmt.Errorf("[func RtmpPullExec Start] 新进程开始运行失败 err= %+v ", err)
	}

	return url,nil
}
func (r *RtmpFfmpeg)RtmpPushExec(ip string,port int,room string) (string, error){
	if RtmpCtx !=nil || RtmpCancelFunc!=nil{
		return "",fmt.Errorf("[func RtmpPushExec]  err= other exec is running")
	}

	cancelCtx, cancelFun := context.WithCancel(context.Background()) //使用ctx保证控制
	RtmpCtx = &cancelCtx
	RtmpCancelFunc = &cancelFun
	url:= fmt.Sprintf(constdef.RtmpFormat,ip,port,room)
	cmdArgs := strings.Split(fmt.Sprintf(constdef.FfmpegArgFormat,url)," ")

	cmd := exec.CommandContext(cancelCtx,constdef.FfmpegCmd,cmdArgs...)
	filePath:= fmt.Sprintf(constdef.PushLogPathFormat, ip,time.Now().Format("2006010215"))
	stdout, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return "", fmt.Errorf("[func RtmpPushExec OpenFile]  err= %+v ", err)
	}

	cmd.Stderr = stdout // 重定向标准输出到文件 注意状态输出在err下
	LogFileStdOut = stdout
	// 执行命令
	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("[func RtmpPushExec Start] 新进程开始运行失败 err= %+v ", err)
	}

	return url,nil

}

func (r *RtmpFfmpeg) Stop()error{
	if RtmpCtx == nil || RtmpCancelFunc == nil{
		return fmt.Errorf("[func Stop] not  err= no exec is running ")
	}
	(*RtmpCancelFunc)() // 运行取消
	RtmpCtx = nil
	RtmpCancelFunc = nil
	if LogFileStdOut != nil{
		LogFileStdOut.Close()
		LogFileStdOut = nil
	}

	return nil
}

func (r *RtmpFfmpeg)FfmpegExecCheck() error{
	var stdout io.ReadCloser
	var err error
	cmd := exec.Command(constdef.FfmpegCmd,"-version")
	if stdout, err = cmd.StdoutPipe(); err != nil {     //获取输出对象，可以从该对象中读取输出结果
		return fmt.Errorf("[func FfmpegExecCheck] 新进程stdout重定向为当前终端失败 err= %+v ", err)
	}
	defer stdout.Close()   // 保证关闭输出流

	if err := cmd.Start(); err != nil {   // 运行命令
		return fmt.Errorf("[func FfmpegExecCheck] 新进程开始运行失败 err= %+v ", err)
	}

	if opBytes, err := ioutil.ReadAll(stdout); err != nil {  // 读取输出结果
		return fmt.Errorf("[func FfmpegExecCheck]  新进程stdout读取失败 err= %+v ", err)
	} else {
		fmt.Printf("检测到ffmpeg :\n%v\n" ,string(opBytes))
	}
	err = cmd.Wait()
	if err != nil{
		return fmt.Errorf("[func FfmpegExecCheck]  新进程没有正常结束 err= %+v ", err)
	}
	return nil
}


