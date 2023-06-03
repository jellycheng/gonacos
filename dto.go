package gonacos

type NacosServerConfig struct {
	ApiUrl       string //http://127.0.0.1:8848
	BeatInterval int64  // 心跳值,单位毫秒，5 * 1000
}

type NacosConfigDto struct {
	Tenant *string `json:"tenant"` //可选，租户信息，对应 Nacos 的命名空间ID字段
	DataId *string `json:"dataId"` //必须，配置ID
	Group  *string `json:"group"`  //必须，配置分组
}
