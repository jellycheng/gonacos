package gonacos

import (
	"errors"
	"fmt"
	"github.com/jellycheng/gcurl"
)

type NacosServiceManage struct {
	NacosServerConfig
}

// CreateService 创建服务
func (m *NacosServiceManage) CreateService(dto NacosServiceDto) (string, error) {
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

// EditService 修改服务
func (m *NacosServiceManage) EditService(dto NacosServiceDto) (string, error) {
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

// DeleteService 删除服务
func (m *NacosServiceManage) DeleteService(dto NacosServiceDto) (string, error) {
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

// GetService 查询一个服务
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

// GetServiceList 获取服务列表
func (m *NacosServiceManage) GetServiceList(dto NacosListServiceReqDto) (string, error) {
	urlStr := fmt.Sprintf("%s/nacos/v1/ns/service/list", m.NacosServerConfig.ApiUrl)
	params := map[string]interface{}{}
	if dto.PageNo != nil {
		params["pageNo"] = *dto.PageNo
	}
	if dto.PageSize != nil {
		params["pageSize"] = *dto.PageSize
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
