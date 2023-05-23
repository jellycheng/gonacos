package gonacos

import (
	"context"
	"errors"
	"fmt"
	"github.com/jellycheng/gcurl"
	"time"
)

type NacosServiceManage struct {
	NacosServerConfig
	isAutoBeat bool
}

func (m NacosServiceManage) GetIsAutoBeat() bool {
	return m.isAutoBeat
}

// 注册服务
func (m *NacosServiceManage) RegisterInstance(dto NacosInstanceDto, isAutoBeat bool) (bool, error) {
	urlStr := fmt.Sprintf("%s/nacos/v1/ns/instance", m.NacosServerConfig.ApiUrl)
	params := map[string]interface{}{}
	if dto.Ip != nil {
		params["ip"] = *dto.Ip
	}
	if dto.Port != nil {
		params["port"] = *dto.Port
	}
	if dto.NamespaceId != nil {
		params["namespaceId"] = *dto.NamespaceId
	}
	if dto.Weight != nil {
		params["weight"] = *dto.Weight
	}
	if dto.Enable != nil {
		params["enabled"] = *dto.Enable
	}
	if dto.Healthy != nil {
		params["healthy"] = *dto.Healthy
	}
	if dto.Metadata != nil {
		params["metadata"] = *dto.Metadata
	}
	if dto.ClusterName != nil {
		params["clusterName"] = *dto.ClusterName
	}
	if dto.ServiceName != nil {
		params["serviceName"] = *dto.ServiceName
	}
	if dto.GroupName != nil {
		params["groupName"] = *dto.GroupName
	}
	if dto.Ephemeral != nil {
		params["ephemeral"] = *dto.Ephemeral
	}

	resp, err := gcurl.Post(urlStr, gcurl.Options{
		Headers: map[string]interface{}{
			//"Content-Type": "application/x-www-form-urlencoded",
			"Content-Type": gcurl.ContentTypeForm,
			"User-Agent":   "gcurl/1.0",
		},
		Query: params,
	})
	if err != nil {
		return false, err
	} else {
		body, _ := resp.GetBody()
		if resp.GetStatusCode() != 200 {
			return false, errors.New(body.ToString())
		} else {
			// 开启心跳
			if isAutoBeat {
				m.isAutoBeat = isAutoBeat
				beatTime := m.NacosServerConfig.BeatInterval
				if beatTime == 0 {
					beatTime = 5 * 1000
				}
				timeBeat := time.NewTimer(time.Millisecond * time.Duration(beatTime))
				ctx, cancel := context.WithCancel(context.Background())
				beatDtoObj := BeatDto{
					ServiceName: dto.ServiceName,
					Ip:          dto.Ip,
					Port:        dto.Port,
				}
				if dto.NamespaceId != nil {
					beatDtoObj.NamespaceId = dto.NamespaceId
				}
				if dto.GroupName != nil {
					beatDtoObj.GroupName = dto.GroupName
				}
				if dto.Ephemeral != nil {
					beatDtoObj.Ephemeral = dto.Ephemeral
				}
				go func() {
					defer timeBeat.Stop()
					for {
						select {
						case <-timeBeat.C:
							if ok, _ := m.SendBeat(beatDtoObj); ok != true {
								cancel()
							}
							timeBeat.Reset(time.Millisecond * time.Duration(beatTime))
						case <-ctx.Done():
							return
						}
					}
				}()
			}
			return true, nil
		}
	}

}

// 发送心跳
func (m *NacosServiceManage) SendBeat(dto BeatDto) (bool, error) {
	urlStr := fmt.Sprintf("%s/nacos/v1/ns/instance/beat", m.NacosServerConfig.ApiUrl)
	params := map[string]interface{}{}
	if dto.ServiceName != nil {
		params["serviceName"] = *dto.ServiceName
	}
	if dto.Ip != nil {
		params["ip"] = *dto.Ip
	}
	if dto.Port != nil {
		params["port"] = *dto.Port
	}
	if dto.NamespaceId != nil {
		params["namespaceId"] = *dto.NamespaceId
	}
	if dto.GroupName != nil {
		params["groupName"] = *dto.GroupName
	}
	if dto.Ephemeral != nil {
		params["ephemeral"] = *dto.Ephemeral
	}
	if dto.Beat != nil {
		params["beat"] = *dto.Beat
	}

	resp, err := gcurl.Put(urlStr, gcurl.Options{
		Headers: map[string]interface{}{
			"Content-Type": gcurl.ContentTypeForm,
			"User-Agent":   "gcurl/1.0",
		},
		Query: params,
	})
	if err != nil {
		return false, err
	} else {
		body, _ := resp.GetBody()
		if resp.GetStatusCode() != 200 {
			return false, errors.New(body.ToString())
		} else {
			return true, nil
		}
	}
}

