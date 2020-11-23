package service

import "lawencon/reimbursement/model"

type ReimbursementRequestService interface {
	CreateReimbursementRequest(data * model.ReimbursementRequest) error
}
