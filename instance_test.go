package gonacos

import (
	"fmt"
	"testing"
)

// go test -run=TestNacosInstanceManage_RegisterInstance
func TestNacosInstanceManage_RegisterInstance(t *testing.T) {
	nacosServerConfig := NacosServerConfig{ApiUrl: "http://console.nacos.io", BeatInterval: 5 * 1000}
	nacosInstanceManage := NewNacosInstanceManage(nacosServerConfig)
	dto := NacosInstanceDto{
		ServiceName: StringPtr("mobile-api"),
		Ip:          StringPtr("127.0.0.1"),
		Port:        Int64Ptr(80),
		NamespaceId: StringPtr("st1"),
		GroupName:   StringPtr("mobile-api"),
	}
	if content, err := nacosInstanceManage.RegisterInstance(dto, true); err == nil {
		fmt.Println("服务注册成功", content)
		if nacosInstanceManage.GetIsAutoBeat() {
			//select {}
		}
	} else {
		fmt.Println("服务注册失败：", err.Error())
	}
}

// go test -run=TestNacosInstanceManage_SendBeat
func TestNacosInstanceManage_SendBeat(t *testing.T) {
	nacosServerConfig := NacosServerConfig{ApiUrl: "http://console.nacos.io", BeatInterval: 5 * 1000}
	nacosInstanceManage := NewNacosInstanceManage(nacosServerConfig)
	dto := BeatDto{
		ServiceName: StringPtr("mobile-api"),
		Ip:          StringPtr("127.0.0.1"),
		Port:        Int64Ptr(80),
		NamespaceId: StringPtr("st1"),
		GroupName:   StringPtr("mobile-api"),
	}
	if content, err := nacosInstanceManage.SendBeat(dto); err == nil {
		fmt.Println("发送成功", content)
	} else {
		fmt.Println("发送失败：", err.Error())
	}

}
