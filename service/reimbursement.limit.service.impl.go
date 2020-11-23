package service

import (
	"lawencon/reimbursement/dao"
	"lawencon/reimbursement/model"
)

var reimbursementLimit dao.ReimbursementLimitDao = dao.ReimbursementLimitDaoImpl{}

// ReimbursementLimitServiceImpl ...
type ReimbursementLimitServiceImpl struct{}

// CreateReimbursementLimit ...
func (ReimbursementLimitServiceImpl) CreateReimbursementLimit(data *model.ReimbursementLimit) (*model.ReimbursementLimit, error) {
	return reimbursementLimit.CreateReimbursementLimit(data)
}

// GetReimbursementLimitAll ...
func (ReimbursementLimitServiceImpl) GetReimbursementLimitAll() ([]model.ReimbursementLimit, error) {
	return reimbursementLimit.GetReimbursementLimitAll()
}

// GetReimbursementLimitByID ...
func (ReimbursementLimitServiceImpl) GetReimbursementLimitByID(id string) (model.ReimbursementLimit, error) {
	return reimbursementLimit.GetReimbursementLimitByID(id)
}

// GetReimbursementLimitByPersonID ...
func (ReimbursementLimitServiceImpl) GetReimbursementLimitByPersonID(id string) ([]model.ReimbursementLimit, error) {
	return reimbursementLimit.GetReimbursementLimitByPersonID(id)
}

// UpdateReimbursementLimit ...
func (ReimbursementLimitServiceImpl) UpdateReimbursementLimit(id string, data *model.ReimbursementLimit) (*model.ReimbursementLimit, error) {
	return reimbursementLimit.UpdateReimbursementLimit(id, data)
}
