package model

import (
	"NyaLog/gin-blog-server/utils/errmsg"
	"encoding/base64"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Uid          string `gorm:"type:varchar(20);not null;primary_key" json:"uid" label:"用户UID"`
	Username     string `gorm:"type:varchar(20);not null" json:"username" label:"用户名"`
	Password     string `gorm:"type:varchar(100);not null" json:"password" label:"用户密码"`
	Profilephoto string `gorm:"type:varchar(1000)" json:"profilephoto" label:"用户头像"`
	Email        string `gorm:"type:varchar(50);not null" json:"email" label:"用户邮箱"`
	Slogan       string `gorm:"type:varchar(50)" json:"slogan" label:"slogan"`
	Salt         string `gorm:"type:varchar(50);not null" json:"salt" label:"salt"`
	Secret       string `gorm:"type:varchar(300)" json:"secret" label:"secret"`
	Lastip       string `gorm:"type:varchar(20)" json:"lastip" label:"lastip"`
	Validateuser int    `gorm:"type:int(5);default:0" json:"validateuser" label:"用户是否验证"`
}

// 检查用户是否存在
// 1存在 0不存在
func UserExist() (bool, int) {
	var count int64
	var user User
	err := db.Find(&user).Count(&count).Error
	if err != nil {
		return false, errmsg.ERROR
	}
	if count > 0 {
		return true, errmsg.SUCCESS
	}
	return false, errmsg.SUCCESS
}

// 新增用户
func NewUser(user *User) int {
	pw, salt := ScryptPw(user.Password)
	user.Password = pw
	user.Salt = salt
	err := db.Create(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 提取用户信息
func SeleUser() (User, int) {
	var user User
	err := db.Limit(1).Find(&user).Error
	if err != nil {
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCESS
}

// 更新用户信息
func UpdateUser(uid string, data *User) int {
	var user User
	var usermap = make(map[string]interface{})
	usermap["username"] = data.Username
	usermap["profilephoto"] = data.Profilephoto
	usermap["email"] = data.Email
	usermap["slogan"] = data.Slogan
	usermap["secret"] = data.Secret
	err := db.Model(&user).Where("uid = ? ", uid).Updates(usermap).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 密码加密
// 输入一个字符串，是密码；返回一个字符串，是加密后的密码
func ScryptPw(password string) (string, string) {
	const KeyLen = 10 //加密后的字符串的长度
	// 这里建议盐用时间戳，从外部作为参数传入，不适用固定的盐
	// salt := make([]byte, 8) // scrypt的盐是一个字符数组，自行设置字符数组的长度
	var salt []byte
	rand.NewSource(time.Now().UnixNano())
	for i := 0; i < 8; i++ {
		randomNumber := rand.Intn(10)
		salt = append(salt, byte(randomNumber))
	}
	N := 32768 // 这个N是决定CPU和内存性能消耗的一个参数，要求大于1并且是2的幂次方
	R := 8
	P := 1 // 这两个是官方推荐的设置参数
	HashPw, err := scrypt.Key([]byte(password), salt, N, R, P, KeyLen)
	if err != nil {
		fmt.Println("Password encryption error")
	}
	pwd := base64.StdEncoding.EncodeToString(HashPw)
	strSalt := ""
	for _, b := range salt {
		strSalt += strconv.Itoa(int(b))
	}
	return pwd, strSalt
}

// 验证密码
func VertifyPw(password string) int {
	const KeyLen = 10
	var user User
	user, e := SeleUser()
	if e != errmsg.SUCCESS {
		return errmsg.ERROR
	}
	var salt []byte
	for _, b := range user.Salt {
		num, _ := strconv.Atoi(string(b))
		salt = append(salt, byte(num))
	}
	N := 32768 // 这个N是决定CPU和内存性能消耗的一个参数，要求大于1并且是2的幂次方
	R := 8
	P := 1 // 这两个是官方推荐的设置参数
	HashPw, err := scrypt.Key([]byte(password), salt, N, R, P, KeyLen)
	if err != nil {
		return errmsg.ERROR
	}
	if user.Password == base64.StdEncoding.EncodeToString(HashPw) {
		return errmsg.SUCCESS
	}
	return errmsg.ERROR
}

// 修改密码
func ModifyPw(password string) int {
	var user User
	data, _ := SeleUser()
	hashpw, salt := ScryptPw(password)
	var pwmap = make(map[string]interface{})
	pwmap["password"] = hashpw
	pwmap["salt"] = salt
	err := db.Model(&user).Where("uid = ?", data.Uid).Updates(pwmap).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除用户
func DeleteUser() int {
	var user User
	user, err := SeleUser()
	if err == errmsg.ERROR {
		return errmsg.ERROR
	}
	e := db.Where("uid = ?", user.Uid).Unscoped().Delete(&user).Error
	if e != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
