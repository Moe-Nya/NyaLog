package model

import (
	"NyaLog/gin-blog-server/utils"
	"fmt"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB
var err error

func InitDb() {
	// %s是占位符 分别代表用户名、密码、地址、端口、数据库名
	// 将数据引入，自动填充占位符
	dbInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassword,
		utils.DbHost,
		strconv.Itoa(utils.DbPort),
		utils.DbName,
	)

	db, err = gorm.Open(mysql.Open(dbInfo), &gorm.Config{
		// gorm日志模式：silent
		Logger: logger.Default.LogMode(logger.Silent),
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})

	if err != nil {
		fmt.Println("DataBase connect failed.", err)
		os.Exit(1)
	}

	// db.AutoMigrate(&User{}, &EmailServer{}, &BlogSet{}, &FindMe{}, &Navigation{}, &Category{}, &Article{}, &Comment{})
	db.AutoMigrate(&User{}, &BlogSet{}, &FindMe{}, &Navigation{}, &Category{}, &Article{}, &Comment{})

	sqlDB, _ := db.DB()
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(30 * time.Second)
}
