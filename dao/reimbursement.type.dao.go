package dao

import "lawencon/reimbursement/model"

// ReimbursementTypeDao ...
type ReimbursementTypeDao interface {
	CreateReimbursementType(data *model.ReimbursementType) (*model.ReimbursementType, error)
	GetReimbursementTypeAll() ([]model.ReimbursementType, error)
	GetReimbursementTypeByID(id string) (model.ReimbursementType, error)
	GetReimbursementTypeByCategoryCode(code string) ([]model.ReimbursementType, error)
	UpdateReimbursementType(id string, data *model.ReimbursementType) (*model.ReimbursementType, error)
}
