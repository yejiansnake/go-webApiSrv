# go-webSrv
居于 echo 框架的二次封装

## 依赖的第三方库
    echo : github.com/labstack/echo
    Mysql : github.com/go-sql-driver/mysql
    go-yedb : github.com/yejiansnake/go-yedb
    yaml : github.com/go-yaml/yaml

## 代码结构说明
1. main.go: 程序启动主文件，一般不需要修改
2. model: 数据模型代码
3. controller: 业务逻辑代码，编写的 controller 结构需要在route中注册，必须实现setRoute方法
4. helper: 一些辅助功能，array 数组操作，convert 数据转换, page 数据分页
5. utility: 操作系统级别的功能封装
6. common: 全局常理声明定义
7. sys: 系统级别功能封装，如需增加配置在 src/sys/config 下编写(使用yaml格式，配置文件在 bin/config.yaml)
