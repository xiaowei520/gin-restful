# gin-restful framework
Go web framework example.
>

> https://github.com/gin-gonic/gin

> Requires
- go1.8+
- gin

## 环境配置

##### 1.源码下载
```shell
$ cd $GOPATH/src
$ git clone git@github.com:xiaowei520/gin-restful.git
```


```
## 目录结构
```sh
conf            项目配置
middleware      中间件
model           模型，数据库连接&ORM
  └ orm         ORM扩展
models          数据库模型
apis            API入口文件
resources       项目资源记录
  └ db          数据库记录
router          路由
  └ api         接口路由
    ├context    自定义Context，便于扩展API层扩展
    └router     路由
util            公共工具
  ├ ext1       *
  ├ ext2       *
  └ ext3       *
```

```
## 运行
```sh
在当前gin-restful 目录下 go run *.go  编译加载路由

编译二进制项目
go build -o gin-restful-test-v1.0.0

