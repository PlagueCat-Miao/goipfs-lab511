package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

//输出当前用户home路径
var myHomePath string

func ShowMyHomePath() (string, error) {
	if myHomePath != "" {
		return myHomePath, nil
	}
	myUserInfo, err := user.Current()
	myHomePath = myUserInfo.HomeDir
	if err != nil {
		return "", fmt.Errorf("[ShowMyHomePath] failed!, err=%v", err)
	}
	return myHomePath, nil
}

// 递归创建文件夹
func MkdirP(paths ...string) error {
	for _, path := range paths {
		exist, _ := PathExists(path)
		if exist {
			continue
		}
		// 创建文件夹
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return fmt.Errorf("[MkdirP] mkdir failed!, err=%v path=%v", err, path)
		}
	}
	return nil
}

func Touch(filename string) error{
	// 创建文件
	fp, err := os.Create(filename)  // 如果文件已存在，会将文件清空。
	defer fp.Close()  //关闭文件，释放资源。
	if err != nil {
		//创建文件失败的原因有：
		//1、路径不存在  2、权限不足  3、打开文件数量超过上限  4、磁盘空间不足等
		return fmt.Errorf("os.Create fail :%v",err)
	}
	return nil
}


/*
*	计算整个目录所占磁盘大小
*	https://blog.csdn.net/HaoDaWang/article/details/80916385?utm_medium=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-2.channel_param&depth_1-utm_source=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-2.channel_param
*/
func DirSize(path string) int64 {
	//文件大小chennel
	fileSize := make(chan int64)
	//文件总大小
	var sizeCount int64

	//计算目录下所有文件占的大小总和
	go func() {
		walkDir(path, fileSize)
		defer close(fileSize)
	}()

	//t := time.Now()
	for size := range fileSize {
		sizeCount += size
	}
	//log.Println("搜索花费的时间为 " + time.Since(t).String())
	//log.Printf("该目录大小为 %vMB\n", sizeCount/1024 /1024)
	return sizeCount / 1024 / 1024

}

//递归计算目录下所有文件大小
func walkDir(path string, fileSize chan<- int64) {
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		log.Printf("[walkDir-debug]err: %v, path:%v", err, path)
	}
	for _, e := range entries {
		if e.IsDir() {
			walkDir(filepath.Join(path, e.Name()), fileSize)
		} else {
			fileSize <- e.Size()
		}
	}
}

//文件详细详细信息 (size, name, err)
func FileDetail(path string) (int64, string, error) {
	fi, err := os.Stat(path)
	if err != nil {
		log.Printf("[FileDetails-err]: %v, path:%v", err, path)
		return 0, "", err
	}
	return fi.Size() / 1024 / 1024, fi.Name(), nil
}


// 判断文件是否存在  存在返回 true 不存在返回false
func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
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