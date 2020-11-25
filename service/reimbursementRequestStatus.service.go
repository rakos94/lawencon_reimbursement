package service

import "lawencon/reimbursement/model"

type ReimbursementRequestStatusService interface {
	CreateReimbursementRequestStatus(data * model.ReimbursementRequestStatus)error
	GetAllReimbursementRequestStatus()([]model.ReimbursementRequestStatus,error)
	GetByIdReimbursementRequestStatus(id string)(model.ReimbursementRequestStatus,error)
	UpdateReimbursementRequestStatus(id string,data * model.ReimbursementRequestStatus)error
	DeleteReimbursementRequestStatus(id string)error
}
