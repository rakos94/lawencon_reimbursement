package model

import "lawencon/reimbursement/datatype"

// ReimbursementProcess ...
type ReimbursementProcess struct {
	BaseModels
	ReimbursementRequestID   string                   `gorm:"not null"`
	PeriodStart              datatype.Date            `gorm:"not null;type:date"`
	PeriodEnd                datatype.Date            `gorm:"not null;type:date"`
	ReimbursementProcessPaid ReimbursementProcessPaid `gorm:"foreignKey:ReimbursementProcessID"`
	BaseCUModels
}

// TableName ...
func (ReimbursementProcess) TableName() string {
	return "reimbursement_process"
}
