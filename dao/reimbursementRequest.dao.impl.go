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
func (ReimbursementRequestDaoImpl)GetAllReimbursementRequest()([]model.ReimbursementRequest,error)  {
	m := []model.ReimbursementRequest{}
	result := g.Find(&m)
	if result.Error != nil{
		return m,result.Error
	}
	return m,nil
}
func (ReimbursementRequestDaoImpl)GetByIdReimbursementRequest(id string)(model.ReimbursementRequest,error)  {
	m := model.ReimbursementRequest{}
	result := g.Where("id",id).First(&m)
	if result.Error != nil{
		return m,result.Error
	}
	return m,nil
}
func (ReimbursementRequestDaoImpl)UpdateReimbursementRequest(id string, data * model.ReimbursementRequest)error  {
	m:= model.ReimbursementRequest{}
	result:= g.Where("id",id).Find(&m)
	if result.Error != nil{
		return result.Error
	}
	result = g.Model(&m).Updates(data)
	if result.Error != nil{
		return result.Error
	}
	return nil
}
func (ReimbursementRequestDaoImpl)DeleteReimbursementRequest(id string)error  {
	m:= model.ReimbursementRequest{}
	result := g.Where("id",id).Delete(&m)
	if result.Error != nil{
		return result.Error
	}
	return nil
}
