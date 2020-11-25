package controller

import (
	"lawencon/reimbursement/model"
	"lawencon/reimbursement/service"
	"net/http"

	"github.com/labstack/echo"
)

var reimbursementService service.ReimbursementRequestService = service.ReimbursementRequestServiceimpl{}

func SetReimbursementRequest(c *echo.Group) {
	c.POST("/reimbursement-request", CreateReimbursement)
	c.GET("/reimbursement-request",GetAllReimbursementRequest)
	c.GET("/reimbursement-request/:id",GetByIdReimbursementRequest)
	c.PUT("/reimbursement-request/:id",UpdateReimbursementRequest)
	c.DELETE("/reimbursement-request/:id",DeleteReimbursementRequest)

}
func CreateReimbursement(c echo.Context) error {
	m := new(model.ReimbursementRequest)
	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err := reimbursementService.CreateReimbursementRequest(m)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, " Create sukses")
}
func GetAllReimbursementRequest(c echo.Context) error {
	result ,err := reimbursementService.GetAllReimbursementRequest()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
func GetByIdReimbursementRequest(c echo.Context)error  {
	id := c.Param("id")
	result, err := reimbursementService.GetByIdReimbursementRequest(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
func UpdateReimbursementRequest(c echo.Context)error  {
	id := c.Param("id")
	m := new(model.ReimbursementRequest)
	if err := c.Bind(m);err!=nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err := reimbursementService.UpdateReimbursementRequest(id,m)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "Updated")
}
func DeleteReimbursementRequest(c echo.Context)error  {
	id := c.Param("id")
	err := reimbursementService.DeleteReimbursementRequest(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "deleted")
}