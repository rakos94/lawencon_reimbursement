package model

// ReimbursementCategory ...
type ReimbursementCategory struct {
	BaseModels
	Name string              `gorm:"not null"`
	Code string              `gorm:"not null;uniqueIndex"`
	Type []ReimbursementType `gorm:"foreignKey:CategoryCode"`
	BaseCUModels
}

// TableName ...
func (ReimbursementCategory) TableName() string {
	return "reimbursement_categories"
}
