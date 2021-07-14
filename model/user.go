package model

import "time"

type User struct {
	Id         uint      `json:"id"`
	NamePrefix string    `json:"namePrefix" gorm:"column:prefix"`
	FirstName  string    `json:"firstName" gorm:"type:varchar(50)"`
	LastName   string    `json:"lastName" gorm:"type:varchar(50)"`
	LineUserId string    `json:"lineUserId" gorm:"not null"`
	CitizenId  string    `json:"citizenId" gorm:"type:varchar(13)"`
	Phone      string    `json:"phone" gorm:"type:varchar(10)"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
