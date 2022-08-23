package entity

type ServiceStageInstanceResponse struct {
	JobId string `json:"job_id"`
}

type ShowInstanceDetailResponse struct {

	// 应用组件实例ID。
	Id *string `json:"id,omitempty"`

	// 应用组件实例名称。
	Name *string `json:"name,omitempty"`

	// 实例描述。
	Description *string `json:"description,omitempty"`

	// 应用组件环境ID。
	EnvironmentId *string `json:"environment_id,omitempty"`

	PlatformType *string `json:"platform_type,omitempty"`

	FlavorId *string `json:"flavor_id,omitempty"`

	// 组件部署件。Key为组件component_name，对于Docker多容器场景，Key为容器名称。
	Artifacts map[string]interface{} `json:"artifacts,omitempty"`

	// 应用组件版本号。
	Version *string `json:"version,omitempty"`

	// 应用组件配置，如环境变量。
	Configuration *interface{} `json:"configuration,omitempty"`

	// 创建人。
	Creator *string `json:"creator,omitempty"`

	// 创建时间。
	CreateTime *int64 `json:"create_time,omitempty"`
 
	// 修改时间。
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 访问方式列表。
	ExternalAccesses *[]ExternalAccesses `json:"external_accesses,omitempty"`

	// 部署资源列表。
	ReferResources *[]ReferResources `json:"refer_resources,omitempty"`

	StatusDetail   *InstanceStatusView `json:"status_detail,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

type ExternalAccesses struct {

	// ID。
	Id *string `json:"id,omitempty"`

	Protocol string `json:"protocol"`

	// 访问地址
    Address string `json:"address"`

	// 应用组件进程监听端口
	ForwardPort int32 `json:"forward_port"`

	Type *string `json:"type,omitempty"`

	Status *string `json:"status,omitempty"`

	// 创建时间。
	CreateTime *int64 `json:"create_time,omitempty"`
 
	// 修改时间。
	UpdateTime *int64 `json:"update_time,omitempty"`
}

type ReferResources struct {

	// 资源ID。
	Id *string `json:"id,omitempty"`

	Type *string `json:"type,omitempty"`

	// 应用别名，doc时才提供，支持“distributed_session”、“distributed_cache”、“distributed_session，distributed_cache”，默认值是“distributed_session，distributed_cache”。
	ReferAlias *string `json:"refer_alias,omitempty"`

	// 引用资源参数
	Parameters *interface{} `json:"parameters,omitempty"`
}

type InstanceStatusView struct {
	Status *string `json:"status,omitempty"`

	// 正常实例副本数。
	AvailableReplica *int32 `json:"available_replica,omitempty"`

	// 实例副本数。
	Replica *int32 `json:"replica,omitempty"`

	FailDetail *string `json:"fail_detail,omitempty"`

	// 最近Job ID。
	LastJobId *string `json:"last_job_id,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}
