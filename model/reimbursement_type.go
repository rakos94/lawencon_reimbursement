package model

// ReimbursementType ...
type ReimbursementType struct {
	BaseModels
	Name         string                `gorm:"not null"`
	Code         string                `gorm:"not null"`
	CategoryCode string                `gorm:"not null"`
	Category     ReimbursementCategory `gorm:"foreignKey:CategoryCode"`
	BaseCUModels
}

// TableName ...
func (ReimbursementType) TableName() string {
	return "reimbursement_type"
}
