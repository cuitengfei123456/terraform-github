package entity

type StructTemplateRequest struct {
	Content      string           `json:"content"`
	LogGroupId   string           `json:"log_group_id"`
	ParseType    string           `json:"parse_type"`
	TemplateId   string           `json:"template_id"`
	TemplateType string           `json:"template_type"`
	TemplateName string           `json:"template_name"`
	LogStreamId  string           `json:"log_stream_id"`
	ProjectId    string           `json:"project_id"`
	RegexRules   *string          `json:"regex_rules,omitempty"`
	Layers       *int             `json:"layers,omitempty"`
	Tokenizer    *string          `json:"tokenizer,omitempty"`
	LogFormat    *string          `json:"log_format,omitempty"`
	DemoFields   []DemoFieldsInfo `json:"demo_fields"`
	TagFields    []TagFieldsInfo  `json:"tag_fields"`
}

type DemoFieldsInfo struct {
	IsAnalysis      bool   `json:"isAnalysis"`
	Content         string `json:"content,omitempty"`
	FieldName       string `json:"fieldName,omitempty"`
	Type            string `json:"type"`
	UserDefinedName string `json:"userDefinedName,omitempty"`
	Index           int    `json:"index,omitempty"`
}

type TagFieldsInfo struct {
	FieldName  string  `json:"fieldName"`
	Type       string  `json:"type"`
	Content    *string `json:"content,omitempty"`
	IsAnalysis *bool   `json:"isAnalysis,omitempty"`
}
