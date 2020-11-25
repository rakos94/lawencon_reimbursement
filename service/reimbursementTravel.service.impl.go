package service

import (
	"lawencon/reimbursement/dao"
	"lawencon/reimbursement/model"
)

var reimbursementTravelDao dao.ReimbursementTravelDao = dao.ReimbursementTravelDaoImpl{}
type ReimbursementTravelServiceImpl struct {}

func (ReimbursementTravelServiceImpl)CreateReimbursementTravel(data *model.ReimbursementTravel)error  {
	return reimbursementTravelDao.CreateReimbursementTravel(data)
}

func (ReimbursementTravelServiceImpl)GetAllReimbursementTravel()([]model.ReimbursementTravel,error)  {
	return reimbursementTravelDao.GetAllReimbursementTravel()
}
func (ReimbursementTravelServiceImpl)GetByIdReimbursementTravel(id string)(model.ReimbursementTravel,error)  {
	return reimbursementTravelDao.GetByIdReimbursementTravel(id)
}

func (ReimbursementTravelServiceImpl)UpdateReimbursementTravel(id string,data * model.ReimbursementTravel)error  {
	return reimbursementTravelDao.UpdateReimbursementTravel(id,data)
}
func (ReimbursementTravelServiceImpl)DeleteReimbursementTravel(id string)error  {
	return reimbursementTravelDao.DeleteReimbursementTravel(id)
}