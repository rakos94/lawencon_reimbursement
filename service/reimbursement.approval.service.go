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
	GetReimbursementApprovalByReimbursementID(id string) ([]model.ReimbursementApproval, error)
	ApproveReimbursement(reimbursementID string, data request.ReimbursementApprovedRequest) (*model.ReimbursementApproval, error)
	// CancelReimbursement(id string, reason string) (*model.ReimbursementProcess, error)
}
