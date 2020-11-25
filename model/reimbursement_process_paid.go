package model

import "lawencon/reimbursement/datatype"

// ReimbursementProcessPaid ...
type ReimbursementProcessPaid struct {
	BaseModels
	ReimbursementProcessID string        `gorm:"not null" json:"reimbursement_process_id"`
	PaidDate               datatype.Date `gorm:"not null;type:date" json:"paid_date"`
	PaidType               string        `gorm:"not null" json:"paid_type"`
	ApprovedAmount         float64       `gorm:"not null" json:"approved_amount"`
	BaseCUModels
}

// TableName ...
func (ReimbursementProcessPaid) TableName() string {
	return "reimbursement_process_paid"
}
