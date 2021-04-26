package main

import (
	"dsidecar/api"
	"dsidecar/constants"
	"dsidecar/http"
	"dsidecar/nacos"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"io"
	"log"
	"net"
	h "net/http"
	"strings"
	"time"
)

var (
	g errgroup.Group
)

// 从iptables的规则表中，根据 "源ip+port" 提取出 "目的ip+port"
func getDestIpPort(conn net.Conn) (string, string) {
	fakeDestIpPort := conn.LocalAddr().String()
	cmd := "cat /etc/sysconfig/iptables | grep \"" + fakeDestIpPort + "\" | awk -F ' ' '{print $4,$10}'"
	result := execute(cmd)
	data := strings.Split(result, " ")
	if len(data) == 1 {
		return "172.20.70.206", "8006"
	}
	// todo 判断data为空或者格式如果不是 "172.20.70.206/32 8004"这种的，就返回错误
	return strings.Split(data[0], "/")[0], data[1]
}

// 解析监听到的请求，提取出目的服务的url、请求参数，用于转发
func analyse(c *gin.Context) (string, string) {
	destIpPort := c.Request.Host
	uri := c.Request.RequestURI
	url := "http://" + destIpPort + uri
	buf := make([]byte, 1024)
	n, err := c.Request.Body.Read(buf)
	if err != io.EOF {
		log.Fatal(err)
	}

	body := string(buf[0:n])
	return url, body
}

// 解析请求并路由
func router() h.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/*params", func(c *gin.Context) {
		url, body := analyse(c)
		result := http.HttpGet(url, body)
		c.JSON(
			h.StatusOK,
			gin.H{
				"code":   h.StatusOK,
				"result": result,
			},
		)
	})
	e.POST("/*params", func(c *gin.Context) {
		url, body := analyse(c)
		result := http.HttpPost(url, body)
		c.JSON(
			h.StatusOK,
			gin.H{
				"code":   h.StatusOK,
				"result": result,
			},
		)
	})
	e.PUT("/*params", func(c *gin.Context) {
		url, body := analyse(c)
		result := http.HttpPut(url, body)
		c.JSON(
			h.StatusOK,
			gin.H{
				"code":   h.StatusOK,
				"result": result,
			},
		)
	})
	e.DELETE("/*params", func(c *gin.Context) {
		url, body := analyse(c)
		result := http.HttpDelete(url, body)
		c.JSON(
			h.StatusOK,
			gin.H{
				"code":   h.StatusOK,
				"result": result,
			},
		)
	})
	return e
}

// 启动一个不使用的端口，用于iptables nat转发
func startUnusedPort() {
	r := gin.Default()
	r.Run(constants.UnusedPort)
}

// 监听端口
func listener(addr string) {
	go startUnusedPort()

	server := &h.Server{
		Addr:         addr,
		Handler:      router(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	g.Go(func() error {
		return server.ListenAndServe()
	})
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}

// 初始化Nacos，注册当前服务到Nacos
func initNacos() {
	nacos.RegisterInstanceToNacos("localhost", 8080, "test-wy")
}

func main() {
	initNacos()
	api.InitApi()
	listener(constants.ServerPort)
}
