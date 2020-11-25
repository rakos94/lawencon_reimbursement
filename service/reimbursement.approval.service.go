package service

import (
	"lawencon/reimbursement/model"
	"lawencon/reimbursement/request"
)

// ReimbursementApprovalService ...
type ReimbursementApprovalService interface {
	CreateReimbursementApproval(data *model.ReimbursementApproval) (*model.ReimbursementApproval, error)
	GetReimbursementApprovalAll() ([]model.ReimbursementApproval, error)
	GetReimbursementApprovalByID(id string) (model.ReimbursementApproval, error)
	GetReimbursementApprovalByRequestID(requestID string) (model.ReimbursementApproval, error)
	ApproveReimbursement(processID string, data *request.ReimbursementApprovedRequest) (*model.ReimbursementApproval, error)
	CancelReimbursement(processID string, data *request.ReimbursementApprovedRequest) (*model.ReimbursementApproval, error)
}
