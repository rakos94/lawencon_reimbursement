package model

// ReimbursementType ...
type ReimbursementType struct {
	BaseModels
	Name         string                 `gorm:"not null" json:"name"`
	Code         string                 `gorm:"not null" json:"code"`
	CategoryCode string                 `gorm:"not null" json:"category_code"`
	Category     *ReimbursementCategory `gorm:"foreignKey:CategoryCode;references:Code" json:"category,omitempty"`
	BaseCUModels
}

// TableName ...
func (ReimbursementType) TableName() string {
	return "reimbursement_types"
}
