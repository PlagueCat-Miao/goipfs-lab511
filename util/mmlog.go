package util

import (
	"fmt"
	"github.com/PlagueCat-Miao/goipfs-lab511/constdef"
	"log"
	"path/filepath"
	"runtime"
	"strings"
)

func ErrLog(err error, errType constdef.ErrLogType, paramFormat string, paramArgs ...interface{}) error {
	if err == nil {
		return nil
	}
	if errType == constdef.Show {
		log.Printf("错误：%v\n", err)
		return nil
	}

	if errType == constdef.Leaf {
		err = fmt.Errorf("{\"err\":\"%v\"}", err)
	}
	filename, line, funcname := "???", 0, "???"
	pc, filename, line, ok := runtime.Caller(1)
	if ok {
		funcname = runtime.FuncForPC(pc).Name()      // main.(*MyStruct).foo
		funcname = filepath.Ext(funcname)            // .foo
		funcname = strings.TrimPrefix(funcname, ".") // foo

		filename = filepath.Base(filename) // /full/path/basename.go => basename.go
	}
	return fmt.Errorf("{\"funcInfo\":\" %s:%d:%s\",\"Param\":\"%v\",\"err\":%v}", filename, line, funcname, fmt.Sprintf(paramFormat, paramArgs...), err)
}
