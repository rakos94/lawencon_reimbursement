package service

import (
	"lawencon/reimbursement/dao"
	"lawencon/reimbursement/model"
)

var reimbursementRequestDetaildao dao.ReimbursementRequestDetailDao = dao.ReimbursementRequestDetailDaoImpl{}
type ReimbursementRequestDetailServiceImpl struct {}

func (ReimbursementRequestDetailServiceImpl)CreateReimbursementRequestDetail(data* model.ReimbursementRequestDetail)error  {
	return reimbursementRequestDetaildao.CreateReimbursementRequestDetail(data)
}
func (ReimbursementRequestDetailServiceImpl)GetAllReimbursementRequestDetai()([]model.ReimbursementRequestDetail,error)  {
	return reimbursementRequestDetaildao.GetAllReimbursementRequestDetai()
}
func (ReimbursementRequestDetailServiceImpl)GetByIdReimbursementRequestDetail(id string)(model.ReimbursementRequestDetail,error) {
	return reimbursementRequestDetaildao.GetByIdReimbursementRequestDetail(id)
}
func (ReimbursementRequestDetailServiceImpl)UpdateReimbursementRequestDetail(id string,data * model.ReimbursementRequestDetail)error{
	return reimbursementRequestDetaildao.UpdateReimbursementRequestDetail(id,data)
}
func (ReimbursementRequestDetailServiceImpl)DeleteReimbursementRequestDetail(id string)error{
	return reimbursementRequestDetaildao.DeleteReimbursementRequestDetail(id)
}