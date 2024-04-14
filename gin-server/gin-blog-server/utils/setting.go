package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

var (
	// System config
	Domain         string
	PasswordMinLen int
	AdminEntrance  string
	GPTURL         string
	GPTKey         string
	QWURL          string
	QWKey          string
	GitHubID       string
	GitHUbSecret   string

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
	Smtphost      string
	Smtpport      string
	Emailusername string
	Emailpassword string
)

func init() {
	configPath := "gin-blog-server/config/config.ini"
	if configPathEnv := os.Getenv("CONFIG_PATH"); configPathEnv != "" {
		configPath = configPathEnv
	}
	file, err := ini.Load(configPath)
	if err != nil {
		fmt.Println("Configration file loading failed.")
	}
	LoadServer(file)
	LoadDatabase(file)
	LoadSystem(file)
	LoadEmaiInfo(file)
}

func LoadSystem(file *ini.File) {
	Domain = file.Section("System").Key("Domain").MustString("http://localhost:8080")
	PasswordMinLen = file.Section("System").Key("PasswordMinLen").MustInt(6)
	AdminEntrance = file.Section("System").Key("AdminEntrance").MustString("admin")
	GPTURL = file.Section("System").Key("GPTURL").MustString("nil")
	GPTKey = file.Section("System").Key("GPTKey").MustString("nil")
	QWURL = file.Section("System").Key("QWURL").MustString("nil")
	QWKey = file.Section("System").Key("QWKey").MustString("nil")
	GitHubID = file.Section("System").Key("GitHubID").MustString("nil")
	GitHUbSecret = file.Section("System").Key("GitHUbSecret").MustString("nil")
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
	Smtphost = file.Section("Email").Key("Smtphost").MustString("127.0.0.1")
	Smtpport = file.Section("Email").Key("Smtpport").MustString("587")
	Emailusername = file.Section("Email").Key("Emailusername").MustString("username")
	Emailpassword = file.Section("Email").Key("Emailpassword").MustString("password")
}
