package model

type ReimbursementRequest struct {
	BaseModels
	EmployeeID            string `gorm:"column:employee_id;not null" json:"employee_id"`
	ReimbursementTypeCode string `gorm:"column:reimbursement_code;not null" json:"reimbursement_type_code"`
	TrxCode               string `gorm:"column:trx_code;not null" json:"trx_code"`
	TrxDate               string `gorm:"column:trx_date;not null" json:"trx_date"`
	StatusCode            string `gorm:"column:status_code:not null" json:"status_code"`
	IsPaid                bool   `gorm:"column:is_paid:not null" json:"is_paid"`
}

func (ReimbursementRequest) TableName() string {
	return "reimbursement_request"
}
