package constants

const (
	NacosAddr                      = "172.24.32.57"
	NacosPort                      = ":8848"
	ServerPort                     = ":8010"
	OpenApiPort                    = ":8011"
	UnusedPort                     = ":8012"
	QueryInstanceListsFromNacosUrl = "/nacos/v1/ns/instance/list"
	RegisterInstanceToNacosUrl     = "/nacos/v1/ns/instance"
	DeleteInstanceFromNacosUrl     = "/nacos/v1/ns/instance"
	UpdateInstanceToNacosUrl       = "/nacos/v1/ns/instance"

	HttpMethodGet    = "GET"
	HttpMethodPost   = "POST"
	HttpMethodDelete = "DELETE"
	HttpMethodPut    = "PUT"
)
