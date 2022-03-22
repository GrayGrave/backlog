package main

import (
	"backlog/dao"
	"backlog/models"
	"backlog/routers"
	"backlog/setting"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		// 可以利用flag包实现 -h 给出不同参数用法的功能，此处只是简单进行终端打印而已
		// 直接go run main.go,命令行获取的参数个数为1个:（涉及 golang 编译运行内部原理）
		// /var/folders/qz/2zdnfpfd4ln1bf1wcy_svwf00000gn/T/go-build931384225/b001/exe/main
		fmt.Printf("Usage: ./backlog  conf/config.ini")
		return
	}

	// 加载配置文件
	if err := setting.Init(os.Args[1]); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}

	// 创建数据库 sql: create database backlog;
	// 连接数据库
	err := dao.InitMysql(setting.Conf.MySQLConfig)
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
	}

	// 程序退出前关闭数据库连接
	defer dao.Close()

	// 模型绑定,自动创建表todos
	dao.DB.AutoMigrate(&models.Todo{})

	// 路由注册
	r := routers.SetupRouter()

	// 启动服务
	if err := r.Run(fmt.Sprintf("localhost:%d", setting.Conf.Port)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}
