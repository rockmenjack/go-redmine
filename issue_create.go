package redmine

type IssueCreate struct {
	Subject      string         `json:"subject,omitempty"`
	Description  string         `json:"description,omitempty"`
	ProjectId    string         `json:"project_id,omitempty"`
	TrackerId    string         `json:"tracker_id,omitempty"`
	ParentId     int            `json:"parent_issue_id,omitempty"`
	PriorityId   int            `json:"priority_id,omitempty"`
	AssignedToId int            `json:"assigned_to_id,omitempty"`
	DueDate      string         `json:"due_date,omitempty"`
	CustomFields []*CustomField `json:"custom_fields,omitempty"`
}
