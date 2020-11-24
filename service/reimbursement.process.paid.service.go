package service

import (
	"lawencon/reimbursement/model"

	"gorm.io/gorm"
)

// ReimbursementProcessPaidService ...
type ReimbursementProcessPaidService interface {
	CreateReimbursementProcessPaid(data *model.ReimbursementProcessPaid) (*model.ReimbursementProcessPaid, error)
	GetReimbursementProcessPaidAll() ([]model.ReimbursementProcessPaid, error)
	GetReimbursementProcessPaidByID(id string) (model.ReimbursementProcessPaid, error)
	GetReimbursementProcessPaidByReimbursementProcessID(id string) (model.ReimbursementProcessPaid, error)

	// DB TRANSACTION
	TxCreateReimbursementProcessPaid(data *model.ReimbursementProcessPaid, tx *gorm.DB) (*model.ReimbursementProcessPaid, error)
}
