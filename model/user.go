package model

import (
	"fmt"

	"github.com/binz96/blog/errmsg"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20); not null" json:"username" validate:"required,min=4,max=12"`
	Password string `gorm:"type:varchar(20); not null" json:"password" validate:"required,min=6,max=20"`
	Role     int    `gorm:"type:int; default:2" json:"role" validate:"required,gte=2"`
}

// name是否存在
func CheckUser(name string) int {
	var u User
	// db.First(&u, "username = ?", name)
	db.Select("id").First(&u, "username = ?", name) //不会更改u的其他字段
	fmt.Println(u)
	if u.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

func CreateUser(u *User) int {
	u.Password = CryptPw(u.Password)
	err = db.Create(&u).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询用户列表
func GetUsers(pageSize, pageNum int) []User {
	var users []User
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

func UpdateUser(id string, user *User) int {
	m := make(map[string]interface{})
	m["username"] = user.Username
	m["role"] = user.Role
	err = db.Model(&User{}).Where("id = ?", id).Updates(m).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func DeleteUser(id string) int {
	err = db.Delete(&User{}, id).Error
	// DELETE FROM users WHERE id = 10;
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func CryptPw(password string) string {
	crypted, _ := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(crypted)
}

func CryptCompare(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CheckLogin(username, password string) int {
	var u User
	db.Where("username = ?", username).First(&u)
	if u.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if !CryptCompare(CryptPw(password), password) {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if u.Role != 1 {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	return errmsg.SUCCESS
}
