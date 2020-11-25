package controller

import (
	"lawencon/reimbursement/service"
	"net/http"

	"github.com/labstack/echo"
)

var approvalService service.ReimbursementApprovalService = service.ReimbursementApprovalServiceImpl{}

// SetReimbursementApproval ...
func SetReimbursementApproval(c *echo.Group) {
	r := c.Group("/reimbursement/approval")
	r.GET("", getApprovalAll)
	r.GET("/:id", getProcessByID)
	r.GET("/request/:requestID", getApprovalByRequestID)
}

func getApprovalAll(c echo.Context) error {
	result, err := approvalService.GetReimbursementApprovalAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func getApprovalByID(c echo.Context) error {
	id := c.Param("id")

	result, err := approvalService.GetReimbursementApprovalByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func getApprovalByRequestID(c echo.Context) error {
	requestID := c.Param("requestID")

	result, err := approvalService.GetReimbursementApprovalByRequestID(requestID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
