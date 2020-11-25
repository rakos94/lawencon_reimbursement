package service

import (
	"errors"
	"lawencon/reimbursement/dao"
	"lawencon/reimbursement/model"
)

var reimbursementLimit dao.ReimbursementLimitDao = dao.ReimbursementLimitDaoImpl{}

// ReimbursementLimitServiceImpl ...
type ReimbursementLimitServiceImpl struct {
	TypeService ReimbursementTypeService
}

// CreateReimbursementLimit ...
func (r ReimbursementLimitServiceImpl) CreateReimbursementLimit(data *model.ReimbursementLimit) (*model.ReimbursementLimit, error) {
	_, err := r.TypeService.GetReimbursementTypeByCode(data.ReimbursementTypeCode)
	if err != nil {
		return nil, errors.New("Reimbursement Type Code not exist")
	}
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

// GetReimbursementLimitByEmployeeID ...
func (ReimbursementLimitServiceImpl) GetReimbursementLimitByEmployeeID(id string) ([]model.ReimbursementLimit, error) {
	return reimbursementLimit.GetReimbursementLimitByEmployeeID(id)
}

// UpdateReimbursementLimit ...
func (r ReimbursementLimitServiceImpl) UpdateReimbursementLimit(id string, data *model.ReimbursementLimit) (*model.ReimbursementLimit, error) {
	_, err := r.TypeService.GetReimbursementTypeByCode(data.ReimbursementTypeCode)
	if err != nil {
		return nil, errors.New("Reimbursement Type Code not exist")
	}
	return reimbursementLimit.UpdateReimbursementLimit(id, data)
}
