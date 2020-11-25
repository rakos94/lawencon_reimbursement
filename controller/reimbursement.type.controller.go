package controller

import (
	"lawencon/reimbursement/model"
	"lawencon/reimbursement/service"
	"net/http"

	"github.com/labstack/echo"
)

var typeService service.ReimbursementTypeService = service.ReimbursementTypeServiceImpl{
	CategoryService: service.ReimbursementCategoryServiceImpl{},
}

// SetReimbursementType ...
func SetReimbursementType(c *echo.Group) {
	r := c.Group("/reimbursement/type")
	r.GET("", getTypeAll)
	r.POST("", createType)
	r.GET("/:id", getTypeByID)
	r.PUT("/:id", updateType)
	c.GET("/reimbursement/category/:code/type", getTypeByCategoryCode)
}

func getTypeAll(c echo.Context) error {
	result, err := typeService.GetReimbursementTypeAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func createType(c echo.Context) error {
	m := &model.ReimbursementType{}
	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := typeService.CreateReimbursementType(m)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func getTypeByID(c echo.Context) error {
	id := c.Param("id")

	result, err := typeService.GetReimbursementTypeByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func updateType(c echo.Context) error {
	id := c.Param("id")

	m := &model.ReimbursementType{}
	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := typeService.UpdateReimbursementType(id, m)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func getTypeByCategoryCode(c echo.Context) error {
	code := c.Param("code")

	result, err := typeService.GetReimbursementTypeByCategoryCode(code)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
