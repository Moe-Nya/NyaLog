package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	// Server config
	AppMode  string
	HttpPort string

	// DataBase config
	DbHost     string
	DbPort     int
	DbUser     string
	DbPassword string
	DbName     string
)

func init() {
	file, err := ini.Load("gin-blog-server/config/config.ini")
	if err != nil {
		fmt.Println("Configration file loading failed.")
	}
	LoadServer(file)
	LoadDatabase(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("Server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("Server").Key("HttpPort").MustString(":3000")
}

func LoadDatabase(file *ini.File) {
	DbName = file.Section("DataBase").Key("DbName").MustString("NyaLog")
	DbHost = file.Section("DataBase").Key("DbHost").MustString("localhost")
	DbPort = file.Section("DataBase").Key("DbPort").MustInt(3306)
	DbUser = file.Section("DataBase").Key("DbUser").MustString("root")
	DbPassword = file.Section("DataBase").Key("DbPassword").MustString("IDontKnow")
}
