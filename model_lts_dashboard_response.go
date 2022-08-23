package entity

type ReadDashBoardResp struct {
	Results []DashBoard `json:"results"`
}

type DashBoard struct {
	ProjectId          string           `json:"project_id"`
	Id                 string           `json:"id"`
	GroupName          string           `json:"group_name"`
	Title              string           `json:"title"`
	Charts             []DashboardChars `json:"charts"`
	Filters            []interface{}    `json:"filters"`
	LastUpdateTime     int              `json:"last_update_time"`
	UseSystemTemplate  bool             `json:"useSystemTemplate"`
}

type DashboardChars struct {
	Width   int                    `json:"width"`
	Height  int                    `json:"height"`
	XPos    int                    `json:"x_pos"`
	YPos    int                    `json:"y_pos"`
	ChartId string                 `json:"chart_id"`
	Chart   map[string]interface{} `json:"chart"`
}