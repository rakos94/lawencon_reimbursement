package request

import "lawencon/reimbursement/datatype"

// ReimbursementApprovedRequest ...
type ReimbursementApprovedRequest struct {
	ReimbursementProcessID string         `json:"reimbursement_process_id" validate:"required"`
	Reason                 string         `json:"reason" validate:"required"`
	StatusCode             string         `json:"status_code" validate:"required"`
	PaidDate               *datatype.Date `json:"paid_date"`
	PaidType               *string        `json:"paid_type"`
	ApprovedAmount         *float64       `json:"approved_amount"`
}
