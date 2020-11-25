package dao

import "lawencon/reimbursement/model"

// ReimbursementCategoryDaoImpl ...
type ReimbursementCategoryDaoImpl struct{}

// CreateReimbursementCategory ...
func (ReimbursementCategoryDaoImpl) CreateReimbursementCategory(data *model.ReimbursementCategory) (*model.ReimbursementCategory, error) {
	result := g.Create(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}

// GetReimbursementCategoryAll ...
func (ReimbursementCategoryDaoImpl) GetReimbursementCategoryAll() ([]model.ReimbursementCategory, error) {
	m := []model.ReimbursementCategory{}
	result := g.Find(&m)
	if result.Error != nil {
		return nil, result.Error
	}
	return m, nil
}

// GetReimbursementCategoryByID ...
func (ReimbursementCategoryDaoImpl) GetReimbursementCategoryByID(id string) (model.ReimbursementCategory, error) {
	m := model.ReimbursementCategory{}
	result := g.Preload("Type").Where("id", id).First(&m)
	if result.Error != nil {
		return m, result.Error
	}
	return m, nil
}

// UpdateReimbursementCategory ...
func (ReimbursementCategoryDaoImpl) UpdateReimbursementCategory(id string, data *model.ReimbursementCategory) (*model.ReimbursementCategory, error) {
	m := &model.ReimbursementCategory{}
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
