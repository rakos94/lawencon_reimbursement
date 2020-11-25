package dao

import "lawencon/reimbursement/model"

type ReimbursementRequestStatusDaoImpl struct {}

func (ReimbursementRequestStatusDaoImpl)CreateReimbursementRequestStatus(data * model.ReimbursementRequestStatus)error  {
	result := g.Create(data)
	if result.Error != nil{
		return  result.Error
	}
	return nil
}
func (ReimbursementRequestStatusDaoImpl)GetAllReimbursementRequestStatus()([]model.ReimbursementRequestStatus,error)  {
	var m []model.ReimbursementRequestStatus
	result := g.Find(&m)
	if result.Error != nil{
		return m, result.Error
	}
	return m, nil
}
func (ReimbursementRequestStatusDaoImpl)GetByIdReimbursementRequestStatus(id string)(model.ReimbursementRequestStatus,error)  {
	var m model.ReimbursementRequestStatus
	result := g.Where("id",id).First(&m)
	if result.Error != nil{
		return m, result.Error
	}
	return m, nil
}
func (ReimbursementRequestStatusDaoImpl)UpdateReimbursementRequestStatus(id string,data * model.ReimbursementRequestStatus)error  {
	var m model.ReimbursementRequestStatus
	result := g.Where("id",id).First(&m)
	if result.Error != nil{
		return result.Error
	}
	result = g.Model(&m).Updates(data)
	if result.Error != nil{
		return result.Error
	}
	return nil
}

func (ReimbursementRequestStatusDaoImpl)DeleteReimbursementRequestStatus(id string)error  {
	var m model.ReimbursementRequestStatus
	result := g.Where("id",id).Delete(&m)
	if result.Error != nil{
		return result.Error
	}
	return nil
}