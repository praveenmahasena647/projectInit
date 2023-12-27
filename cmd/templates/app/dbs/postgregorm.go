package dbs

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Connect() error {
	var err error
	var connectionUri = fmt.Sprintf("")
	DB, err = gorm.Open(postgres.Open(connectionUri), &gorm.Config{})
	return err
}
