package gonacos

// NacosNamespaceCreateDto 命名空间创建
type NacosNamespaceCreateDto struct {
	CustomNamespaceId *string `json:"customNamespaceId"` //必须，命名空间ID
	NamespaceName     *string `json:"namespaceName"`     //必须，命名空间名
	NamespaceDesc     *string `json:"namespaceDesc"`     //可选，命名空间描述

}

// NacosNamespaceEditDto 命名空间修改
type NacosNamespaceEditDto struct {
	Namespace         *string `json:"namespace"`         //必须，命名空间ID
	NamespaceShowName *string `json:"namespaceShowName"` //必须，命名空间名
	NamespaceDesc     *string `json:"namespaceDesc"`     //可选，命名空间描述，不传则置为空

}

type NacosNamespaceDataDto struct {
	Namespace         string      `json:"namespace"`
	NamespaceShowName string      `json:"namespaceShowName"`
	NamespaceDesc     interface{} `json:"namespaceDesc"`
	Quota             int         `json:"quota"`
	ConfigCount       int         `json:"configCount"`
	Type              int         `json:"type"`
}

// NacosNamespaceListRespDto 命名空间列表
type NacosNamespaceListRespDto struct {
	Code    int                     `json:"code"`
	Message interface{}             `json:"message"`
	Data    []NacosNamespaceDataDto `json:"data"`
}
