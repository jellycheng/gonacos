package gonacos

import (
	"errors"
	"fmt"
	"github.com/jellycheng/gcurl"
)

type NacosConfigManage struct {
	NacosServerConfig
}

// 获取配置
func (m *NacosConfigManage) GetConfig(dto NacosConfigDto) (string, error) {
	urlStr := fmt.Sprintf("%s/nacos/v1/cs/configs", m.NacosServerConfig.ApiUrl)
	params := map[string]interface{}{}
	if dto.Tenant != nil {
		params["tenant"] = *dto.Tenant
	}
	if dto.DataId != nil {
		params["dataId"] = *dto.DataId
	}
	if dto.Group != nil {
		params["group"] = *dto.Group
	}
	resp, err := gcurl.Get(urlStr, gcurl.Options{
		Headers: map[string]interface{}{
			"Content-Type": gcurl.ContentTypeForm,
			"User-Agent":   "gcurl/1.0",
		},
		Query: params,
	})
	if err != nil {
		return "", err
	} else {
		body, _ := resp.GetBody()
		if resp.GetStatusCode() != 200 {
			return "", errors.New(body.ToString())
		} else {
			return body.ToString(), nil
		}
	}

}

func NewNacosConfigManage(config NacosServerConfig) *NacosConfigManage {
	return &NacosConfigManage{
		NacosServerConfig: config,
	}
}
