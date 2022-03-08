# ginblog

## 项目介绍

    使用gin框架封装web服务

    我的博客地址： [**编码会馆**](https://www.seedblog.cn)

#### 热更新加载 github.com/cosmtrek/air 

    air是 Go 语言的热加载工具，它可以监听文件或目录的变化，自动编译，重启程序。大大提高开发期的工作效率。

### 性能统计 "github.com/gin-contrib/pprof"

    访问 如下地址 即可（端口号就是你的server）
    
    http://localhost:8022/debug/pprof/
    
    然后我们就可以在命令行中 分析了
    
    终端输入：
    
    go tool pprof http://127.0.0.1:8022/debug/pprof/profile
    
    然后等待30s（你可以在这30s内访问你的server 试试看）