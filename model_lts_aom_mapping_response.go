package entity

type CreateAomMappingRulesResponse struct {
	Body           *[]AomMappingRuleResp `json:"body,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

type AomMappingRuleResp struct {

	// 项目id
	ProjectId string `json:"project_id"`

	// 接入规则名称
	RuleName string `json:"rule_name"`

	// 接入规则id
	RuleId string `json:"rule_id"`

	RuleInfo *AomMappingRuleInfo `json:"rule_info"`
}