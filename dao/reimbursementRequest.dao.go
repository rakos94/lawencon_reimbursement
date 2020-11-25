package dao

import "lawencon/reimbursement/model"

type ReimbursementRequestDao interface {
	CreateReimbursementRequest(data * model.ReimbursementRequest) error
	GetAllReimbursementRequest()([]model.ReimbursementRequest,error)
	GetByIdReimbursementRequest(id string)(model.ReimbursementRequest,error)
	UpdateReimbursementRequest(id string, data * model.ReimbursementRequest)error
	DeleteReimbursementRequest(id string)error

}
