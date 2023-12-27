package dbs

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Connect() error {
	var err error
	var dsn = fmt.Sprintf("")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err

}
