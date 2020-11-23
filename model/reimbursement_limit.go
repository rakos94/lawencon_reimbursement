package model

// ReimbursementLimit ...
type ReimbursementLimit struct {
	BaseModels
	EmployeeID        string `gorm:"not null"`
	ReimbursementCode string `gorm:"not null"`
	TotalAmount       int64  `gorm:"not null;type:numeric"`
	UsedAmount        int64  `gorm:"not null;type:numeric;default:0"`
	BaseCUModels
}

// TableName ...
func (ReimbursementLimit) TableName() string {
	return "reimbursement_limit"
}
