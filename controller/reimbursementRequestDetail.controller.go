package controller

import (
	"github.com/labstack/echo"
	"lawencon/reimbursement/model"
	"lawencon/reimbursement/service"
	"net/http"
)

var reimbursementRequestDetailService service.ReimbursementRequestDetailService =service.ReimbursementRequestDetailServiceImpl{}

func SetReimbursementRequestDetail(c *echo.Group)  {
	c.POST("/reimbursement-request-detail",CreateReimbursementRequestDetail)
	c.GET("/reimbursement-request-detail",GetAllReimbursementRequestDetai)
	c.GET("/reimbursement-request-detail/:id",GetByIdReimbursementRequestDetail)
	c.PUT("/reimbursement-request-detail/:id",updateReimbursementRequestDetail)
	c.DELETE("/reimbursement-request-detail/:id",deleteReimbursementRequestDetai)
}

func CreateReimbursementRequestDetail(c echo.Context)error  {
	m := new(model.ReimbursementRequestDetail)
	err := c.Bind(&m)
	if err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = reimbursementRequestDetailService.CreateReimbursementRequestDetail(m)
	if err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated,"Created")
}

func GetAllReimbursementRequestDetai(c echo.Context)error  {

	result,err := reimbursementRequestDetailService.GetAllReimbursementRequestDetai()
	if err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK,result)
}
func GetByIdReimbursementRequestDetail(c echo.Context)error  {
	id:= c.Param("id")
	result,err := reimbursementRequestDetailService.GetByIdReimbursementRequestDetail(id)
	if err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK,result)
}
func updateReimbursementRequestDetail(c echo.Context)error  {
	id:= c.Param("id")
	m:= new(model.ReimbursementRequestDetail)
	err := c.Bind(&m)
	if err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err =  reimbursementRequestDetailService.UpdateReimbursementRequestDetail(id,m)
	return c.JSON(http.StatusCreated,"Updated")
}

func deleteReimbursementRequestDetai(c echo.Context)error  {
	id:= c.Param("id")
	err:= reimbursementRequestDetailService.DeleteReimbursementRequestDetail(id)
	if err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated,"Updated")
}