package controller

import (
	"lawencon/reimbursement/model"
	"lawencon/reimbursement/service"
	"net/http"

	"github.com/labstack/echo"
)

var categoryService service.ReimbursementCategoryService = service.ReimbursementCategoryServiceImpl{}

// SetReimbursementCategory ...
func SetReimbursementCategory(c *echo.Group) {
	r := c.Group("/reimbursement/category")
	r.GET("", getCategoryAll)
	r.POST("", createCategory)
	r.GET("/:id", getCategoryByID)
	r.PUT("/:id", updateCategory)
}

func getCategoryAll(c echo.Context) error {
	result, err := categoryService.GetReimbursementCategoryAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func createCategory(c echo.Context) error {
	m := &model.ReimbursementCategory{}
	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := categoryService.CreateReimbursementCategory(m)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func getCategoryByID(c echo.Context) error {
	id := c.Param("id")

	result, err := categoryService.GetReimbursementCategoryByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func updateCategory(c echo.Context) error {
	id := c.Param("id")

	m := &model.ReimbursementCategory{}
	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := categoryService.UpdateReimbursementCategory(id, m)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
