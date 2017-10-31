package httplib

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"jupiter/libs"
	"net/http"

	"github.com/dghubble/sling"
)

var (
	nxhttp *HttpService
)

type HttpService struct {
	Sling  *sling.Sling
	Client *http.Client // test
}

func Service(domain string) *HttpService {
	httpClient := GetHttp()
	return &HttpService{
		Sling:  sling.New().Client(httpClient).Base(domain),
		Client: GetHttp(),
	}
}

func (s *HttpService) Get(url string, params interface{}, v interface{}) (string, bool) {
	fmt.Println("called: Get")
	req, err := s.Sling.New().Get(url).QueryStruct(params).Request()
	fmt.Printf("req:%s, err:%s\n", req, err)
	resp, _ := GetHttp().Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("response: %s\n", string(body))
	return string(body), libs.JsonUnmarshal(body, v)
}

// test
func (s *HttpService) PostEx(url string, h http.Header, post interface{}, v interface{}) (string, bool) {
	data := libs.JsonMarshal(post)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	req.Header.Set("Content-Type", "application/json")

	for k, vv := range h {
		req.Header.Set(k, vv[0])
	}

	if resp, err := s.Client.Do(req); err == nil {
		defer resp.Body.Close()
		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("response Body:", string(body))
		libs.JsonUnmarshal(body, v)
		return string(body), true
	} else {
		fmt.Printf("PostEx: error - %s\n", err)
		return "", false
	}
}
