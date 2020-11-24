package dao

import (
	"lawencon/reimbursement/model"

	"gorm.io/gorm"
)

// ReimbursementApprovalDaoImpl ...
type ReimbursementApprovalDaoImpl struct{}

// CreateReimbursementApproval ...
func (ReimbursementApprovalDaoImpl) CreateReimbursementApproval(data *model.ReimbursementApproval) (*model.ReimbursementApproval, error) {
	result := g.Create(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}

// GetReimbursementApprovalAll ...
func (ReimbursementApprovalDaoImpl) GetReimbursementApprovalAll() ([]model.ReimbursementApproval, error) {
	m := []model.ReimbursementApproval{}
	result := g.Find(&m)
	if result.Error != nil {
		return nil, result.Error
	}
	return m, nil
}

// GetReimbursementApprovalByID ...
func (ReimbursementApprovalDaoImpl) GetReimbursementApprovalByID(id string) (model.ReimbursementApproval, error) {
	m := model.ReimbursementApproval{}
	result := g.Where("id", id).First(&m)
	if result.Error != nil {
		return m, result.Error
	}
	return m, nil
}

// GetReimbursementApprovalByReimbursementID ...
func (ReimbursementApprovalDaoImpl) GetReimbursementApprovalByReimbursementID(id string) ([]model.ReimbursementApproval, error) {
	m := []model.ReimbursementApproval{}
	result := g.Where("reimbursement_id", id).Find(&m)
	if result.Error != nil {
		return m, result.Error
	}
	return m, nil
}

// TxCreateReimbursementApproval ...
func (ReimbursementApprovalDaoImpl) TxCreateReimbursementApproval(data *model.ReimbursementApproval, tx *gorm.DB) (*model.ReimbursementApproval, error) {
	result := tx.Create(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}
