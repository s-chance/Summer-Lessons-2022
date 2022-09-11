package model

import (
	"demo/db"
)

type User struct {
	UserID      uint64  `gorm:"primaryKey" json:"user_id"`
	Name        string  `gorm:"not null" json:"name" validate:"required"`
	Email       string  `gorm:"not null" json:"email" validate:"required"`
	Phone       string  `json:"not null"`
	Password    string  `gorm:"not null" json:"password" validate:"required"`
	School      string  `json:"school"`
	StuID       string  `json:"stu_id"`
	ClassAmount uint64  `gorm:"not null; default:12" json:"class_amount"`
	WeekAmount  uint64  `gorm:"not null; default:16" json:"week_amount"`
	Classes     []Class `json:"classes"`
}

type APIUser struct {
	Name        string `gorm:"primaryKey" json:"name"`
	Email       string `gorm:"primaryKey" json:"email"`
	Phone       string `json:"phone"`
	School      string `json:"school"`
	StuID       string `jsonn:"stu_id"`
	ClassAmount uint64 `gorm:"not null; default:12" json:"class_amount"`
	WeekAmount  uint64 `gorm:"not null; default:16" json:"week_amount"`
}

func AddUser(user *User) error {
	//Omit忽略指定字段(一般是指不需要进行修改的字段)
	return db.DB.Omit("UserID", "Classes", "ClassAmount", "WeekAmount").Create(user).Error
}

func GetUserByEmail(email string) (*User, error) {
	var user *User
	err := db.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	} else {
		return user, nil
	}
}

func GetUserByID(id uint64) (*APIUser, error) {
	var user APIUser
	err := db.DB.Model(&User{}).Where("user_id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	} else {
		return &user, nil
	}
}

func UpdateUser(user *User) error {
	var u User
	err := db.DB.Where("user_id = ?", user.UserID).Take(&u).Error
	if err != nil {
		return err
	}
	return db.DB.Omit("UserID", "Classes").Save(user).Error
}

func DeleteUserByID(id uint64) error {
	return db.DB.Where("user_id = ?", id).Delete(&User{}).Error
}
