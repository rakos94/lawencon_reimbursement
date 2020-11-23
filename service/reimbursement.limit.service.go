package service

import "lawencon/reimbursement/model"

// ReimbursementLimitService ...
type ReimbursementLimitService interface {
	CreateReimbursementLimit(data *model.ReimbursementLimit) (*model.ReimbursementLimit, error)
	GetReimbursementLimitAll() ([]model.ReimbursementLimit, error)
	GetReimbursementLimitByID(id string) (model.ReimbursementLimit, error)
	GetReimbursementLimitByEmployeeID(id string) ([]model.ReimbursementLimit, error)
	UpdateReimbursementLimit(id string, data *model.ReimbursementLimit) (*model.ReimbursementLimit, error)
}
