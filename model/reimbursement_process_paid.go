package model

import "lawencon/reimbursement/datatype"

// ReimbursementProcessPaid ...
type ReimbursementProcessPaid struct {
	BaseModels
	ReimbursementProcessID string        `gorm:"not null"`
	PaidDate               datatype.Date `gorm:"not null;type:date"`
	PaidType               string        `gorm:"not null;type:date"`
	ApprovedAmount         float64       `gorm:"not null"`
	BaseCUModels
}

// TableName ...
func (ReimbursementProcessPaid) TableName() string {
	return "reimbursement_process_paid"
}
