package dao

import (
	"lawencon/reimbursement/model"

	"gorm.io/gorm"
)

// ReimbursementProcessPaidDao ...
type ReimbursementProcessPaidDao interface {
	CreateReimbursementProcessPaid(data *model.ReimbursementProcessPaid) (*model.ReimbursementProcessPaid, error)
	GetReimbursementProcessPaidAll() ([]model.ReimbursementProcessPaid, error)
	GetReimbursementProcessPaidByID(id string) (model.ReimbursementProcessPaid, error)
	GetReimbursementProcessPaidByReimbursementProcessID(reimbursementProcessID string) (model.ReimbursementProcessPaid, error)

	// DB TRANSACTION
	TxCreateReimbursementProcessPaid(data *model.ReimbursementProcessPaid, tx *gorm.DB) (*model.ReimbursementProcessPaid, error)
}
