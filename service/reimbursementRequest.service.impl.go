package service

import (
	"lawencon/reimbursement/dao"
	"lawencon/reimbursement/model"
)

var reimbursementDao dao.ReimbursementRequestDao =dao.ReimbursementRequestDaoImpl{}
type ReimbursementRequestServiceimpl struct {}

func (ReimbursementRequestServiceimpl)CreateReimbursementRequest(data * model.ReimbursementRequest) error  {
	return  reimbursementDao.CreateReimbursementRequest(data)
}
func (ReimbursementRequestServiceimpl)GetAllReimbursementRequest()([]model.ReimbursementRequest,error)  {
	return reimbursementDao.GetAllReimbursementRequest();
}

func (ReimbursementRequestServiceimpl)GetByIdReimbursementRequest(id string)(model.ReimbursementRequest,error)  {
	return reimbursementDao.GetByIdReimbursementRequest(id)
}
func (ReimbursementRequestServiceimpl)UpdateReimbursementRequest(id string, data * model.ReimbursementRequest)error  {
	return reimbursementDao.UpdateReimbursementRequest(id,data)
}

func (ReimbursementRequestServiceimpl)DeleteReimbursementRequest(id string)error  {
	return reimbursementDao.DeleteReimbursementRequest(id)
}