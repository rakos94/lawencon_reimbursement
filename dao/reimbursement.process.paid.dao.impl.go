package dao

import (
	"lawencon/reimbursement/model"

	"gorm.io/gorm"
)

// ReimbursementProcessPaidDaoImpl ...
type ReimbursementProcessPaidDaoImpl struct{}

// CreateReimbursementProcessPaid ...
func (ReimbursementProcessPaidDaoImpl) CreateReimbursementProcessPaid(data *model.ReimbursementProcessPaid) (*model.ReimbursementProcessPaid, error) {
	result := g.Create(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}

// GetReimbursementProcessPaidAll ...
func (ReimbursementProcessPaidDaoImpl) GetReimbursementProcessPaidAll() ([]model.ReimbursementProcessPaid, error) {
	m := []model.ReimbursementProcessPaid{}
	result := g.Find(&m)
	if result.Error != nil {
		return nil, result.Error
	}
	return m, nil
}

// GetReimbursementProcessPaidByID ...
func (ReimbursementProcessPaidDaoImpl) GetReimbursementProcessPaidByID(id string) (model.ReimbursementProcessPaid, error) {
	m := model.ReimbursementProcessPaid{}
	result := g.Where("id", id).First(&m)
	if result.Error != nil {
		return m, result.Error
	}
	return m, nil
}

// GetReimbursementProcessPaidByReimbursementProcessID ...
func (ReimbursementProcessPaidDaoImpl) GetReimbursementProcessPaidByReimbursementProcessID(reimbursementProcessID string) (model.ReimbursementProcessPaid, error) {
	m := model.ReimbursementProcessPaid{}
	result := g.Where("reimbursement_process_id", reimbursementProcessID).First(&m)
	if result.Error != nil {
		return m, result.Error
	}
	return m, nil
}

// TxCreateReimbursementProcessPaid ...
func (ReimbursementProcessPaidDaoImpl) TxCreateReimbursementProcessPaid(data *model.ReimbursementProcessPaid, tx *gorm.DB) (*model.ReimbursementProcessPaid, error) {
	result := tx.Create(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}
