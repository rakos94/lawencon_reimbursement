package dao

import (
	"lawencon/reimbursement/model"

	"gorm.io/gorm"
)

// ReimbursementApprovalDao ...
type ReimbursementApprovalDao interface {
	CreateReimbursementApproval(data *model.ReimbursementApproval) (*model.ReimbursementApproval, error)
	GetReimbursementApprovalAll() ([]model.ReimbursementApproval, error)
	GetReimbursementApprovalByID(id string) (model.ReimbursementApproval, error)
	GetReimbursementApprovalByRequestID(requestID string) (model.ReimbursementApproval, error)

	// DB TRANSACTION
	TxCreateReimbursementApproval(data *model.ReimbursementApproval, tx *gorm.DB) (*model.ReimbursementApproval, error)
}
