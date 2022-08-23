package entity

type ServiceStageInstanceRequest struct {
	// 应用组件实例名称
	Name string `json:"name"`

	//环境ID
	EnviromentId string `json:"environment_id"`

	FlavorId *string `json:"flavor_id,omitempty"`

	// 实例副本数
	Replice int32 `json:"replica"`

	//组件部署件。Key为组件component_name, 对于Docker多容器场景，Key为容器名称。
	Artifacts map[string]interface{} `json:"artifacts"`

	//应用组件版本号，满足版本语义，如1.0.0。。
	Version string `json:"version"`

	// 应用配置，环境变量等， 如{“env”：[{“name”：“log-level”：“warn”}]},默认空。
	Configuration *interface{} `json:"configuration,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 访问方式
	ExternalAccesses *[]ExternalAccessesCreate `json:"external_accesses,omitempty"`

	// 部署资源
	ReferResources []ReferResourceCreate `json:"refer_resources"`
}

type ExternalAccessesCreate struct {

	// 协议，支持http、https。
	Protocol string `json:"protocol"`

	// 访问地址
	Address string `json:"address"`

	// 端口号
	ForwardPort int32 `json:"forward_port"`
}

type ReferResourceCreate struct {

	//资源ID。
	Id string `json:"id"`

	Type *string `json:"type"`

	// 应用别名，doc时才提供，支持“distributed_session”、“distributed_cache”、“distributed_session，distributed_cache”，默认值是“distributed_session，distributed_cache”。
	ReferAlias *string `json:"refer_alias,omitempty"`

	// 引用资源参数
	Parameters *interface{} `json:"parameters,omitempty"`
}

type InstanceAction struct {
	Action *string `json:"action"`

	Parameters *InstanceActionParameters `json:"parameters,omitempty"`
}

type InstanceActionParameters struct {

	//实例数，在scale操作时提供。
	Replica *int32 `json:"replica,omitempty"`

	//ECS ID列表，指定虚机扩容时部署的ECS主机。
	Hosts []string `json:"hosts,omitempty"`

	// 版本号，在rollback操作时提供，通过查询快照接口获取。
	Version *string `json:"version,omitempty"`
}
