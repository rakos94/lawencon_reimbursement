package service

import "lawencon/reimbursement/model"

type ReimbursementRequestDetailService interface {
	CreateReimbursementRequestDetail(data* model.ReimbursementRequestDetail)error
	GetAllReimbursementRequestDetai()([]model.ReimbursementRequestDetail,error)
	GetByIdReimbursementRequestDetail(id string)(model.ReimbursementRequestDetail,error)
	UpdateReimbursementRequestDetail(id string,data * model.ReimbursementRequestDetail)error
	DeleteReimbursementRequestDetail(id string)error
}
