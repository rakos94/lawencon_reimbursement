package service

import (
	"lawencon/reimbursement/dao"
	"lawencon/reimbursement/model"

	"gorm.io/gorm"
)

var reimbursementProcessPaid dao.ReimbursementProcessPaidDao = dao.ReimbursementProcessPaidDaoImpl{}

// ReimbursementProcessPaidServiceImpl ...
type ReimbursementProcessPaidServiceImpl struct{}

// CreateReimbursementProcessPaid ...
func (ReimbursementProcessPaidServiceImpl) CreateReimbursementProcessPaid(data *model.ReimbursementProcessPaid) (*model.ReimbursementProcessPaid, error) {
	return reimbursementProcessPaid.CreateReimbursementProcessPaid(data)
}

// GetReimbursementProcessPaidAll ...
func (ReimbursementProcessPaidServiceImpl) GetReimbursementProcessPaidAll() ([]model.ReimbursementProcessPaid, error) {
	return reimbursementProcessPaid.GetReimbursementProcessPaidAll()
}

// GetReimbursementProcessPaidByID ...
func (ReimbursementProcessPaidServiceImpl) GetReimbursementProcessPaidByID(id string) (model.ReimbursementProcessPaid, error) {
	return reimbursementProcessPaid.GetReimbursementProcessPaidByID(id)
}

// GetReimbursementProcessPaidByReimbursementProcessID ...
func (ReimbursementProcessPaidServiceImpl) GetReimbursementProcessPaidByReimbursementProcessID(id string) (model.ReimbursementProcessPaid, error) {
	return reimbursementProcessPaid.GetReimbursementProcessPaidByReimbursementProcessID(id)
}

// TxCreateReimbursementProcessPaid ...
func (ReimbursementProcessPaidServiceImpl) TxCreateReimbursementProcessPaid(data *model.ReimbursementProcessPaid, tx *gorm.DB) (*model.ReimbursementProcessPaid, error) {
	return reimbursementProcessPaid.TxCreateReimbursementProcessPaid(data, tx)
}
