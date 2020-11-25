package controller

import (
	"github.com/labstack/echo"
	"lawencon/reimbursement/model"
	"lawencon/reimbursement/service"
	"net/http"
)
var reimbursementTravelService service.ReimbursementTravelService = service.ReimbursementTravelServiceImpl{}

func SetReimbursementTravel(c *echo.Group)  {
	c.POST("/reimbursement-travel",createReimbursementTravel)
	c.GET("/reimbursement-travel",getAllReimbursementTravel)
	c.GET("/reimbursement-travel/:id",getByIdReimbursementTravel)
	c.PUT("/reimbursement-travel/:id",updateReimbursementTravel)
	c.DELETE("/reimbursement-travel/:id",deleteReimbursementTravel)
}

func createReimbursementTravel(c echo.Context)error  {
	var m = new(model.ReimbursementTravel)
	if err:= c.Bind(m);err!=nil{
		return c.JSON(http.StatusBadRequest,err)
	}
	err:=reimbursementTravelService.CreateReimbursementTravel(m)
	if err!=nil{
		return c.JSON(http.StatusBadRequest,err)
	}
	return c.JSON(http.StatusCreated,m)
}
func getAllReimbursementTravel(c echo.Context)error  {
	result,err := reimbursementTravelService.GetAllReimbursementTravel()
	if err!=nil{
		return c.JSON(http.StatusBadRequest,err)
	}
	return c.JSON(http.StatusCreated,result)
}
func getByIdReimbursementTravel(c echo.Context)error  {
	id := c.Param("id")
	result,err := reimbursementTravelService.GetByIdReimbursementTravel(id)
	if err!=nil{
		return c.JSON(http.StatusBadRequest,err)
	}
	return c.JSON(http.StatusCreated,result)
}
func updateReimbursementTravel(c echo.Context)error  {
	id := c.Param("id")
	m := new(model.ReimbursementTravel)
	if err:= c.Bind(m);err!=nil{
		return c.JSON(http.StatusBadRequest,err)
	}
	err:= reimbursementTravelService.UpdateReimbursementTravel(id,m)
	if err!=nil{
		return c.JSON(http.StatusBadRequest,err)
	}
	return c.JSON(http.StatusCreated,"updated")
}
func deleteReimbursementTravel(c echo.Context)error  {
	id := c.Param("id")
	err:= reimbursementTravelService.DeleteReimbursementTravel(id)
	if err!=nil{
		return c.JSON(http.StatusBadRequest,err)
	}
	return c.JSON(http.StatusCreated,"Deleted")
}