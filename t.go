package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	//	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	baidu_url = "http://openapi.baidu.com/public/2.0/bmt/translate"
	api_key   = "CimNEmGPnYGpyyprUvhpGWAV"
)

var (
	w string
)

func translate(s string) (string, error) {
	url := baidu_url + "?from=auto&to=auto&client_id=" + api_key + "&q=" + w
	//fmt.Println(url)
	rsp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	res, err := ioutil.ReadAll(rsp.Body)
	rsp.Body.Close()
	if err != nil {
		return "", err
	}

	var m map[string]interface{}
	err = json.Unmarshal(res, &m)

	if err != nil {
		return "", err
	}

	if len(m["trans_result"].([]interface{})) <= 0 {
		return "", errors.New("no translate result")
	}

	return m["trans_result"].([]interface{})[0].(map[string]interface{})["dst"].(string), err
}

func main() {
	flag.StringVar(&w, "t", "", "input a word")
	flag.Parse()
	w = strings.Join(flag.Args(), " ")
	if w == "" {
		fmt.Println("enter a word")
		return
	}
	rs, err := translate(w)
	if err != nil {
		fmt.Printf("error occour: %s", err)
	}
	fmt.Println(rs)
}
