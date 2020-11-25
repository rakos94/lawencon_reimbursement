package service

import (
	"errors"
	"lawencon/reimbursement/dao"
	"lawencon/reimbursement/model"
	"lawencon/reimbursement/request"
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

// GetReimbursementApprovalByRequestID ...
func (ReimbursementApprovalServiceImpl) GetReimbursementApprovalByRequestID(requestID string) (model.ReimbursementApproval, error) {
	return reimbursementApproval.GetReimbursementApprovalByRequestID(requestID)
}

// ApproveReimbursement ...
func (ReimbursementApprovalServiceImpl) ApproveReimbursement(processID string, data *request.ReimbursementApprovedRequest) (*model.ReimbursementApproval, error) {
	// Start DB TRANSACTION
	tx := g.Begin()

	// Check reimbursement process exist
	process, err := reimbursementProcess.GetReimbursementProcessByID(processID)
	if err != nil {
		return nil, errors.New("Reimbursement Process not exist")
	}

	if data.StatusCode != "approved" {
		return nil, errors.New("Status not approved")
	}

	if data.PaidDate == nil {
		return nil, errors.New("paid_date required")
	}
	if data.PaidType == nil {
		return nil, errors.New("paid_type required")
	}
	if data.ApprovedAmount == nil {
		return nil, errors.New("paid_amount required")
	}

	// Create process paid
	mProcessPaid := &model.ReimbursementProcessPaid{
		ReimbursementProcessID: processID,
		PaidDate:               *data.PaidDate,
		PaidType:               *data.PaidType,
		ApprovedAmount:         *data.ApprovedAmount,
	}
	processPaid, err := reimbursementProcessPaidService.TxCreateReimbursementProcessPaid(mProcessPaid, tx)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("Fail create Process Paid")
	}
	process.ReimbursementProcessPaid = processPaid

	// Create approval with action status
	mApproval := &model.ReimbursementApproval{
		ReimbursementRequestID: process.ReimbursementRequestID,
		Reason:                 data.Reason,
		StatusCode:             data.StatusCode,
	}
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

// CancelReimbursement ...
func (ReimbursementApprovalServiceImpl) CancelReimbursement(processID string, data *request.ReimbursementApprovedRequest) (*model.ReimbursementApproval, error) {
	// Start DB TRANSACTION
	tx := g.Begin()

	// Check reimbursement process exist
	process, err := reimbursementProcess.GetReimbursementProcessByID(processID)
	if err != nil {
		return nil, errors.New("Reimbursement Process not exist")
	}

	if data.StatusCode != "cancel" {
		return nil, errors.New("Status not cancel")
	}

	// Create approval with action status
	mApproval := &model.ReimbursementApproval{
		ReimbursementRequestID: process.ReimbursementRequestID,
		Reason:                 data.Reason,
		StatusCode:             data.StatusCode,
	}
	cancel, err := reimbursementApproval.TxCreateReimbursementApproval(mApproval, tx)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("Fail create approval")
	}

	// TODO
	// UPDATE REIMBURSEMENT STATUS_CODE TO APPROVED

	tx.Commit()
	return cancel, err
}
