package util

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

/*
*   计算整个目录所占磁盘大小
 */
// https://blog.csdn.net/HaoDaWang/article/details/80916385?utm_medium=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-2.channel_param&depth_1-utm_source=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-2.channel_param

func DirSize(path string) int64{
	//文件大小chennel
	fileSize := make(chan int64)
	//文件总大小
	var sizeCount int64

	//计算目录下所有文件占的大小总和
	go func(){
		walkDir(path, fileSize)
		defer close(fileSize)
	}()

	//t := time.Now()
	for size := range fileSize {
		sizeCount += size
	}
	//log.Println("搜索花费的时间为 " + time.Since(t).String())
	//log.Printf("该目录大小为 %vMB\n", sizeCount/1024 /1024)
	return sizeCount/1024 /1024

}

//递归计算目录下所有文件
func walkDir(path string, fileSize chan <- int64){
	entries, err := ioutil.ReadDir(path)
	if err !=nil{
		log.Printf("[walkDir-debug]err: %v, path:%v",err,path)
	}
	for _, e := range entries{
		if e.IsDir() {
			walkDir(filepath.Join(path, e.Name()), fileSize)
		} else {
			fileSize <- e.Size()
		}
	}
}
func FileDetail(path string) (int64,string,error){
	fi,err:=os.Stat(path)
	if err !=nil {
		log.Printf("[FileDetails-err]: %v, path:%v",err,path)
		return 0,"",err
	}
	return  fi.Size()/1024/1024,fi.Name(),nil
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}