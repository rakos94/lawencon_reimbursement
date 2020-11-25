package model

// ReimbursementCategory ...
type ReimbursementCategory struct {
	BaseModels
	Name string              `gorm:"not null" json:"name"`
	Code string              `gorm:"not null;uniqueIndex" json:"code"`
	Type []ReimbursementType `gorm:"foreignKey:CategoryCode;references:Code" json:"reimbursement_type,omitempty"`
	BaseCUModels
}

// TableName ...
func (ReimbursementCategory) TableName() string {
	return "reimbursement_categories"
}
