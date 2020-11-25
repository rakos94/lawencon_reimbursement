package service

import "lawencon/reimbursement/model"

// ReimbursementProcessService ...
type ReimbursementProcessService interface {
	CreateReimbursementProcess(data *model.ReimbursementProcess) (*model.ReimbursementProcess, error)
	GetReimbursementProcessAll() ([]model.ReimbursementProcess, error)
	GetReimbursementProcessByID(id string) (model.ReimbursementProcess, error)
	GetReimbursementProcessByRequestID(id string) (model.ReimbursementProcess, error)
	UpdateReimbursementProcess(id string, data *model.ReimbursementProcess) (*model.ReimbursementProcess, error)
}
