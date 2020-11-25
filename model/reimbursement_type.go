package model

// ReimbursementType ...
type ReimbursementType struct {
	BaseModels
	CategoryCode string                 `gorm:"not null" json:"category_code"`
	Code         string                 `gorm:"not null;unique;uniqueIndex" json:"code"`
	Name         string                 `gorm:"not null" json:"name"`
	Category     *ReimbursementCategory `gorm:"foreignKey:CategoryCode;references:Code" json:"category,omitempty"`
	BaseCUModels
}

// TableName ...
func (ReimbursementType) TableName() string {
	return "reimbursement_types"
}
