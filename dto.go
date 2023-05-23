package gonacos

type NacosServerConfig struct {
	ApiUrl       string //http://127.0.0.1:8848
	BeatInterval int64  // 心跳值,单位毫秒，5 * 1000
}

// 包含实例的注册、修改、删除参数
type NacosInstanceDto struct {
	Ip          *string  `json:"ip"`          //必须，服务实例IP
	Port        *int64   `json:"port"`        //必须，服务实例port
	NamespaceId *string  `json:"namespaceId"` //可选，命名空间ID
	Weight      *float64 `json:"weight"`      //可选，权重
	Enable      *bool    `json:"enabled"`     //可选，是否上线
	Healthy     *bool    `json:"healthy"`     //可选，是否健康
	Metadata    *string  `json:"metadata"`    //可选，扩展信息,json字符串
	ClusterName *string  `json:"clusterName"` //可选，集群名
	ServiceName *string  `json:"serviceName"` //必须，服务名
	GroupName   *string  `json:"groupName"`   //可选,默认DEFAULT_GROUP，分组名
	Ephemeral   *bool    `json:"ephemeral"`   //可选，是否临时实例
}

type BeatDto struct {
	ServiceName *string `json:"serviceName"` //必须，服务名
	Ip          *string `json:"ip"`          //必须，服务实例IP
	Port        *int64  `json:"port"`        // 必须，服务实例PORT
	NamespaceId *string `json:"namespaceId"` //可选，命名空间ID
	GroupName   *string `json:"groupName"`   //可选，分组名
	Ephemeral   *bool   `json:"ephemeral"`   //可选，是否临时实例
	Beat        *string `json:"beat"`        //必须，实例心跳内容，json字符串
}

// 服务创建、删除、修改
type NacosServiceDto struct {
	ServiceName      *string  `json:"serviceName"`      //必须，服务名
	GroupName        *string  `json:"groupName"`        //可选,默认DEFAULT_GROUP，分组名
	NamespaceId      *string  `json:"namespaceId"`      //可选，命名空间ID
	ProtectThreshold *float64 `json:"protectThreshold"` //可选，保护阈值,取值0到1,默认0
	Metadata         *string  `json:"metadata"`         //可选，元数据
	Selector         *string  `json:"selector"`         //可选，访问策略,json字符串
}
