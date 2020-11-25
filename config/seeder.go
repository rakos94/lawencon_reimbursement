package config

import (
	"lawencon/reimbursement/dao"
	"lawencon/reimbursement/model"
	"log"

	"gorm.io/gorm"
)

var g *gorm.DB
var typeDao dao.ReimbursementTypeDao = dao.ReimbursementTypeDaoImpl{}
var categoryDao dao.ReimbursementCategoryDao = dao.ReimbursementCategoryDaoImpl{}
var statusDao dao.ReimbursementRequestStatusDao = dao.ReimbursementRequestStatusDaoImpl{}

// SeederRun ...
func SeederRun() {
	setConnection()
	categorySeeder()
	typeSeeder()
	statusSeeder()
}

func setConnection() {
	conn, err := Conn()
	if err != nil {
		log.Println(err.Error())
	}
	g = conn
}

func categorySeeder() {
	truncateTable(model.ReimbursementCategory{}.TableName())
	data := []model.ReimbursementCategory{
		{Code: "MDC", Name: "Medical"},
		{Code: "TRV", Name: "Travel"},
	}

	for _, d := range data {
		_, err := categoryDao.CreateReimbursementCategory(&d)
		if err != nil {
			log.Println(err.Error())
		}
	}

	log.Println(model.ReimbursementCategory{}.TableName(), "seeder created")
}

func typeSeeder() {
	truncateTable(model.ReimbursementType{}.TableName())
	data := []model.ReimbursementType{
		{CategoryCode: "MDC", Code: "MDC-SGR", Name: "Surgery"},
		{CategoryCode: "MDC", Code: "MDC-DTL", Name: "Dentals"},
		{CategoryCode: "TRV", Code: "TRV-TRN", Name: "Train"},
		{CategoryCode: "TRV", Code: "TRV-BUS", Name: "Bus"},
	}

	for _, d := range data {
		_, err := typeDao.CreateReimbursementType(&d)
		if err != nil {
			log.Println(err.Error())
		}
	}
	log.Println(model.ReimbursementType{}.TableName(), "seeder created")
}

func statusSeeder() {
	truncateTable(model.ReimbursementRequestStatus{}.TableName())
	data := []model.ReimbursementRequestStatus{
		{StatusCode: "PEND", Name: "Pending"},
		{StatusCode: "APRV", Name: "Approve"},
		{StatusCode: "CANC", Name: "Cancel"},
		{StatusCode: "PROC", Name: "Process"},
	}

	for _, d := range data {
		err := statusDao.CreateReimbursementRequestStatus(&d)
		if err != nil {
			log.Println(err.Error())
		}
	}
	log.Println(model.ReimbursementRequestStatus{}.TableName(), "seeder created")
}

func truncateTable(tableName string) {
	truncateQuery := "DELETE FROM " + tableName
	if rs := g.Exec(truncateQuery); rs.Error != nil {
		log.Println(rs.Error)
		return
	}
}
