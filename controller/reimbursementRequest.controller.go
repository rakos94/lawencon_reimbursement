package controller

import (
	"github.com/labstack/echo"
	"lawencon/reimbursement/model"
	"lawencon/reimbursement/service"
	"net/http"
)

var reimbursementService service.ReimbursementRequestService =service.ReimbursementRequestServiceimpl{}

func SetReimbursemetRequest(c *echo.Group)  {
	c.POST("/reimbursement-request",CreateReimbursement)
}
func CreateReimbursement(c echo.Context)error  {
	m := new(model.ReimbursementRequest)
	if err:= c.Bind(m); err!=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	err := reimbursementService.CreateReimbursementRequest(m)
	if err != nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	return  c.JSON(http.StatusOK," Create sukses")
}