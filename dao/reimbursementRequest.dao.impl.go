package dao

import "lawencon/reimbursement/model"

type ReimbursementRequestDaoImpl struct {}

func (ReimbursementRequestDaoImpl) CreateReimbursementRequest(data * model.ReimbursementRequest) error {
	result := g.Create(data)
	if result.Error != nil{
		return result.Error
	}
	return nil
}