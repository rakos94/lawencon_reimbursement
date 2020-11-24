package request

// ReimbursementApprovedRequest ...
type ReimbursementApprovedRequest struct {
	ReimbursementID string  `json:"reimbursement_id" validate:"required"`
	Reason          string  `json:"reason" validate:"required"`
	StatusCode      string  `json:"status_code" validate:"required"`
	PaidType        string  `json:"paid_type" validate:"required"`
	ApprovedAmount  float64 `json:"approved_amount" validate:"required"`
}
