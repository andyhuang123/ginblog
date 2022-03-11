# ginblog

## 项目介绍

使用gin框架封装web服务

我的博客地址： 

[**编码会馆**](https://www.seedblog.cn)

#### 热更新加载 github.com/cosmtrek/air 

    air是 Go 语言的热加载工具，它可以监听文件或目录的变化，自动编译，重启程序。大大提高开发期的工作效率。
    
    cd /项目根目录

    air -c  air.conf

#### docker部署
    编辑
    Dockerfile 文件

    构建镜像
    docker build --rm  -t hello:v1 .
    
    添加容器
    docker run --rm -it -p 127.0.0.1:8022:8022 hello:v1

### 性能统计 "github.com/gin-contrib/pprof"

    访问 如下地址 即可（端口号就是你的server）
    
    http://localhost:8022/debug/pprof/
    
    然后我们就可以在命令行中 分析了
    
    终端输入：
    
    go tool pprof http://127.0.0.1:8022/debug/pprof/profile
    
    然后等待30s（你可以在这30s内访问你的server 试试看）