// 创建服务
func (m *NacosServiceManage) CreateService(dto NacosServiceDto) (bool, error) {
	urlStr := fmt.Sprintf("%s/nacos/v1/ns/service", m.NacosServerConfig.ApiUrl)
	params := map[string]interface{}{}
	if dto.ServiceName != nil {
		params["serviceName"] = *dto.ServiceName
	}
	if dto.GroupName != nil {
		params["groupName"] = *dto.GroupName
	}
	if dto.NamespaceId != nil {
		params["namespaceId"] = *dto.NamespaceId
	}
	if dto.ProtectThreshold != nil {
		params["protectThreshold"] = *dto.ProtectThreshold
	}
	if dto.Metadata != nil {
		params["metadata"] = *dto.Metadata
	}
	if dto.Selector != nil {
		params["selector"] = *dto.Selector
	}
	resp, err := gcurl.Post(urlStr, gcurl.Options{
		Headers: map[string]interface{}{
			"Content-Type": gcurl.ContentTypeForm,
			"User-Agent":   "gcurl/1.0",
		},
		Query: params,
	})
	if err != nil {
		return false, err
	} else {
		body, _ := resp.GetBody()
		if resp.GetStatusCode() != 200 {
			return false, errors.New(body.ToString())
		} else {
			return true, nil
		}
	}
}

// 修改服务
func (m *NacosServiceManage) EditService(dto NacosServiceDto) (bool, error) {
	urlStr := fmt.Sprintf("%s/nacos/v1/ns/service", m.NacosServerConfig.ApiUrl)
	params := map[string]interface{}{}
	if dto.ServiceName != nil {
		params["serviceName"] = *dto.ServiceName
	}
	if dto.GroupName != nil {
		params["groupName"] = *dto.GroupName
	}
	if dto.NamespaceId != nil {
		params["namespaceId"] = *dto.NamespaceId
	}
	if dto.ProtectThreshold != nil {
		params["protectThreshold"] = *dto.ProtectThreshold
	}
	if dto.Metadata != nil {
		params["metadata"] = *dto.Metadata
	}
	if dto.Selector != nil {
		params["selector"] = *dto.Selector
	}
	resp, err := gcurl.Put(urlStr, gcurl.Options{
		Headers: map[string]interface{}{
			"Content-Type": gcurl.ContentTypeForm,
			"User-Agent":   "gcurl/1.0",
		},
		Query: params,
	})
	if err != nil {
		return false, err
	} else {
		body, _ := resp.GetBody()
		if resp.GetStatusCode() != 200 {
			return false, errors.New(body.ToString())
		} else {
			return true, nil
		}
	}
}

// 删除服务
func (m *NacosServiceManage) DeleteService(dto NacosServiceDto) (bool, error) {
	urlStr := fmt.Sprintf("%s/nacos/v1/ns/service", m.NacosServerConfig.ApiUrl)
	params := map[string]interface{}{}
	if dto.ServiceName != nil {
		params["serviceName"] = *dto.ServiceName
	}
	if dto.GroupName != nil {
		params["groupName"] = *dto.GroupName
	}
	if dto.NamespaceId != nil {
		params["namespaceId"] = *dto.NamespaceId
	}

	resp, err := gcurl.Delete(urlStr, gcurl.Options{
		Headers: map[string]interface{}{
			"Content-Type": gcurl.ContentTypeForm,
			"User-Agent":   "gcurl/1.0",
		},
		Query: params,
	})
	if err != nil {
		return false, err
	} else {
		body, _ := resp.GetBody()
		if resp.GetStatusCode() != 200 {
			return false, errors.New(body.ToString())
		} else {
			return true, nil
		}
	}
}

// 查询一个服务
func (m *NacosServiceManage) GetService(dto NacosServiceDto) (string, error) {
	urlStr := fmt.Sprintf("%s/nacos/v1/ns/service", m.NacosServerConfig.ApiUrl)
	params := map[string]interface{}{}
	if dto.ServiceName != nil {
		params["serviceName"] = *dto.ServiceName
	}
	if dto.GroupName != nil {
		params["groupName"] = *dto.GroupName
	}
	if dto.NamespaceId != nil {
		params["namespaceId"] = *dto.NamespaceId
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

func NewNacosServiceManage(config NacosServerConfig) NacosServiceManage {
	return NacosServiceManage{
		NacosServerConfig: config,
	}
}
