package utils

import (
	"sort"
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"io/ioutil"
	"fmt"
	"os"
	"encoding/json"
	"github.com/kataras/go-errors"
)

const (
	HOST = "http://www.kuaifazs.com"
	ORDER_DETAIL_PATH = "/forgoapi/order/export"
	SECRET_KEY = "5f5a740a75ed31c67c5d16eded29d30d"
)

type OrderDetailData struct {
	Url string
}

type OrderDetail struct {
	Status        string
	Error         string
	Error_message string
	Data          OrderDetailData
}

func GetDetailUrl(path string, params map[string]string) (string, error){
	data := GetData(path, params)
	//fmt.Println("data:", data)
	var order_detail OrderDetail
	json.Unmarshal([]byte(data), &order_detail)
	fmt.Println("order detail url detail:", order_detail.Data.Url)
	if order_detail.Status == "0" {
		return HOST + order_detail.Data.Url, nil
	} else {
		return "", errors.New(order_detail.Error_message)
	}
}

func parseURL(path string, parms map[string]string) string {
	keys := make([]string, len(parms))
	i := 0
	for k, _ := range parms {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	ret := ""
	for _, key := range keys {
		ret = ret + key + "=" + parms[key] + "&"
	}
	sign := getSign(ret[:len(ret) - 1])
	return HOST + path + "?" + ret + "sign=" + sign
}

func getSign(parms string) string {
	m := md5.New()
	m.Write([]byte(parms))
	tmp_md5 := hex.EncodeToString(m.Sum(nil))
	mm := md5.New()
	mm.Write([]byte(tmp_md5 + SECRET_KEY))
	return hex.EncodeToString(mm.Sum(nil))
}

func GetData(path string, params map[string]string) string {
	addr := parseURL(path, params)
	fmt.Println("addr:", addr)
	rsp, err := http.Get(addr)
	checkError(err)
	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)
	checkError(err)
	//fmt.Println("data type:", reflect.TypeOf(body))
	return string(body)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}