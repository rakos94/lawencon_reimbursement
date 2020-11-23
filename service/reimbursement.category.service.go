package service

import "lawencon/reimbursement/model"

// ReimbursementCategoryService ...
type ReimbursementCategoryService interface {
	CreateReimbursementCategory(data *model.ReimbursementCategory) (*model.ReimbursementCategory, error)
	GetReimbursementCategoryAll() ([]model.ReimbursementCategory, error)
	GetReimbursementCategoryByID(id string) (model.ReimbursementCategory, error)
	UpdateReimbursementCategory(id string, data *model.ReimbursementCategory) (*model.ReimbursementCategory, error)
}
