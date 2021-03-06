package model

// ReimbursementLimit ...
type ReimbursementLimit struct {
	BaseModels
	EmployeeID            string             `gorm:"not null;uniqueIndex:idx_name" json:"employee_id"`
	ReimbursementTypeCode string             `gorm:"not null;uniqueIndex:idx_name" json:"reimbursement_type_code"`
	TotalAmount           float64            `gorm:"not null;type:numeric" json:"total_amount"`
	UsedAmount            float64            `gorm:"not null;type:numeric;default:0" json:"used_amount"`
	Type                  *ReimbursementType `gorm:"foreignKey:ReimbursementTypeCode;references:Code" json:"reimbursement_type,omitempty"`
	BaseCUModels
}

// TableName ...
func (ReimbursementLimit) TableName() string {
	return "reimbursement_limit"
}
