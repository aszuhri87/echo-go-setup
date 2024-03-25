package seeds

import (
	"echo-go/app/utils"
	"echo-go/config"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"column:id; gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Name     string    `gorm:"column:name;comment:''" json:"name"`
	Username string    `gorm:"column:username;comment:''" json:"username"`
	Password string    `gorm:"column:password;comment:''" json:"password"`
}

func (s Seed) UserSeed() {
	hash, _ := utils.HashPassword("123456")
	input := User{ID: uuid.New(), Name: "admin", Username: "admin", Password: hash}
	err := config.DB.Create(&input)

	if err != nil {
		panic(err)
	}

}
