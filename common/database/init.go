package database

import (
	"fmt"
)

func Setup(Driver string) error {

	if Driver == "mysql" {
		db := new(Mysql)
		db.Setup()
		return nil
	} else if Driver == "pgsql" {
		db := new(Pgsql)
		db.Setup()
		return nil
	} else {
		return fmt.Errorf("not support database type for : [%s]", Driver)
	}

}
