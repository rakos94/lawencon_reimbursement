package service

import (
	"lawencon/reimbursement/dao"
	"lawencon/reimbursement/model"
)

var reimbursementCategory dao.ReimbursementCategoryDao = dao.ReimbursementCategoryDaoImpl{}

// ReimbursementCategoryServiceImpl ...
type ReimbursementCategoryServiceImpl struct{}

// CreateReimbursementCategory ...
func (ReimbursementCategoryServiceImpl) CreateReimbursementCategory(data *model.ReimbursementCategory) (*model.ReimbursementCategory, error) {
	return reimbursementCategory.CreateReimbursementCategory(data)
}

// GetReimbursementCategoryAll ...
func (ReimbursementCategoryServiceImpl) GetReimbursementCategoryAll() ([]model.ReimbursementCategory, error) {
	return reimbursementCategory.GetReimbursementCategoryAll()
}

// GetReimbursementCategoryByID ...
func (ReimbursementCategoryServiceImpl) GetReimbursementCategoryByID(id string) (model.ReimbursementCategory, error) {
	return reimbursementCategory.GetReimbursementCategoryByID(id)
}

// GetReimbursementCategoryByCode ...
func (ReimbursementCategoryServiceImpl) GetReimbursementCategoryByCode(code string) (model.ReimbursementCategory, error) {
	return reimbursementCategory.GetReimbursementCategoryByCode(code)
}

// UpdateReimbursementCategory ...
func (ReimbursementCategoryServiceImpl) UpdateReimbursementCategory(id string, data *model.ReimbursementCategory) (*model.ReimbursementCategory, error) {
	return reimbursementCategory.UpdateReimbursementCategory(id, data)
}
