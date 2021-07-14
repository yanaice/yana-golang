package dbmongo

import "yana-golang/db_interface"

func Init() *db_interface.MYSQL {

	return &db_interface.MYSQL{
		User: InitUserDB(),
	}
}
