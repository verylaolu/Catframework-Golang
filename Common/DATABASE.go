package Common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"        // 用于记录日志
	"sso/Config" // 引入配置文件相关
)

var DB *gorm.DB

// SetDB func
func SetDB(database *gorm.DB) {
	DB = database
}

// ConnectToDB func
func ConnectToDB() *gorm.DB {
	// 读取配置信息
	connectingStr := Config.GetMysqlConnectingString()
	fmt.Println(connectingStr)
	// 记录日志
	log.Println("Connet to db...")

	// gorm数据库操作，手册：http://gorm.book.jasperxu.com/
	db, err := gorm.Open("mysql", connectingStr)

	if err != nil {
		log.Println(err)
	}else{
		log.Println("Connet to db success")
	}
	// 全局禁用表名复数
	// 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
	db.SingularTable(true)
	return db
}
