package model

type ReimbursementRequestStatus struct {
	BaseModels
	StatusCode	string `gorm:"column:status_code; not null" json:"status_code"`
	Name 		string `gorm:"column:name; not null" json:"name"`
}
func (ReimbursementRequestStatus)TableName()string  {
	return "reimbursement_request_status"
}