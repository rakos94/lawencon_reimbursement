package dao

import "lawencon/reimbursement/model"

type ReimbursementRequestDao interface {
	CreateReimbursementRequest(data * model.ReimbursementRequest) error
}
