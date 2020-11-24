package service

import (
	"errors"
	"lawencon/reimbursement/dao"
	"lawencon/reimbursement/datatype"
	"lawencon/reimbursement/model"
	"lawencon/reimbursement/request"
	"time"
)

var reimbursementApproval dao.ReimbursementApprovalDao = dao.ReimbursementApprovalDaoImpl{}
var reimbursementRequestService ReimbursementRequestService = ReimbursementRequestServiceimpl{}
var reimbursementProcessService ReimbursementProcessService = ReimbursementProcessServiceImpl{}
var reimbursementProcessPaidService ReimbursementProcessPaidService = ReimbursementProcessPaidServiceImpl{}

// ReimbursementApprovalServiceImpl ...
type ReimbursementApprovalServiceImpl struct{}

// CreateReimbursementApproval ...
func (ReimbursementApprovalServiceImpl) CreateReimbursementApproval(data *model.ReimbursementApproval) (*model.ReimbursementApproval, error) {
	return reimbursementApproval.CreateReimbursementApproval(data)
}

// GetReimbursementApprovalAll ...
func (ReimbursementApprovalServiceImpl) GetReimbursementApprovalAll() ([]model.ReimbursementApproval, error) {
	return reimbursementApproval.GetReimbursementApprovalAll()
}

// GetReimbursementApprovalByID ...
func (ReimbursementApprovalServiceImpl) GetReimbursementApprovalByID(id string) (model.ReimbursementApproval, error) {
	return reimbursementApproval.GetReimbursementApprovalByID(id)
}

// GetReimbursementApprovalByReimbursementID ...
func (ReimbursementApprovalServiceImpl) GetReimbursementApprovalByReimbursementID(id string) ([]model.ReimbursementApproval, error) {
	return reimbursementApproval.GetReimbursementApprovalByReimbursementID(id)
}

// ApproveReimbursement ...
func (ReimbursementProcessServiceImpl) ApproveReimbursement(reimbursementID string, data request.ReimbursementApprovedRequest) (*model.ReimbursementApproval, error) {
	// Start DB TRANSACTION
	tx := g.Begin()

	// TODO
	// CHECK REIMBURSEMENT ID EXIST OR NOT
	// _, err := re.(id)
	// if err != nil {
	// 	return nil, errors.New("Reimbursement Process not exist")
	// }

	// Check reimbursement process exist
	process, err := reimbursementProcess.GetReimbursementProcessByReimbursementID(reimbursementID)
	if err != nil {
		return nil, errors.New("Reimbursement Process not exist")
	}

	// Create process paid
	mProcessPaid := &model.ReimbursementProcessPaid{
		ReimbursementProcessID: reimbursementID,
		PaidDate:               datatype.Date(time.Now()),
		PaidType:               data.PaidType,
	}
	processPaid, err := reimbursementProcessPaidService.TxCreateReimbursementProcessPaid(mProcessPaid, tx)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("Fail create Process Paid")
	}
	process.ReimbursementProcessPaid = *processPaid

	// Create approval with action status
	mApproval := &model.ReimbursementApproval{ReimbursementID: data.ReimbursementID, Reason: data.Reason, StatusCode: "approved"}
	approve, err := reimbursementApproval.TxCreateReimbursementApproval(mApproval, tx)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("Fail create approval")
	}

	// TODO
	// UPDATE REIMBURSEMENT STATUS_CODE TO APPROVED

	tx.Commit()
	return approve, err
}
