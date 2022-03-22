package dao

import (
	"backlog/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB // 全局变量，和数据库交互的桥梁
)

func InitMysql(cfg *setting.MySQLConfig) (err error) {
	// dsn := "root@tcp(127.0.0.1:3306)/backlog?charset=utf8mb4&parseTime=true&loc=Local"
	// dsn := "user:password@tcp(host:port)/database?charset=utf8mb4&parseTime=true&loc=Local"
	dsn := fmt.Sprintf("%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		cfg.User, cfg.Host, cfg.Port, cfg.DB)
	DB, err = gorm.Open("mysql", dsn) // 注意！使用:=会重新生成一个DB变量
	if err != nil {
		return
	}
	return DB.DB().Ping() // 测试数据库连通性
}

func Close() {
	DB.Close()
}
