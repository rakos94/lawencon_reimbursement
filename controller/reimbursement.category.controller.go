package controller

import (
	"lawencon/reimbursement/model"
	"lawencon/reimbursement/service"
	"net/http"

	"github.com/labstack/echo"
)

var categoryService service.ReimbursementCategoryService = service.ReimbursementCategoryServiceImpl{}

// SetReimbursemetCategory ...
func SetReimbursemetCategory(c *echo.Group) {
	group := c.Group("/reimbursement/category")
	group.GET("", getCategoryAll)
	group.POST("", createCategory)
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
