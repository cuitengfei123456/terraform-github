package entity

type DashBoardResquest struct {
	// 日志组id
	LogGroupId string `json:"log_group_id"`

	// 目标日志组名称。
	LogGroupName string `json:"log_group_name"`

	// 日志流id
	LogStreamId string `json:"log_stream_id"`
	// 目标日志组名称。
	LogStreamName string `json:"log_stream_name"`

	TemplateTiltle []string `json:"template_title"`
	TemplateType   []string `json:"template_type"`
	GroupName      string   `json:"group_name"`
}
