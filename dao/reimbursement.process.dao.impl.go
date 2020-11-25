package dao

import (
	"lawencon/reimbursement/datatype"
	"lawencon/reimbursement/model"
	"time"
)

// ReimbursementProcessDaoImpl ...
type ReimbursementProcessDaoImpl struct{}

// CreateReimbursementProcess ...
func (ReimbursementProcessDaoImpl) CreateReimbursementProcess(data *model.ReimbursementProcess) (*model.ReimbursementProcess, error) {
	now := time.Now()
	data.PeriodStart = datatype.Date(now)
	data.PeriodEnd = datatype.Date(now.AddDate(0, 0, 7))
	result := g.Create(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}

// GetReimbursementProcessAll ...
func (ReimbursementProcessDaoImpl) GetReimbursementProcessAll() ([]model.ReimbursementProcess, error) {
	m := []model.ReimbursementProcess{}
	result := g.Find(&m)
	if result.Error != nil {
		return nil, result.Error
	}
	return m, nil
}

// GetReimbursementProcessByID ...
func (ReimbursementProcessDaoImpl) GetReimbursementProcessByID(id string) (model.ReimbursementProcess, error) {
	m := model.ReimbursementProcess{}
	result := g.Where("id", id).First(&m)
	if result.Error != nil {
		return m, result.Error
	}
	return m, nil
}

// GetReimbursementProcessByRequestID ...
func (ReimbursementProcessDaoImpl) GetReimbursementProcessByRequestID(id string) (model.ReimbursementProcess, error) {
	m := model.ReimbursementProcess{}
	result := g.Where("reimbursement_request_id", id).Find(&m)
	if result.Error != nil {
		return m, result.Error
	}
	return m, nil
}

// UpdateReimbursementProcess ...
func (ReimbursementProcessDaoImpl) UpdateReimbursementProcess(id string, data *model.ReimbursementProcess) (*model.ReimbursementProcess, error) {
	m := &model.ReimbursementProcess{}
	first := g.Where("id", id).First(&m)
	if first.Error != nil {
		return m, first.Error
	}
	result := g.Model(m).Updates(data)
	if result.Error != nil {
		return m, result.Error
	}
	return m, nil
}
