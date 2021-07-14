package dbmock

import (
	"yana-golang/db_interface"
	"yana-golang/model"
)

type userMockImpl struct {
}

func InitUserDB() db_interface.UserDatabaseInterface {
	return &userMockImpl{}
}

func (s *userMockImpl) CreateUser(req model.User) error {
	return nil
}
func (s *userMockImpl) GetUser(Id string) (model.User, error) {
	return model.User{}, nil
}
func (s *userMockImpl) UpdateUser(Id string, req model.User) error {
	return nil
}
func (s *userMockImpl) DeleteUser(Id string) error {
	return nil
}
func (s *userMockImpl) ListUser() ([]model.User, error) {
	return []model.User{}, nil
}
