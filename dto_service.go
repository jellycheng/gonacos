package gonacos

// NacosServiceDto 服务创建、删除、修改、查询
type NacosServiceDto struct {
	ServiceName      *string  `json:"serviceName"`      //必须，服务名
	GroupName        *string  `json:"groupName"`        //可选,默认DEFAULT_GROUP，分组名
	NamespaceId      *string  `json:"namespaceId"`      //可选，命名空间ID
	ProtectThreshold *float64 `json:"protectThreshold"` //可选，保护阈值,取值0到1,默认0
	Metadata         *string  `json:"metadata"`         //可选，元数据
	Selector         *string  `json:"selector"`         //可选，访问策略,json字符串
}

type NacosListServiceReqDto struct {
	PageNo      *int64  `json:"pageNo"`      //必须，当前页码
	PageSize    *int64  `json:"pageSize"`    //必须，分页大小
	GroupName   *string `json:"groupName"`   //可选，分组名
	NamespaceId *string `json:"namespaceId"` //可选，命名空间ID
}
