package service

import (
	"lawencon/reimbursement/dao"
	"lawencon/reimbursement/model"
)

var reimbursementType dao.ReimbursementTypeDao = dao.ReimbursementTypeDaoImpl{}

// ReimbursementTypeServiceImpl ...
type ReimbursementTypeServiceImpl struct{}

// CreateReimbursementType ...
func (ReimbursementTypeServiceImpl) CreateReimbursementType(data *model.ReimbursementType) (*model.ReimbursementType, error) {
	return reimbursementType.CreateReimbursementType(data)
}

// GetReimbursementTypeAll ...
func (ReimbursementTypeServiceImpl) GetReimbursementTypeAll() ([]model.ReimbursementType, error) {
	return reimbursementType.GetReimbursementTypeAll()
}

// GetReimbursementTypeByID ...
func (ReimbursementTypeServiceImpl) GetReimbursementTypeByID(id string) (model.ReimbursementType, error) {
	return reimbursementType.GetReimbursementTypeByID(id)
}

// GetReimbursementTypeByCode ...
func (ReimbursementTypeServiceImpl) GetReimbursementTypeByCode(code string) (model.ReimbursementType, error) {
	return reimbursementType.GetReimbursementTypeByCode(code)
}

// GetReimbursementTypeByCategoryCode ...
func (ReimbursementTypeServiceImpl) GetReimbursementTypeByCategoryCode(code string) ([]model.ReimbursementType, error) {
	return reimbursementType.GetReimbursementTypeByCategoryCode(code)
}

// UpdateReimbursementType ...
func (ReimbursementTypeServiceImpl) UpdateReimbursementType(id string, data *model.ReimbursementType) (*model.ReimbursementType, error) {
	return reimbursementType.UpdateReimbursementType(id, data)
}
