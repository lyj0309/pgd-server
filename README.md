#派工单服务端



## 运行方式
电脑下载go语言，在当前目录执行
`go run ./`
默认监听9000端口

## 目录文件说明
main.go 程序入口
controller/notice 通知
controller/order 订单
controller/user  用户
controller/main.go  返回封装
global 全局量，如数据库，以及表的封装

## 使用技术
+ mysql存储数据
+ jwt验证 
+ 库：gorm,gin,gin-jwt