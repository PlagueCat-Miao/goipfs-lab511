package rtmpffmpeg

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
	"time"
)



type IRmptFfmpeg interface {
	FfmpegExecCheck() error
	//RtmpPushExec() error
	//RtmpPullExec(ip string) error
	//Stop() error
}

type RmptFfmpeg struct{}

func NewRmptFfmpeg() IRmptFfmpeg {
	return &RmptFfmpeg{}
}
func rtmpPushExec() error{
	var stdout io.ReadCloser
	var err error

	cancelCtx, _ := context.WithDeadline(context.Background(),time.Now().Add(2*time.Second)) //使用ctx保证控制
	cmd := exec.CommandContext(cancelCtx,"ffmpeg", "-version")
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
		log.Printf("检测到ffmpeg :\n%v\n",string(opBytes))
	}

	return nil

}


func (r *RmptFfmpeg)FfmpegExecCheck() error{
	var stdout io.ReadCloser
	var err error

	cmd := exec.Command("ffmpeg", "-version")
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
		log.Printf("[func FfmpegExecCheck]  新进程没有正常结束 err= %+v ", err)
	}
	return nil
}


