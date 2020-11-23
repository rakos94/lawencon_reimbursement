package model

import "lawencon/reimbursement/datatype"

// ReimbursementProcess ...
type ReimbursementProcess struct {
	BaseModels
	ReimbursementID string        `gorm:"not null"`
	PeriodStart     datatype.Date `gorm:"not null;type:date"`
	PeriodEnd       datatype.Date `gorm:"not null;type:date"`
	PaidDate        datatype.Date `gorm:"not null;type:date"`
	PaidType        string        `gorm:"not null;type:date"`
	BaseCUModels
}

// TableName ...
func (ReimbursementProcess) TableName() string {
	return "reimbursement_process"
}
