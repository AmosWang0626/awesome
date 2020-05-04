# GO 2020

> 利用2020年春节休息期间, 系统地学习golang。(武汉加油！中国加油！)
>
> [Golang标准库文档](https://studygolang.com/pkgdoc)

## GO学习归档
> src/main/archive

- base >>> json、切片、指针、异常、面向对象(封装、继承、多态)
- ds >>> [数据结构与算法](https://github.com/AmosWang0626/awesome/tree/master/src/main/archive/ds)
- file >>> 文件操作
- function >>> 函数
- goroutine >>> 协程
- redis >>> Redis
- reflect >>> 反射
- socket >>> TCP socket
- struct >>> 结构体 & 结构体方法
- time >>> 时间

## GO应用程序
> src/main/application

### 1.间歇提醒休息
> 结合《为什么精英都是时间控》碰撞点小火花
>
> [间歇提醒休息](src/main/application/elite_time/EliteTimeControl.go)

### 2.IM
- 相关技术：tcp通讯、[自定义协议](src/main/application/im/common/utils/transfer.go)、Redis([连接池](src/main/application/im/common/utils/redis.go))
- 已实现功能
    - 登录
    - 注册
    - 退出
    - 在线用户列表
    - 查看所有注册用户
- 未实现功能
    - 点对点聊天
    - 离线留言