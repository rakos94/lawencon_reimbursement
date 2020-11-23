package main

import (
	"fmt"
	"github.com/labstack/echo"
	"lawencon/reimbursement/config"
	"lawencon/reimbursement/controller"
	"lawencon/reimbursement/dao"
	"lawencon/reimbursement/service"

	"gorm.io/gorm"
)

func main() {
	e:= echo.New()
	g := NewConn()
	//config.Conn();
	dao.SetDao(g)
	service.SetService(g)

	jwtGroup := config.SetJwt(e)
	config.ClientConnect()
	controller.SetReimbursemetRequest(jwtGroup)
	e.Logger.Fatal(e.Start(":" + config.AppPort))
}

// NewConn ...
func NewConn() *gorm.DB {
	g, err := config.Conn()


	if err == nil {
		config.MigrateSchema(g)
		fmt.Println("Success Create Table")
	}

	return g
}
