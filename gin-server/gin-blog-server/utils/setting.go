package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	// System config
	PasswordMinLen int
	AdminEntrance  string
	AIsummaryURL   string
	AIsummaryKEY   string
	GitHubKey      string

	// Server config
	AppMode  string
	HttpPort string
	JwtKey   string
	LoginNum int

	// DataBase config
	Db         string
	DbHost     string
	DbPort     int
	DbUser     string
	DbPassword string
	DbName     string

	// Email server
	Stmphost      string
	Stmpport      string
	Emailusername string
	Emailpassword string
)

func init() {
	file, err := ini.Load("gin-blog-server/config/config.ini")
	if err != nil {
		fmt.Println("Configration file loading failed.")
	}
	LoadServer(file)
	LoadDatabase(file)
	LoadSystem(file)
	LoadEmaiInfo(file)
}

func LoadSystem(file *ini.File) {
	PasswordMinLen = file.Section("System").Key("PasswordMinLen").MustInt(6)
	AdminEntrance = file.Section("System").Key("AdminEntrance").MustString("admin")
	AIsummaryURL = file.Section("System").Key("AIsummaryURL").MustString("nil")
	AIsummaryKEY = file.Section("System").Key("AIsummaryKEY").MustString("nil")
	GitHubKey = file.Section("System").Key("GitHubKey").MustString("nil")
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("Server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("Server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("Server").Key("JwtKey").MustString("yourkey")
	LoginNum = file.Section("Server").Key("LoginNum").MustInt(10)
}

func LoadDatabase(file *ini.File) {
	Db = file.Section("DataBase").Key("Db").MustString("mysql")
	DbName = file.Section("DataBase").Key("DbName").MustString("NyaLog")
	DbHost = file.Section("DataBase").Key("DbHost").MustString("localhost")
	DbPort = file.Section("DataBase").Key("DbPort").MustInt(3306)
	DbUser = file.Section("DataBase").Key("DbUser").MustString("root")
	DbPassword = file.Section("DataBase").Key("DbPassword").MustString("IDontKnow")
}

func LoadEmaiInfo(file *ini.File) {
	Stmphost = file.Section("Email").Key("Stmphost").MustString("127.0.0.1")
	Stmpport = file.Section("Email").Key("Stmpport").MustString("587")
	Emailusername = file.Section("Email").Key("Emailusername").MustString("username")
	Emailpassword = file.Section("Email").Key("Emailpassword").MustString("password")
}
