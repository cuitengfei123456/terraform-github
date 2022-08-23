package entity

type ShowStructTemplateResponse struct {

	// 结构化字段
	DemoFields *[]StructFieldInfoReturn `json:"demoFields,omitempty"`

	// 关键词详细信息
	TagFields *[]StructTagFieldsInfo `json:"tagFields,omitempty"`

	// 示例日志
	DemoLog *string `json:"demoLog,omitempty"`

	// 测试
	DemoLabel *string `json:"demoLabel,omitempty"`

	// id
	Id string `json:"id,omitempty"`

	// 日志组ID
	LogGroupId *string `json:"logGroupId,omitempty"`

	Rule *ShowStructTemplateRule `json:"rule,omitempty"`

	// 日志流ID
	LogStreamId *string `json:"logStreamId,omitempty"`

	// 项目ID
	ProjectId *string `json:"projectId,omitempty"`

	// 测试
	TemplateName *string `json:"templateName,omitempty"`

	// 为了兼容前台数据格式
	Regex *string `json:"regex,omitempty"`
	HttpStatusCode int `json:"-"`
}

type StructFieldInfoReturn struct {

	// 字段名称
	FieldName *string `json:"fieldName,omitempty"`

	// 字段数据类型
	Type *string `json:"type,omitempty"`

	// 字段内容
	Content *string `json:"conteny,omitempty"`

	// 结构化方式
	IsAnalysis *bool `json:""isAnalysis,omitempty`

	// 序号
	Index *int32 `json:"index,omitempty"`
}

type StructTagFieldsInfo struct {

    // 字段名称
	FieldName *string `json:"fieldName,omitempty"`

	// 字段类型
	Type *string `json:"type,omitempty"`

	// 内容
	Content *string `json:"conteny,omitempty"`

	// 是否解析
	IsAnalysis *bool `json:"isAnalysis,omitempty"`

	// 字段名称
	Index *int32 `json:"index,omitempty"`
}

type ShowStructTemplateRule struct {

	// 测试
	Param *string `json:"param,omitempty"`

	// 结构化类型
	Type *string `json:"type,omitempty"`

}

type DeleteStructTemplateReqBody struct {

	// 结构化规则ID
	Id string `json:"id"`
}