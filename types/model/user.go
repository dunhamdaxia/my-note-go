package model

import (
	"errors"
	"fmt"
	"gintest/stores/mysql"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

// User 用户表
type User struct {
	ID              int64     `gorm:"primaryKey;unique;column:id;type:bigint(20);not null"`
	Name            string    `gorm:"column:name;type:varchar(100);not null"`
	Phone           string    `gorm:"column:phone;type:varchar(20);default:''"`
	Birthday        time.Time `gorm:"column:birthday;type:datetime"`
	CreatedAt       time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP"`
	Height          int       `gorm:"column:height;type:int(11);default:0"` // 身高 cm
	Weight          int       `gorm:"column:weight;type:int(11);default:0"` // 体重 kg
	Account         string    `gorm:"unique;column:account;type:varchar(20);not null"`
	Password        string    `gorm:"column:password;type:varchar(60);not null"`
	Token           string    `gorm:"column:token;type:varchar(100)"`
	TokenExpiration time.Time `gorm:"column:token_expiration;type:datetime"` // token过期时间
}

// TableName get sql table name.获取数据库表名
func (m *User) TableName() string {
	return "user"
}

func CreateUser(user User) (result User, err error) {
	if user.Name == "" {
		return result, errors.New("缺少必要的参数name")
	}

	if user.Account == "" {
		return result, errors.New("缺少必要的参数account")
	}

	if user.Password == "" {
		return result, errors.New("缺少必要的参数password")
	}

	db, err := mysql.NewDb()
	db.Where("account = ?", user.Account).First(&result)

	if result != (User{}) {
		return result, errors.New("account已存在")
	}

	password, err := hashAndSalt(user.Password)
	if err != nil {
		return result, errors.New("加密密码出错")
	}

	user.Password = password

	newUser := User{Name: user.Name, Account: user.Account, Password: user.Password}
	db.Create(&newUser)
	if db.Error != nil {
		return result, db.Error
	}

	return newUser, nil
}

func CheckToken(token string) (user User, code int, msg string) {
	if token == "" {
		return user, 0, "token不能为空"
	}

	db, err := mysql.NewDb()
	if err != nil {
		return user, 0, err.Error()
	}

	db.Where("token = ?", token).First(&user)

	if user == (User{}) {
		return user, 1, "未找到相关token"
	}

	if user.TokenExpiration.Unix() < time.Now().Unix() {
		return user, 1, "token已过期,请重新登录"
	}

	return user, 2, "ok"
}

func Login(account, password string) (string, error) {
	db, err := mysql.NewDb()
	if err != nil {
		return "", err
	}

	user := User{}
	db.Where("account = ?", account).First(&user)
	if user == (User{}) || !comparePasswords(user.Password, password) {
		return "", errors.New("账户或密码错误")
	}

	//刷新token
	now := time.Now().AddDate(0, 0, 30)
	tokenByte, err := bcrypt.GenerateFromPassword([]byte(fmt.Sprintf("ID%vTIME%v", user.ID, now.Unix())), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	tokenStr := string(tokenByte)

	db.Model(&user).Updates(User{Token: tokenStr, TokenExpiration: now, UpdatedAt: time.Now()})

	return tokenStr, nil
}

// 加密密码
func hashAndSalt(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// 验证密码
func comparePasswords(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
