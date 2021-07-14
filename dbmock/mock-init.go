package dbmock

import (
	"yana-golang/db_interface"

	_ "github.com/go-sql-driver/mysql"
)

func Init() *db_interface.MYSQL {

	return &db_interface.MYSQL{
		User: InitUserDB(),
	}
}
