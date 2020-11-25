package controller

import (
	"github.com/labstack/echo"
	"lawencon/reimbursement/model"
	"lawencon/reimbursement/service"
	"net/http"
)

var reimbursementRequestStatus service.ReimbursementRequestStatusService = service.ReimbursementRequestStatusServiceImpl{}

func SetReimbursementRequestStatus(c *echo.Group)  {
	c.POST("/reimbursement-request-status",createReimbursementRequestStatus)
	c.GET("/reimbursement-request-status",getAllReimbursementRequestStatus)
	c.GET("/reimbursement-request-status/:id",getByIdReimbursementRequestStatus)
	c.PUT("/reimbursement-request-status/:id",updateReimbursementRequestStatus)
	c.DELETE("/reimbursement-request-status/:id",DeleteReimbursementRequestStatus)
}

func createReimbursementRequestStatus(c echo.Context)error  {
	m:= new(model.ReimbursementRequestStatus)
	if err := c.Bind(m); err!=nil{
		return c.JSON(http.StatusBadRequest,err)
	}
	err:= reimbursementRequestStatus.CreateReimbursementRequestStatus(m)
	if err!=nil{
		return c.JSON(http.StatusBadRequest,err)
	}
	return c.JSON(http.StatusOK,"Created")
}
func getAllReimbursementRequestStatus(c echo.Context)error  {
	result,err := reimbursementRequestStatus.GetAllReimbursementRequestStatus()
	if err!=nil{
		return c.JSON(http.StatusBadRequest,err)
	}
	return c.JSON(http.StatusOK,result)
}

func getByIdReimbursementRequestStatus(c echo.Context)error  {
	id:= c.Param("id")
	result,err := reimbursementRequestStatus.GetByIdReimbursementRequestStatus(id)
	if err!=nil{
		return c.JSON(http.StatusBadRequest,err)
	}
	return c.JSON(http.StatusOK,result)
}
func updateReimbursementRequestStatus(c echo.Context)error  {
	id := c.Param("id")
	m := new(model.ReimbursementRequestStatus)
	if err := c.Bind(m); err!=nil{
		return c.JSON(http.StatusBadRequest,err)
	}
	err:= reimbursementRequestStatus.UpdateReimbursementRequestStatus(id,m)
	if err!=nil{
		return c.JSON(http.StatusBadRequest,err)
	}
	return c.JSON(http.StatusOK,"Updated")
}
func DeleteReimbursementRequestStatus(c echo.Context)error  {
	id:= c.Param("id")
	err := reimbursementRequestStatus.DeleteReimbursementRequestStatus(id)
	if err!=nil{
		return c.JSON(http.StatusBadRequest,err)
	}
	return c.JSON(http.StatusOK,"deleted")
}