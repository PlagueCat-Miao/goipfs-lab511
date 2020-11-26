package util

import (
	"fmt"
	"os"
	"os/user"
)
//输出当前用户home路径
var myHomePath string
func ShowMyHomePath() (string,error){
	if myHomePath != ""{
		return myHomePath,nil
	}
	myUserInfo , err := user.Current()
	myHomePath = myUserInfo.HomeDir
	if err != nil{
		return "",fmt.Errorf("[ShowMyHomePath] failed!, err=%v", err)
	}
	return myHomePath,nil
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func MkdirP(paths ...string) error{
	for _,path :=range paths{
		exist,_:=PathExists(path)
		if exist {
			continue
		}
		// 创建文件夹
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return fmt.Errorf("[MkdirP] mkdir failed!, err=%v path=%v", err,path)
		}
	}
	return nil
}