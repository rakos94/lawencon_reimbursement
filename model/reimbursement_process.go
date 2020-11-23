package model

// ReimbursementProcess ...
type ReimbursementProcess struct {
	BaseModels
	ReimbursementID string `gorm:"not null"`
	PeriodStart     string `gorm:"not null;type:date"`
	PeriodEnd       string `gorm:"not null;type:date"`
	PaidDate        string `gorm:"not null;type:date"`
	PaidType        string `gorm:"not null;type:date"`
	BaseCUModels
}

// TableName ...
func (ReimbursementProcess) TableName() string {
	return "reimbursement_process"
}
