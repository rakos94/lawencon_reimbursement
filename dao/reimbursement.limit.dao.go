package dao

import "lawencon/reimbursement/model"

// ReimbursementLimitDao ...
type ReimbursementLimitDao interface {
	CreateReimbursementLimit(data *model.ReimbursementLimit) (*model.ReimbursementLimit, error)
	GetReimbursementLimitAll() ([]model.ReimbursementLimit, error)
	GetReimbursementLimitByID(id string) (model.ReimbursementLimit, error)
	GetReimbursementLimitByPersonID(id string) ([]model.ReimbursementLimit, error)
	UpdateReimbursementLimit(id string, data *model.ReimbursementLimit) (*model.ReimbursementLimit, error)
}
