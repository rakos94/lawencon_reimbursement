package dao

import (
	"errors"
	"lawencon/reimbursement/model"
)

// ReimbursementLimitDaoImpl ...
type ReimbursementLimitDaoImpl struct{}

// CreateReimbursementLimit ...
func (ReimbursementLimitDaoImpl) CreateReimbursementLimit(data *model.ReimbursementLimit) (*model.ReimbursementLimit, error) {
	err := checkExistEmployeeIDAndTypeCode(data)
	if err != nil {
		return nil, err
	}

	result := g.Create(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}

// GetReimbursementLimitAll ...
func (ReimbursementLimitDaoImpl) GetReimbursementLimitAll() ([]model.ReimbursementLimit, error) {
	m := []model.ReimbursementLimit{}
	result := g.Find(&m)
	if result.Error != nil {
		return nil, result.Error
	}
	return m, nil
}

// GetReimbursementLimitByID ...
func (ReimbursementLimitDaoImpl) GetReimbursementLimitByID(id string) (model.ReimbursementLimit, error) {
	m := model.ReimbursementLimit{}
	result := g.Preload("Type").Where("id", id).First(&m)
	if result.Error != nil {
		return m, result.Error
	}
	return m, nil
}

// GetReimbursementLimitByEmployeeID ...
func (ReimbursementLimitDaoImpl) GetReimbursementLimitByEmployeeID(id string) ([]model.ReimbursementLimit, error) {
	m := []model.ReimbursementLimit{}
	result := g.Where("employee_id", id).Find(&m)
	if result.Error != nil {
		return m, result.Error
	}
	return m, nil
}

// UpdateReimbursementLimit ...
func (ReimbursementLimitDaoImpl) UpdateReimbursementLimit(id string, data *model.ReimbursementLimit) (*model.ReimbursementLimit, error) {
	m := &model.ReimbursementLimit{}
	first := g.Where("id", id).First(&m)
	if first.Error != nil {
		return m, first.Error
	}

	// check employee already have new ReimbursementTypeCode or not
	if data.ReimbursementTypeCode != m.ReimbursementTypeCode {
		data.EmployeeID = m.EmployeeID
		err := checkExistEmployeeIDAndTypeCode(data)
		if err != nil {
			return nil, err
		}
	}

	result := g.Model(m).Updates(data)
	if result.Error != nil {
		return m, result.Error
	}
	return m, nil
}

func checkExistEmployeeIDAndTypeCode(data *model.ReimbursementLimit) error {
	var count int64
	m := &model.ReimbursementLimit{}
	result := g.Model(m).Where("employee_id", data.EmployeeID).
		Where("reimbursement_type_code", data.ReimbursementTypeCode).
		Count(&count)
	if result.Error != nil {
		return result.Error
	}
	if count > 0 {
		return errors.New("Employee id & Reimbursement type code already exist")
	}
	return nil
}
