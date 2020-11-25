package dao

import "lawencon/reimbursement/model"

type ReimbursementRequestDetailDaoImpl struct {}

func (ReimbursementRequestDetailDaoImpl)CreateReimbursementRequestDetail(data* model.ReimbursementRequestDetail)error  {
	result:= g.Create(data)
	if result.Error != nil{
		return result.Error
	}
	return nil
}
func (ReimbursementRequestDetailDaoImpl)GetAllReimbursementRequestDetai()([]model.ReimbursementRequestDetail,error)  {
	m :=[]model.ReimbursementRequestDetail{}
	result:=g.Find(&m)
	if result.Error != nil{
		return m, result.Error
	}
	return m, nil
}

func (ReimbursementRequestDetailDaoImpl)GetByIdReimbursementRequestDetail(id string)(model.ReimbursementRequestDetail,error)  {
	m:= model.ReimbursementRequestDetail{}
	result := g.Where("id",id).Find(&m)
	if result.Error != nil{
		return m, result.Error
	}
	return m, nil
}

func (ReimbursementRequestDetailDaoImpl)UpdateReimbursementRequestDetail(id string,data * model.ReimbursementRequestDetail)error  {
	m:= model.ReimbursementRequestDetail{}
	result := g.Where("id",id).Find(&m)
	if result.Error != nil{
		return result.Error
	}
	result = g.Model(&m).Updates(data)
	if result.Error != nil{
		return result.Error
	}
	return nil
}

func (ReimbursementRequestDetailDaoImpl)DeleteReimbursementRequestDetail(id string)error  {
	m:= model.ReimbursementRequestDetail{}
	result := g.Where("id",id).Delete(&m)
	if result.Error != nil{
		return result.Error
	}
	return nil
}