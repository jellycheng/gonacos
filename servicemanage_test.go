package gonacos

import (
	"fmt"
	"testing"
)

// go test -run=TestNacosServiceManage_CreateService
func TestNacosServiceManage_CreateService(t *testing.T) {
	nacosServerConfig := NacosServerConfig{ApiUrl: "http://console.nacos.io"}
	nacosServiceManage := NewNacosServiceManage(nacosServerConfig)
	dto := NacosServiceDto{
		ServiceName: StringPtr("mobile-api-01"),
		NamespaceId: StringPtr("dde761c3-96be-4a98-b349-3c2289033322"),
		GroupName:   StringPtr("mobile-api"),
		//ProtectThreshold: Float64Ptr(1),
	}
	if content, err := nacosServiceManage.CreateService(dto); err == nil {
		fmt.Println("服务创建成功", content)
	} else {
		fmt.Println("服务创建失败：", err.Error())
	}
}

// go test -run=TestNacosServiceManage_EditService
func TestNacosServiceManage_EditService(t *testing.T) {
	nacosServerConfig := NacosServerConfig{ApiUrl: "http://console.nacos.io"}
	nacosServiceManage := NewNacosServiceManage(nacosServerConfig)
	dto := NacosServiceDto{
		ServiceName:      StringPtr("mobile-api-01"),
		NamespaceId:      StringPtr("dde761c3-96be-4a98-b349-3c2289033322"),
		GroupName:        StringPtr("mobile-api"),
		ProtectThreshold: Float64Ptr(1),
		Metadata:         StringPtr("helloworld=abc"),
	}
	if content, err := nacosServiceManage.EditService(dto); err == nil {
		fmt.Println("修改服务成功", content)
	} else {
		fmt.Println("修改服务失败：", err.Error())
	}
}
