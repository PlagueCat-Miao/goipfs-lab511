package httppack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PlagueCat-Miao/goipfs-lab511/constdef"
	"io/ioutil"
	"net/http"
	"time"
)

func PostJson(url string, data interface{}) ([]byte, error) {
	// 超时时间：5秒
	client := &http.Client{Timeout: 3 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, constdef.JsonContentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, fmt.Errorf("[Post-err]:%v", err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return result, nil
}
