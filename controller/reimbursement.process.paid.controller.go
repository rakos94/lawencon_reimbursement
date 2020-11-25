package controller

import (
	"lawencon/reimbursement/model"
	"lawencon/reimbursement/service"
	"net/http"

	"github.com/labstack/echo"
)

var processPaidService service.ReimbursementProcessPaidService = service.ReimbursementProcessPaidServiceImpl{}

// SetReimbursementProcessPaid ...
func SetReimbursementProcessPaid(c *echo.Group) {
	r := c.Group("/reimbursement/process/paid")
	r.GET("", getProcessPaidAll)
	r.POST("", createProcessPaid)
	r.GET("/:id", getProcessPaidByID)
	c.GET("/reimbursement/process/:processID/paid", getProcessPaidByProcessID)
}

func getProcessPaidAll(c echo.Context) error {
	result, err := processPaidService.GetReimbursementProcessPaidAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func createProcessPaid(c echo.Context) error {
	m := &model.ReimbursementProcessPaid{}
	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := processPaidService.CreateReimbursementProcessPaid(m)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func getProcessPaidByID(c echo.Context) error {
	id := c.Param("id")

	result, err := processPaidService.GetReimbursementProcessPaidByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func getProcessPaidByProcessID(c echo.Context) error {
	processID := c.Param("processID")

	result, err := processPaidService.GetReimbursementProcessPaidByReimbursementProcessID(processID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
