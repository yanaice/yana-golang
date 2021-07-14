package dbmongo

import (
	"yana-golang/db_interface"
	"yana-golang/model"
)

type userMongoImpl struct {
}

func InitUserDB() db_interface.UserDatabaseInterface {
	return &userMongoImpl{}
}

func (s *userMongoImpl) CreateUser(req model.User) error {
	return nil
}
func (s *userMongoImpl) GetUser(Id string) (model.User, error) {
	return model.User{}, nil
}
func (s *userMongoImpl) UpdateUser(Id string, req model.User) error {
	return nil
}
func (s *userMongoImpl) DeleteUser(Id string) error {
	return nil
}
func (s *userMongoImpl) ListUser() ([]model.User, error) {
	return []model.User{}, nil
}
