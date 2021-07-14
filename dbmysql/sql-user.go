package dbmysql

import (
	"log"
	"yana-golang/db_interface"
	"yana-golang/model"

	"gorm.io/gorm"
)

type userMysqlImpl struct {
	db *gorm.DB
	// createUser    *sql.Stmt
	// updateUser    *sql.Stmt
	// findByCitizen *sql.Stmt
	// getUser       *sql.Stmt
	// deleteUser    *sql.Stmt
}

func InitUserDB(db *gorm.DB) db_interface.UserDatabaseInterface {

	err := db.AutoMigrate(model.User{})
	if err != nil {
		log.Panic("init user db")
	}

	return &userMysqlImpl{
		db: db,
	}
}

func (s *userMysqlImpl) CreateUser(user model.User) (model.User, error) {
	result := s.db.Create(&user)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return user, nil
}
func (s *userMysqlImpl) UpdateUserByLineId(lineId string, user model.User) (model.User, error) {
	updUser := model.User{}
	result := s.db.Where("line_user_id = ?", lineId).Limit(1).Find(&updUser)

	if result.Error != nil {
		return model.User{}, result.Error
	}

	updUser.CitizenId = user.CitizenId
	updUser.FirstName = user.FirstName
	updUser.LastName = user.LastName
	updUser.NamePrefix = user.NamePrefix
	updUser.Phone = user.Phone

	s.db.Save(&updUser)
	return updUser, nil
}

// single row = Limit 1
// .First(&user) (asc), .Last(&user) (desc), .Take(&user)
// .First(&user, "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a")
// .Limit(1).Find(&user, primaryKey)

// multiple rows
// .Find(&user)

func (s *userMysqlImpl) GetUserByLineId(lineId string) (model.User, error) {
	user := model.User{}
	result := s.db.Where("line_user_id = ? ", lineId).Limit(1).Find(&user)

	if result.Error != nil {
		return model.User{}, result.Error
	}

	return user, nil
}

func (s *userMysqlImpl) GetUser(userId string) (model.User, error) {
	user := model.User{}
	result := s.db.Limit(1).Find(&user, userId)

	if result.Error != nil {
		// if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// 	return model.User{},
		// }
		return model.User{}, result.Error
	}

	return user, nil
}

func (s *userMysqlImpl) GetUserByCitizenId(citizenId string) (model.User, error) {
	user := model.User{}
	result := s.db.Where("citizen_id = ?", citizenId).Limit(1).Find(&user) // limit1, find = avoid not found

	if result.Error != nil {
		return model.User{}, result.Error
	}

	return user, nil
}

func (s *userMysqlImpl) DeleteUser(lineUserId string) error {

	return nil
}

func (s *userMysqlImpl) SearchUser(searchText string) ([]model.User, error) {
	users := []model.User{}
	searchQuery := "%" + searchText + "%"
	result := s.db.Where("first_name LIKE ? OR last_name LIKE ? OR phone LIKE ? OR citizen_id LIKE ? OR line_user_id LIKE ?", searchQuery, searchQuery, searchQuery, searchQuery, searchQuery).Find(&users)

	if result.Error != nil {
		return []model.User{}, result.Error
	}

	return users, nil
}

func (s *userMysqlImpl) ListUser() ([]model.User, error) {
	users := []model.User{}
	result := s.db.Find(&users) // Get all records
	if result.Error != nil {
		return []model.User{}, result.Error
	}

	return users, nil
}
