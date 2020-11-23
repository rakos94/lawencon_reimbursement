package dao

import "lawencon/reimbursement/model"

// ReimbursementProcessDao ...
type ReimbursementProcessDao interface {
	CreateReimbursementProcess(data *model.ReimbursementProcess) (*model.ReimbursementProcess, error)
	GetReimbursementProcessAll() ([]model.ReimbursementProcess, error)
	GetReimbursementProcessByID(id string) (model.ReimbursementProcess, error)
	GetReimbursementProcessByReimbursementID(id string) (model.ReimbursementProcess, error)
	UpdateReimbursementProcess(id string, data *model.ReimbursementProcess) (*model.ReimbursementProcess, error)
}
