package service

import "lawencon/reimbursement/model"

type ReimbursementTravelService interface {
	CreateReimbursementTravel(data *model.ReimbursementTravel)error
	GetAllReimbursementTravel()([]model.ReimbursementTravel,error)
	GetByIdReimbursementTravel(id string)(model.ReimbursementTravel,error)
	UpdateReimbursementTravel(id string,data * model.ReimbursementTravel)error
	DeleteReimbursementTravel(id string)error
}
