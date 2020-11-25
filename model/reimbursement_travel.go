package model

import "lawencon/reimbursement/datatype"

type ReimbursementTravel struct {
	BaseModels
	DepartDate             datatype.Date        `gorm:"column:depart_date" json:"depart_date"`
	ReturnDate             datatype.Date        `gorm:"column:return_date" json:"return_date"`
	From                   string               `gorm:"column:from" json:"from"`
	To                     string               `gorm:"column:to" json:"to"`
	ReimbursementRequestID string               `gorm:"column:reimbursement_request_id" json:"reimbursement_request_id"`
	ReimbursementRequest   ReimbursementRequest `gorm:"foreignKey:reimbursement_request_id" json:"reimbursement_request"`
}

func (ReimbursementTravel) TableName() string {
	return "reimbursement_travel"
}
