package nacos

import (
	"dsidecar/constants"
	"dsidecar/http"
	"fmt"
	"strconv"
)

// 查询实例
func QueryInstanceListsFromNacos(serviceName string) string {
	url := "http://" + constants.NacosAddr + constants.NacosPort + constants.QueryInstanceListsFromNacosUrl
	url += "?serviceName=" + serviceName
	result := http.HttpGet(url, nil)
	fmt.Println(result)
	return result
}

// 注册实例
func RegisterInstanceToNacos(ip string, port int, serviceName string) string {
	url := "http://" + constants.NacosAddr + constants.NacosPort + constants.RegisterInstanceToNacosUrl
	url += "?ip=" + ip + "&port=" + strconv.Itoa(port) + "&serviceName=" + serviceName
	result := http.HttpPost(url, nil)
	fmt.Println(result)
	return result
}

// 注销实例
func DeleteInstanceFromNacos(ip string, port int, serviceName string) string {
	url := "http://" + constants.NacosAddr + constants.NacosPort + constants.DeleteInstanceFromNacosUrl
	url += "?ip=" + ip + "&port=" + strconv.Itoa(port) + "&serviceName=" + serviceName
	result := http.HttpDelete(url, nil)
	fmt.Println(result)
	return result
}

// 修改实例
func UpdateInstanceToNacos(ip string, port int, serviceName string) string {
	url := "http://" + constants.NacosAddr + constants.NacosPort + constants.UpdateInstanceToNacosUrl
	url += "?ip=" + ip + "&port=" + strconv.Itoa(port) + "&serviceName=" + serviceName
	result := http.HttpPut(url, nil)
	fmt.Println(result)
	return result
}
