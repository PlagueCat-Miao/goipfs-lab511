package util

import (
	"fmt"
	"go.uber.org/config"
	"os"
)

func LoadYaml(path string, target interface{},key string)error{
	fi,err := os.Open(path)
	if err != nil {
		return fmt.Errorf("[LoadYaml-Open-err]:%v",err)
	}
	provider, err := config.NewYAML(config.Source(fi))
	if  err != nil {
		return fmt.Errorf("[LoadYaml-NewYAML-err]:%v",err)
	}
	err= provider.Get(key).Populate(target)
	if err != nil {
		return fmt.Errorf("[LoadYaml-Get-err]:%v",err)
	}
	return nil
}
