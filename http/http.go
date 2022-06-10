package http

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	string2 "github.com/jeffcail/ctools/string"
)

var client *http.Client

func init() {
	def := http.DefaultTransport
	defPot, ok := def.(*http.Transport)
	if !ok {
		panic("init transport出错")
	}
	defPot.MaxIdleConns = 100
	defPot.MaxIdleConnsPerHost = 100
	defPot.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	//defPot.IdleConnTimeout = 60
	client = &http.Client{
		Timeout:   time.Second * time.Duration(20),
		Transport: defPot,
	}
}

func Get(url string, header map[string]string, params map[string]interface{}) ([]byte, error) {
	//client := &http.Client{
	//	Timeout: time.Second * time.Duration(20),
	//}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if header != nil {
		for k, v := range header {
			req.Header.Add(k, v)
		}
	}

	q := req.URL.Query()
	if params != nil {
		for Key, val := range params {
			v, _ := string2.ToString(val)
			q.Add(Key, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	r, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	if r.StatusCode != 200 {
		return nil, err
	}
	bb, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return bb, nil
}

func Post(url string, header map[string]string, param map[string]interface{}) ([]byte, error) {
	//client := &http.Client{
	//	Timeout: time.Second * time.Duration(20),
	//}
	dd, _ := json.Marshal(param)
	re := bytes.NewReader(dd)
	req, err := http.NewRequest("POST", url, re)
	if err != nil {
		return nil, err
	}
	if header != nil {
		for k, v := range header {
			req.Header.Add(k, v)
		}
	}
	r, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	bb, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return bb, nil
}
