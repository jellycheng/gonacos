package gonacos

import (
	"errors"
	"fmt"
	"github.com/jellycheng/gcurl"
)

type NacosNamespaceManage struct {
	NacosServerConfig
}

// 查询命名空间列表
func (m *NacosNamespaceManage) GetNamespaces() (string, error) {
	urlStr := fmt.Sprintf("%s/nacos/v1/console/namespaces", m.NacosServerConfig.ApiUrl)

	resp, err := gcurl.Get(urlStr, gcurl.Options{
		Headers: map[string]interface{}{
			"Content-Type": gcurl.ContentTypeForm,
			"User-Agent":   "gcurl/1.0",
		},
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

// 创建命名空间
func (m *NacosNamespaceManage) CreateNamespace(dto NacosNamespaceCreateDto) (string, error) {
	urlStr := fmt.Sprintf("%s/nacos/v1/console/namespaces", m.NacosServerConfig.ApiUrl)
	params := map[string]interface{}{}
	if dto.CustomNamespaceId != nil {
		params["customNamespaceId"] = *dto.CustomNamespaceId
	}
	if dto.NamespaceName != nil {
		params["namespaceName"] = *dto.NamespaceName
	}
	if dto.NamespaceDesc != nil {
		params["namespaceDesc"] = *dto.NamespaceDesc
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

// 修改命名空间
func (m *NacosNamespaceManage) EditNamespace(dto NacosNamespaceEditDto) (string, error) {
	urlStr := fmt.Sprintf("%s/nacos/v1/console/namespaces", m.NacosServerConfig.ApiUrl)
	params := map[string]interface{}{}
	if dto.Namespace != nil {
		params["namespace"] = *dto.Namespace
	}
	if dto.NamespaceShowName != nil {
		params["namespaceShowName"] = *dto.NamespaceShowName
	}
	if dto.NamespaceDesc != nil {
		params["namespaceDesc"] = *dto.NamespaceDesc
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

// 删除命名空间
func (m *NacosNamespaceManage) DeleteNamespace(namespaceId string) (string, error) {
	urlStr := fmt.Sprintf("%s/nacos/v1/console/namespaces", m.NacosServerConfig.ApiUrl)
	if namespaceId == "" {
		return "", errors.New("命名空间id不能为空")
	}
	params := map[string]interface{}{
		"namespaceId": namespaceId,
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

func NewNacosNamespaceManage(config NacosServerConfig) *NacosNamespaceManage {
	return &NacosNamespaceManage{
		NacosServerConfig: config,
	}
}
