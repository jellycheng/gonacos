package gonacos

import (
	"context"
	"errors"
	"fmt"
	"github.com/jellycheng/gcurl"
	"time"
)

type NacosInstanceManage struct {
	NacosServerConfig
	isAutoBeat bool
}

func (m NacosInstanceManage) GetIsAutoBeat() bool {
	return m.isAutoBeat
}

// RegisterInstance 注册服务
func (m *NacosInstanceManage) RegisterInstance(dto NacosInstanceDto, isAutoBeat bool) (string, error) {
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
		return "", err
	} else {
		body, _ := resp.GetBody()
		if resp.GetStatusCode() != 200 {
			return "", errors.New(body.ToString())
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
							if _, errBeat := m.SendBeat(beatDtoObj); errBeat != nil {
								cancel()
							} else {
								timeBeat.Reset(time.Millisecond * time.Duration(beatTime))
							}
						case <-ctx.Done():
							return
						}
					}
				}()
			}
			return body.ToString(), nil
		}
	}

}

// SendBeat 发送心跳
func (m *NacosInstanceManage) SendBeat(dto BeatDto) (string, error) {
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

// DeleteInstance 删除实例,注销实例
func (m *NacosInstanceManage) DeleteInstance(dto NacosInstanceDto) (string, error) {
	urlStr := fmt.Sprintf("%s/nacos/v1/ns/instance", m.NacosServerConfig.ApiUrl)
	params := map[string]interface{}{}
	if dto.ServiceName != nil {
		params["serviceName"] = *dto.ServiceName
	}
	if dto.GroupName != nil {
		params["groupName"] = *dto.GroupName
	}

	if dto.Ip != nil {
		params["ip"] = *dto.Ip
	}
	if dto.Port != nil {
		params["port"] = *dto.Port
	}
	if dto.ClusterName != nil {
		params["clusterName"] = *dto.ClusterName
	}
	if dto.NamespaceId != nil {
		params["namespaceId"] = *dto.NamespaceId
	}

	if dto.Ephemeral != nil {
		params["ephemeral"] = *dto.Ephemeral
	}

	resp, err := gcurl.Delete(urlStr, gcurl.Options{
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

// EditInstance 修改实例，主要是为了修改weight、metadata、enabled、ephemeral等字段，其它字段为查询条件
func (m *NacosInstanceManage) EditInstance(dto NacosInstanceDto) (string, error) {
	urlStr := fmt.Sprintf("%s/nacos/v1/ns/instance", m.NacosServerConfig.ApiUrl)
	params := map[string]interface{}{}
	if dto.ServiceName != nil {
		params["serviceName"] = *dto.ServiceName
	}
	if dto.GroupName != nil {
		params["groupName"] = *dto.GroupName
	}
	if dto.Ip != nil {
		params["ip"] = *dto.Ip
	}
	if dto.Port != nil {
		params["port"] = *dto.Port
	}
	if dto.ClusterName != nil {
		params["clusterName"] = *dto.ClusterName
	}
	if dto.NamespaceId != nil {
		params["namespaceId"] = *dto.NamespaceId
	}
	if dto.Weight != nil {
		params["weight"] = *dto.Weight
	}
	if dto.Metadata != nil {
		params["metadata"] = *dto.Metadata
	}
	if dto.Enable != nil {
		params["enabled"] = *dto.Enable
	}
	if dto.Ephemeral != nil {
		params["ephemeral"] = *dto.Ephemeral
	}

	resp, err := gcurl.Put(urlStr, gcurl.Options{
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

// GetInstanceList 查询服务下的实例列表,发现服务
func (m *NacosInstanceManage) GetInstanceList(dto NacosInstanceListReqDto) (string, error) {
	urlStr := fmt.Sprintf("%s/nacos/v1/ns/instance/list", m.NacosServerConfig.ApiUrl)
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
	if dto.Clusters != nil {
		params["clusters"] = *dto.Clusters
	}
	if dto.HealthyOnly != nil {
		params["healthyOnly"] = *dto.HealthyOnly
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

// GetInstanceView 查询实例详情：查询一个服务下个某个实例详情
func (m *NacosInstanceManage) GetInstanceView(dto NacosInstanceViewReqDto) (string, error) {
	urlStr := fmt.Sprintf("%s/nacos/v1/ns/instance", m.NacosServerConfig.ApiUrl)
	params := map[string]interface{}{}
	if dto.ServiceName != nil {
		params["serviceName"] = *dto.ServiceName
	}
	if dto.GroupName != nil {
		params["groupName"] = *dto.GroupName
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
	if dto.Clusters != nil {
		params["clusters"] = *dto.Clusters
	}
	if dto.HealthyOnly != nil {
		params["healthyOnly"] = *dto.HealthyOnly
	}
	if dto.Ephemeral != nil {
		params["ephemeral"] = *dto.Ephemeral
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

func NewNacosInstanceManage(config NacosServerConfig) NacosInstanceManage {
	return NacosInstanceManage{
		NacosServerConfig: config,
	}
}
