package gonacos

import (
	"context"
	"errors"
	"fmt"
	"github.com/jellycheng/gcurl"
	"strings"
	"time"
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

// 监听配置变化
func (m *NacosConfigManage) ListenConfig(namespace, group, dataId string, fn func(cnf string)) {
	dto := NacosConfigDto{
		Tenant: StringPtr(namespace),
		Group:  StringPtr(group),
		DataId: StringPtr(dataId),
	}
	if cfgContent, err := m.GetConfig(dto); err == nil {
		fn(cfgContent)
		contentMd5 := MD5(cfgContent)
		beatTime := 2
		timeBeat := time.NewTimer(time.Second * time.Duration(beatTime))
		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			defer timeBeat.Stop()
			for {
				select {
				case <-timeBeat.C:
					if ischange, listenErr := m.GetListenConfig(namespace, group, dataId, contentMd5); listenErr != nil {
						cancel()
					} else if ischange { // 获取新配置
						if cfgContentNew, errNew := m.GetConfig(dto); errNew == nil {
							contentMd5 = MD5(cfgContentNew)
							fn(cfgContentNew)
						}
					}
					timeBeat.Reset(time.Second * time.Duration(beatTime))
				case <-ctx.Done():
					return
				}
			}
		}()
	} else {
		fmt.Println(err.Error())
	}

}

func (m *NacosConfigManage) GetListenConfig(namespace, group, dataId, contentMd5 string) (bool, error) {
	urlStr := fmt.Sprintf("%s/nacos/v1/cs/configs/listener", m.NacosServerConfig.ApiUrl)
	content := dataId + SplitConfigInner + group + SplitConfigInner + contentMd5 + SplitConfigInner + namespace + SplitConfig
	params := map[string]interface{}{}
	params["Listening-Configs"] = content
	resp, err := gcurl.Post(urlStr, gcurl.Options{
		Headers: map[string]interface{}{
			"Content-Type":         gcurl.ContentTypeForm,
			"User-Agent":           "gcurl/1.0",
			"Long-Pulling-Timeout": 3000,
		},
		FormParams: params,
	})
	if err != nil {
		return false, err
	} else {
		body, _ := resp.GetBody()
		if resp.GetStatusCode() != 200 {
			return false, errors.New(body.ToString())
		}

		str := strings.Split(body.ToString(), "%02")
		// 返回内容不为空且是传入的dataId 则说明内容有变化
		if len(str) > 0 && str[0] == dataId {
			return true, nil
		}
	}
	return false, nil
}

func NewNacosConfigManage(config NacosServerConfig) *NacosConfigManage {
	return &NacosConfigManage{
		NacosServerConfig: config,
	}
}
