package gonacos

import "encoding/json"

// NacosInstanceDto 包含实例的注册、修改、删除参数
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

// NacosInstanceListReqDto 查询服务下的实例列表
type NacosInstanceListReqDto struct {
	ServiceName *string `json:"serviceName"` //必须，服务名
	GroupName   *string `json:"groupName"`   //可选,分组名,默认DEFAULT_GROUP
	NamespaceId *string `json:"namespaceId"` //可选，命名空间ID，默认public
	HealthyOnly *bool   `json:"healthyOnly"` //可选，是否只返回健康实例,默认false
	Clusters    *string `json:"clusters"`    //可选，集群名，多个用逗号分隔，默认DEFAULT
}

type NacosInstanceHostDto struct {
	InstanceID                string          `json:"instanceId"`
	Ip                        string          `json:"ip"`
	Port                      int             `json:"port"`
	Weight                    float64         `json:"weight"`
	Healthy                   bool            `json:"healthy"`
	Enabled                   bool            `json:"enabled"`
	Ephemeral                 bool            `json:"ephemeral"`
	ClusterName               string          `json:"clusterName"`
	ServiceName               string          `json:"serviceName"`
	Metadata                  json.RawMessage `json:"metadata"`
	InstanceHeartBeatInterval int             `json:"instanceHeartBeatInterval"`
	InstanceHeartBeatTimeOut  int             `json:"instanceHeartBeatTimeOut"`
	IPDeleteTimeout           int             `json:"ipDeleteTimeout"`
}

type NacosInstanceListRespDto struct {
	Name                     string                 `json:"name"`
	GroupName                string                 `json:"groupName"`
	Clusters                 string                 `json:"clusters"`
	CacheMillis              int                    `json:"cacheMillis"`
	Hosts                    []NacosInstanceHostDto `json:"hosts"`
	LastRefTime              int64                  `json:"lastRefTime"`
	Checksum                 string                 `json:"checksum"`
	AllIPs                   bool                   `json:"allIPs"`
	ReachProtectionThreshold bool                   `json:"reachProtectionThreshold"`
	Valid                    bool                   `json:"valid"`
}

// NacosInstanceViewReqDto 查询实例详情：查询一个服务下个某个实例详情
type NacosInstanceViewReqDto struct {
	ServiceName *string `json:"serviceName"` //必须，服务名
	GroupName   *string `json:"groupName"`   //可选,分组名,默认DEFAULT_GROUP
	Ip          *string `json:"ip"`          //实例ip
	Port        *int    `json:"port"`        //实例端口
	NamespaceId *string `json:"namespaceId"` //可选，命名空间ID，默认public
	Clusters    *string `json:"clusters"`    //可选，集群名，多个用逗号分隔，默认DEFAULT
	HealthyOnly *bool   `json:"healthyOnly"` //可选，是否只返回健康实例,默认false
	Ephemeral   *bool   `json:"ephemeral"`   //可选，是否临时实例
}
