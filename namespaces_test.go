package gonacos

import (
	"fmt"
	"testing"
)

// go test -run=TestNacosNamespaceManage_CreateNamespace
func TestNacosNamespaceManage_CreateNamespace(t *testing.T) {
	nacosServerConfig := NacosServerConfig{ApiUrl: "http://console.nacos.io", BeatInterval: 5 * 1000}
	nacosNamespaceManage := NewNacosNamespaceManage(nacosServerConfig)
	dto := NacosNamespaceCreateDto{
		CustomNamespaceId: StringPtr("st1"),
		NamespaceName:     StringPtr("st1"),
		NamespaceDesc:     StringPtr("st1环境"),
	}
	if content, err := nacosNamespaceManage.CreateNamespace(dto); err == nil {
		fmt.Println("命名空间创建成功", content)
	} else {
		fmt.Println("命名空间创建失败：", err.Error())
	}
}
