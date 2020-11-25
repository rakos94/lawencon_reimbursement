package controller

import (
	"lawencon/reimbursement/model"
	"lawencon/reimbursement/service"
	"net/http"

	"github.com/labstack/echo"
)

var processService service.ReimbursementProcessService = service.ReimbursementProcessServiceImpl{}

// SetReimbursementProcess ...
func SetReimbursementProcess(c *echo.Group) {
	r := c.Group("/reimbursement/process")
	r.GET("", getProcessAll)
	r.POST("", createProcess)
	r.GET("/:id", getProcessByID)
	r.PUT("/:id", updateProcess)
	r.GET("/request/:requestID", getProcessByEmployeeID)
}

func getProcessAll(c echo.Context) error {
	result, err := processService.GetReimbursementProcessAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func createProcess(c echo.Context) error {
	m := &model.ReimbursementProcess{}
	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := processService.CreateReimbursementProcess(m)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func getProcessByID(c echo.Context) error {
	id := c.Param("id")

	result, err := processService.GetReimbursementProcessByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func updateProcess(c echo.Context) error {
	id := c.Param("id")

	m := &model.ReimbursementProcess{}
	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := processService.UpdateReimbursementProcess(id, m)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func getProcessByEmployeeID(c echo.Context) error {
	requestID := c.Param("requestID")

	result, err := processService.GetReimbursementProcessByRequestID(requestID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
