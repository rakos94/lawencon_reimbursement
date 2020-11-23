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

// GetReimbursementTypeByCategoryID ...
func (ReimbursementTypeServiceImpl) GetReimbursementTypeByCategoryID(id string) ([]model.ReimbursementType, error) {
	return reimbursementType.GetReimbursementTypeByCategoryID(id)
}

// UpdateReimbursementType ...
func (ReimbursementTypeServiceImpl) UpdateReimbursementType(id string, data *model.ReimbursementType) (*model.ReimbursementType, error) {
	return reimbursementType.UpdateReimbursementType(id, data)
}
