package gonacos

import (
	"fmt"
	"testing"
)

// go test -run=TestNacosServiceManage_RegisterInstance
func TestNacosServiceManage_RegisterInstance(t *testing.T) {
	nacosServerConfig := NacosServerConfig{ApiUrl: "http://127.0.0.1:8848", BeatInterval: 5 * 1000}
	nacosServiceManage := NewNacosServiceManage(nacosServerConfig)
	dto := NacosInstanceDto{
		ServiceName: StringPtr("mobile-api"),
		Ip:          StringPtr("127.0.0.1"),
		Port:        Int64Ptr(80),
		NamespaceId: StringPtr("dev"),
		GroupName:   StringPtr("mobile-api"),
	}
	if ok, err := nacosServiceManage.RegisterInstance(dto, true); ok {
		fmt.Println("服务注册成功")
		if nacosServiceManage.GetIsAutoBeat() {
			//select {}
		}
	} else {
		fmt.Println("服务注册失败：", err.Error())
	}
}

// go test -run=TestNacosServiceManage_SendBeat
func TestNacosServiceManage_SendBeat(t *testing.T) {
	nacosServerConfig := NacosServerConfig{ApiUrl: "http://127.0.0.1:8848", BeatInterval: 5 * 1000}
	nacosServiceManage := NewNacosServiceManage(nacosServerConfig)
	dto := BeatDto{
		ServiceName: StringPtr("mobile-api"),
		Ip:          StringPtr("127.0.0.1"),
		Port:        Int64Ptr(80),
		NamespaceId: StringPtr("dev"),
		GroupName:   StringPtr("mobile-api"),
	}
	if ok, err := nacosServiceManage.SendBeat(dto); ok {
		fmt.Println("发送成功")
	} else {
		fmt.Println("发送失败：", err.Error())
	}

}
