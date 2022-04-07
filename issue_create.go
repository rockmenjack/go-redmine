package redmine

type IssueCreate struct {
	Subject      string         `json:"subject"`
	Description  string         `json:"description"`
	ProjectId    string         `json:"project_id"`
	TrackerId    string         `json:"tracker_id"`
	ParentId     int            `json:"parent_issue_id,omitempty"`
	PriorityId   int            `json:"priority_id,omitempty"`
	AssignedTo   string         `json:"assigned_to"`
	CustomFields []*CustomField `json:"custom_fields,omitempty"`
}
