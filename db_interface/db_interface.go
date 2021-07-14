package db_interface

import "yana-golang/model"

type MYSQL struct {
	User UserDatabaseInterface
}

type UserDatabaseInterface interface {
	CreateUser(req model.User) (model.User, error)
	GetUser(userId string) (model.User, error)
	GetUserByCitizenId(citizenId string) (model.User, error)
	GetUserByLineId(lineId string) (model.User, error)
	UpdateUserByLineId(lineId string, req model.User) (model.User, error)
	DeleteUser(Id string) error
	ListUser() ([]model.User, error)
	SearchUser(searchText string) ([]model.User, error)
}
