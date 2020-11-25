package dao

import "lawencon/reimbursement/model"

// ReimbursementTypeDaoImpl ...
type ReimbursementTypeDaoImpl struct{}

// CreateReimbursementType ...
func (ReimbursementTypeDaoImpl) CreateReimbursementType(data *model.ReimbursementType) (*model.ReimbursementType, error) {
	result := g.Create(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}

// GetReimbursementTypeAll ...
func (ReimbursementTypeDaoImpl) GetReimbursementTypeAll() ([]model.ReimbursementType, error) {
	m := []model.ReimbursementType{}
	result := g.Find(&m)
	if result.Error != nil {
		return nil, result.Error
	}
	return m, nil
}

// GetReimbursementTypeByID ...
func (ReimbursementTypeDaoImpl) GetReimbursementTypeByID(id string) (model.ReimbursementType, error) {
	m := model.ReimbursementType{}
	result := g.Preload("Category").Where("id", id).First(&m)
	if result.Error != nil {
		return m, result.Error
	}
	return m, nil
}

// GetReimbursementTypeByCategoryID ...
func (ReimbursementTypeDaoImpl) GetReimbursementTypeByCategoryID(id string) ([]model.ReimbursementType, error) {
	m := []model.ReimbursementType{}
	result := g.Where("category_id", id).Find(&m)
	if result.Error != nil {
		return m, result.Error
	}
	return m, nil
}

// UpdateReimbursementType ...
func (ReimbursementTypeDaoImpl) UpdateReimbursementType(id string, data *model.ReimbursementType) (*model.ReimbursementType, error) {
	m := &model.ReimbursementType{}
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
