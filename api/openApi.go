package api

import (
	"dsidecar/constants"
	"dsidecar/nacos"
	"github.com/gin-gonic/gin"
)

func queryInstancesFromNacos(c *gin.Context) {
	serviceName := c.Query("serviceName")
	c.String(200, nacos.QueryInstanceListsFromNacos(serviceName))
}

func InitApi() {
	r := gin.Default()

	// 注册接口
	r.GET("/instances", queryInstancesFromNacos)

	// 端口要和注册到zipkin中的一致，即与serviceAddr的值一致
	go r.Run(constants.OpenApiPort)
}
