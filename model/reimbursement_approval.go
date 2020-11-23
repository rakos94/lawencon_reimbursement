package model

// ReimbursementApproval ...
type ReimbursementApproval struct {
	BaseModels
	ReimbursementID string                     `gorm:"not null"`
	Reason          string                     `gorm:"not null;type:text"`
	StatusCode      string                     `gorm:"not null"`
	Status          ReimbursementRequestStatus `gorm:"foreignKey:StatusCode"`
	BaseCUModels
}

// TableName ...
func (ReimbursementApproval) TableName() string {
	return "reimbursement_approval"
}
