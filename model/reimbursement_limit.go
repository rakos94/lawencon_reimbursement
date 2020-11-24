package model

// ReimbursementLimit ...
type ReimbursementLimit struct {
	BaseModels
	EmployeeID        string  `gorm:"not null"`
	ReimbursementCode string  `gorm:"not null"`
	TotalAmount       float64 `gorm:"not null;type:numeric"`
	UsedAmount        float64 `gorm:"not null;type:numeric;default:0"`
	BaseCUModels
}

// TableName ...
func (ReimbursementLimit) TableName() string {
	return "reimbursement_limit"
}
