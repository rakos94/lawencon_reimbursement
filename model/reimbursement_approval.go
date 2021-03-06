package model

// ReimbursementApproval ...
type ReimbursementApproval struct {
	BaseModels
	ReimbursementRequestID string                      `gorm:"not null" json:"reimbursement_request_id"`
	Reason                 string                      `gorm:"not null" json:"reason"`
	StatusCode             string                      `gorm:"not null" json:"status_code"`
	Status                 *ReimbursementRequestStatus `gorm:"foreignKey:StatusCode;references:StatusCode" json:"status,omitempty"`
	BaseCUModels
}

// TableName ...
func (ReimbursementApproval) TableName() string {
	return "reimbursement_approval"
}
