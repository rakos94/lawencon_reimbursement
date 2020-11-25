package model

import "lawencon/reimbursement/datatype"

// ReimbursementProcess ...
type ReimbursementProcess struct {
	BaseModels
	ReimbursementRequestID   string                    `gorm:"not null" json:"reimbursement_request_id"`
	PeriodStart              datatype.Date             `gorm:"not null;type:date" json:"period_start"`
	PeriodEnd                datatype.Date             `gorm:"not null;type:date" json:"period_end"`
	ReimbursementProcessPaid *ReimbursementProcessPaid `gorm:"foreignKey:ReimbursementProcessID" json:"reimbursement_process_paid,omitempty"`
	ReimbursementRequest     *ReimbursementRequest     `gorm:"foreignKey:ReimbursementRequestID" json:"reimbursement_request,omitempty"`
	BaseCUModels
}

// TableName ...
func (ReimbursementProcess) TableName() string {
	return "reimbursement_process"
}
