package model

// ReimbursementCategory ...
type ReimbursementCategory struct {
	BaseModels
	Code string              `gorm:"not null;unique;uniqueIndex" json:"code"`
	Name string              `gorm:"not null" json:"name"`
	Type []ReimbursementType `gorm:"foreignKey:CategoryCode;references:Code" json:"reimbursement_type,omitempty"`
	BaseCUModels
}

// TableName ...
func (ReimbursementCategory) TableName() string {
	return "reimbursement_categories"
}
