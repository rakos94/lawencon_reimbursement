package dao

import "lawencon/reimbursement/model"

// ReimbursementLimitDaoImpl ...
type ReimbursementLimitDaoImpl struct{}

// CreateReimbursementLimit ...
func (ReimbursementLimitDaoImpl) CreateReimbursementLimit(data *model.ReimbursementLimit) (*model.ReimbursementLimit, error) {
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
	result := g.Where("id", id).First(&m)
	if result.Error != nil {
		return m, result.Error
	}
	return m, nil
}

// GetReimbursementLimitByPersonID ...
func (ReimbursementLimitDaoImpl) GetReimbursementLimitByPersonID(id string) ([]model.ReimbursementLimit, error) {
	m := []model.ReimbursementLimit{}
	result := g.Where("person_id", id).Find(&m)
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
	result := g.Model(m).Updates(data)
	if result.Error != nil {
		return m, result.Error
	}
	return m, nil
}
