package dao

import "lawencon/reimbursement/model"

type ReimbursementTravelDaoImpl struct {}

func (ReimbursementTravelDaoImpl)CreateReimbursementTravel(data *model.ReimbursementTravel)error  {
	result := g.Create(data)
	if result.Error !=nil{
		return result.Error
	}
	return nil
}
func (ReimbursementTravelDaoImpl)GetAllReimbursementTravel()([]model.ReimbursementTravel,error)  {
	var m []model.ReimbursementTravel
	result := g.Find(&m)
	if result.Error !=nil{
		return m,result.Error
	}
	return  m,nil
}
func (ReimbursementTravelDaoImpl)GetByIdReimbursementTravel(id string)(model.ReimbursementTravel,error)  {
	var m model.ReimbursementTravel
	result := g.Where("id",id).First(&m)
	if result.Error !=nil{
		return m,result.Error
	}
	return  m,nil
}
func (ReimbursementTravelDaoImpl)UpdateReimbursementTravel(id string,data * model.ReimbursementTravel)error  {
	var m model.ReimbursementTravel
	result := g.Where("id",id).First(&m)
	if result.Error !=nil{
		return result.Error
	}
	result = g.Model(m).Updates(data)
	if result.Error !=nil{
		return result.Error
	}
	return  nil

}
func (ReimbursementTravelDaoImpl)DeleteReimbursementTravel(id string)error  {
	var m model.ReimbursementTravel
	result := g.Where("id",id).Delete(&m)
	if result.Error !=nil{
		return result.Error
	}
	return  nil
}