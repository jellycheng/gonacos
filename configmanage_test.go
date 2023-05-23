package gonacos

import (
	"fmt"
	"testing"
)

// go test -run=TestNewNacosConfigManage
func TestNewNacosConfigManage(t *testing.T) {
	nacosServerConfig := NacosServerConfig{ApiUrl: "http://console.nacos.io", BeatInterval: 5 * 1000}
	nacosConfigManage := NewNacosConfigManage(nacosServerConfig)

	dto := NacosConfigDto{
		Tenant: StringPtr("public"),
		DataId: StringPtr("application-dev.yaml"),
		Group:  StringPtr("DEFAULT_GROUP"),
	}
	if res, err := nacosConfigManage.GetConfig(dto); err == nil {
		fmt.Println(res)
	} else {
		fmt.Println("获取配置失败：", err.Error())
	}
}
