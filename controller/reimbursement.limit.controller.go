package controller

import (
	"lawencon/reimbursement/model"
	"lawencon/reimbursement/service"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

var limitService service.ReimbursementLimitService = service.ReimbursementLimitServiceImpl{}

// SetReimbursemetLimit ...
func SetReimbursemetLimit(c *echo.Group) {
	r := c.Group("/reimbursement/limit")
	r.GET("", getLimitAll)
	r.POST("", createLimit)
	r.GET("/:id", getLimitByID)
	r.PUT("/:id", updateLimit)
	r.GET("/employee/:employeeID", getLimitByEmployeeID)
}

func getLimitAll(c echo.Context) error {
	result, err := limitService.GetReimbursementLimitAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func createLimit(c echo.Context) error {
	m := &model.ReimbursementLimit{}
	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	log.Println(m)

	result, err := limitService.CreateReimbursementLimit(m)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func getLimitByID(c echo.Context) error {
	id := c.Param("id")

	result, err := limitService.GetReimbursementLimitByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func updateLimit(c echo.Context) error {
	id := c.Param("id")

	m := &model.ReimbursementLimit{}
	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := limitService.UpdateReimbursementLimit(id, m)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func getLimitByEmployeeID(c echo.Context) error {
	employeeID := c.Param("employeeID")

	result, err := limitService.GetReimbursementLimitByEmployeeID(employeeID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
