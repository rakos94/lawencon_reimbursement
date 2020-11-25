package service

import (
	"lawencon/reimbursement/dao"
	"lawencon/reimbursement/model"
)

var reimbursementProcess dao.ReimbursementProcessDao = dao.ReimbursementProcessDaoImpl{}

// ReimbursementProcessServiceImpl ...
type ReimbursementProcessServiceImpl struct{}

// CreateReimbursementProcess ...
func (ReimbursementProcessServiceImpl) CreateReimbursementProcess(data *model.ReimbursementProcess) (*model.ReimbursementProcess, error) {
	return reimbursementProcess.CreateReimbursementProcess(data)
}

// GetReimbursementProcessAll ...
func (ReimbursementProcessServiceImpl) GetReimbursementProcessAll() ([]model.ReimbursementProcess, error) {
	return reimbursementProcess.GetReimbursementProcessAll()
}

// GetReimbursementProcessByID ...
func (ReimbursementProcessServiceImpl) GetReimbursementProcessByID(id string) (model.ReimbursementProcess, error) {
	return reimbursementProcess.GetReimbursementProcessByID(id)
}

// GetReimbursementProcessByRequestID ...
func (ReimbursementProcessServiceImpl) GetReimbursementProcessByRequestID(id string) (model.ReimbursementProcess, error) {
	return reimbursementProcess.GetReimbursementProcessByRequestID(id)
}

// UpdateReimbursementProcess ...
func (ReimbursementProcessServiceImpl) UpdateReimbursementProcess(id string, data *model.ReimbursementProcess) (*model.ReimbursementProcess, error) {
	return reimbursementProcess.UpdateReimbursementProcess(id, data)
}
