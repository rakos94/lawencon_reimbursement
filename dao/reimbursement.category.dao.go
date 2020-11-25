package dao

import "lawencon/reimbursement/model"

// ReimbursementCategoryDao ...
type ReimbursementCategoryDao interface {
	CreateReimbursementCategory(data *model.ReimbursementCategory) (*model.ReimbursementCategory, error)
	GetReimbursementCategoryAll() ([]model.ReimbursementCategory, error)
	GetReimbursementCategoryByID(id string) (model.ReimbursementCategory, error)
	GetReimbursementCategoryByCode(code string) (model.ReimbursementCategory, error)
	UpdateReimbursementCategory(id string, data *model.ReimbursementCategory) (*model.ReimbursementCategory, error)
}
