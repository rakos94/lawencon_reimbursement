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