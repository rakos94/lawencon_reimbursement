package service

import (
	"lawencon/reimbursement/dao"
	"lawencon/reimbursement/model"
)

var reimbursementRequestStatusDao dao.ReimbursementRequestStatusDao = dao.ReimbursementRequestStatusDaoImpl{}
type ReimbursementRequestStatusServiceImpl struct {}

func (ReimbursementRequestStatusServiceImpl)CreateReimbursementRequestStatus(data * model.ReimbursementRequestStatus)error  {
	return reimbursementRequestStatusDao.CreateReimbursementRequestStatus(data)
}
func (ReimbursementRequestStatusServiceImpl)GetAllReimbursementRequestStatus()([]model.ReimbursementRequestStatus,error)  {
	return reimbursementRequestStatusDao.GetAllReimbursementRequestStatus()
}
func (ReimbursementRequestStatusServiceImpl)GetByIdReimbursementRequestStatus(id string)(model.ReimbursementRequestStatus,error)  {
	return reimbursementRequestStatusDao.GetByIdReimbursementRequestStatus(id)
}
func (ReimbursementRequestStatusServiceImpl)UpdateReimbursementRequestStatus(id string,data * model.ReimbursementRequestStatus)error  {
	return reimbursementRequestStatusDao.UpdateReimbursementRequestStatus(id,data)
}
func (ReimbursementRequestStatusServiceImpl)DeleteReimbursementRequestStatus(id string)error  {
	return reimbursementRequestStatusDao.DeleteReimbursementRequestStatus(id)
}