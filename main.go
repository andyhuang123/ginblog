package main

import (
	"fmt"
	"gin-blog/news/global"
	"gin-blog/news/initialize"
	"github.com/fatih/color"
	"go.uber.org/zap"
	"time"
)

func main() {
	//1.初始化yaml配置
	initialize.InitConfig()

	//2. 初始化routers
	Router := initialize.Routers()
	color.Cyan("go-gin服务开始了")

	// 3.初始化日志信息
	initialize.InitLogger()

	//zap.L().Info("this is hello func", zap.String("test", "test--------->"))

	//4. 初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}

	// 5.初始化mysql
	initialize.InitMysqlDB()

	//6. 初始化redis
	initialize.InitRedis()

	color.Yellow(">>>>>>>>>>>>>>>gin服务开始")
    global.Redis.Set("test","testvalue",time.Second)
	time.Sleep(time.Second*2)
	value := global.Redis.Get("test")
	color.Blue(value.Val())

	//启动gin,并配置端口,global.Settings.Port是yaml配置过的
	err := Router.Run(fmt.Sprintf("0.0.0.0:%d", global.Settings.Port))
	if err != nil {
		zap.L().Info("this is hello func", zap.String("error", "启动错误!"))
	}


}