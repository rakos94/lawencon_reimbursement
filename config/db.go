package config

import (
	"fmt"
	"lawencon/reimbursement/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var tables = []interface{}{
	&model.ReimbursementCategory{},
	&model.ReimbursementType{},
	&model.ReimbursementLimit{},
	&model.ReimbursementApproval{},
	&model.ReimbursementProcess{},
	&model.ReimbursementRequest{},
	&model.ReimbursementRequestDetail{},
	&model.ReimbursementRequestStatus{},
	&model.ReimbursementTravel{},
	&model.ReimbursementProcessPaid{},
}

// Conn connect to db
// @return connection
func Conn() (*gorm.DB, error) {
	pg := fmt.Sprintf("host= %v user=%v password=%v dbname=%v port=%v sslmode=%v", host, user, pass, dbname, port, sslmode)
	db, err := gorm.Open(postgres.Open(pg), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}

// MigrateSchema create listed schema based on var tables
func MigrateSchema(db *gorm.DB) error {
	return db.AutoMigrate(tables...)
}
