package http

import (
	"bytes"
	"dsidecar/constants"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func HttpRequest(method string, url string, requestBody interface{}) string {
	jsonStr, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	req.Header.Add("application", "json")
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

// 发送GET请求
// getUrl：         请求地址
func HttpGet(getUrl string, requestBody interface{}) string {
	return HttpRequest(constants.HttpMethodGet, getUrl, nil)
}

// 发送POST请求
// postUrl：		请求地址
// requestBody：	POST请求提交的数据
func HttpPost(postUrl string, requestBody interface{}) string {
	return HttpRequest(constants.HttpMethodPost, postUrl, requestBody)

}

// 发送DELETE请求
// deleteUrl：		请求地址
// requestBody：	DELETE请求提交的数据
func HttpDelete(deleteUrl string, requestBody interface{}) string {
	return HttpRequest(constants.HttpMethodDelete, deleteUrl, requestBody)
}

// 发送PUT请求
// putUrl：			请求地址
// requestBody：	PUT请求提交的数据
func HttpPut(putUrl string, requestBody interface{}) string {
	return HttpRequest(constants.HttpMethodPut, putUrl, requestBody)
}
