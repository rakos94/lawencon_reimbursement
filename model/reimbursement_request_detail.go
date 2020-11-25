package model

type ReimbursementRequestDetail struct {
	BaseModels
	Reciver                string               `gorm:"column:reciver" json:"reciver"`
	Description            string               `gorm:"column:description" json:"description"`
	Attachment             string               `gorm:"column:attachment" json:"attachment"`
	TotalAmount            float64              `gorm:"column:total_amount" json:"total_amount"`
	ReimbursementRequestID string               `gorm:"column:reimbursement_request_id" json:"reimbursement_request_id"`
	ReimbursementRequest   ReimbursementRequest `gorm:"foreignKey:reimbursement_request_id" json:"reimbursement_request"`
}

func (ReimbursementRequestDetail) TableName() string {
	return "reimbursement_request_detail"
}
