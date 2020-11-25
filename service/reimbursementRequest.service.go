package service

import "lawencon/reimbursement/model"

type ReimbursementRequestService interface {
	CreateReimbursementRequest(data * model.ReimbursementRequest) error
	GetAllReimbursementRequest()([]model.ReimbursementRequest,error)
	GetByIdReimbursementRequest(id string)(model.ReimbursementRequest,error)
	UpdateReimbursementRequest(id string, data * model.ReimbursementRequest)error
	DeleteReimbursementRequest(id string)error
}
