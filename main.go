package main

import (
	"lawencon/reimbursement/config"
	"lawencon/reimbursement/controller"
	"lawencon/reimbursement/dao"
	"lawencon/reimbursement/service"
	"log"
	"os"

	"github.com/labstack/echo"

	"gorm.io/gorm"
)

func main() {
	e := echo.New()
	g := NewConn()

	dao.SetDao(g)

	// CLI HANDLE
	if len(os.Args) > 1 {
		handleArgs()
		return
	}

	service.SetService(g)

	jwtGroup := config.SetJwt(e)
	config.ClientConnect()
	controller.SetReimbursementRequest(jwtGroup)
	controller.SetReimbursementCategory(jwtGroup)
	controller.SetReimbursementType(jwtGroup)
	controller.SetReimbursementLimit(jwtGroup)
	controller.SetReimbursementProcess(jwtGroup)
	controller.SetReimbursementRequestDetail(jwtGroup)
	controller.SetReimbursementRequestStatus(jwtGroup)
	controller.SetReimbursementProcessPaid(jwtGroup)
	controller.SetReimbursementApproval(jwtGroup)

	e.Logger.Fatal(e.Start(":" + config.AppPort))
}

// NewConn ...
func NewConn() *gorm.DB {
	g, err := config.Conn()
	if err != nil {
		log.Println(err)
	}

	return g
}

func handleArgs() {
	arg := os.Args[1]

	if arg == "seeder" {
		config.SeederRun()
	} else if arg == "migrate" {
		migrateTable()
	} else if arg == "migrate:seed" {
		migrateTable()
		config.SeederRun()
	}
}

func migrateTable() {
	g := NewConn()
	err := config.MigrateSchema(g)
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("Success Create Table")
	}
}
