package main

import (
	"fmt"
	"lawencon/reimbursement/config"

	"gorm.io/gorm"
)

func main() {
	_ = NewConn()
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
