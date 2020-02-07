# go 学习 2020

## init project
- go mod init amos.wang/awesome
- 会生成 `go.mod`

## 一、基础相关
1. Hello World!
    ```go
    package main
    
    import "fmt"
    
    func main() {
        fmt.Printf("Hello World!")
    }
    ```

2. Go环境配置
    - GOROOT: golang安装目录
    - GOBIN: 可执行文件生成目录
    - GOPATH: workspace（可设置多个）

3. `go build hello.go`
    - 编译后会生成`hello.exe`，非`Windows`系统则不会带`.exe`后缀，皆为可执行文件
    - 执行：`./hello.exe`
    - 重命名生成的文件(注意：名字一定要为`*.exe`)：`go build -o myhello.exe hello.go`
    - 执行：`./myhello.exe`

4. `go run hello.go`
    - 直接运行`.go`文件，需要`go`执行环境
    - go 语言 定义的变量或者 import 的包如果没有使用到，代码不能编译通过
        - 引入未使用：`imported and not used: "fmt"`
        - 定义未使用：`subj declared and not used`

5. 转义字符
    ```
    \t 制表符
    \n 换行
    \\ \
    \" "
    \r 比较奇葩 fmt.Println("落霞与孤鹜齐飞\r滕王阁序") 输出：滕王阁序鹜齐飞
    ```

6. 代码格式化`gofmt`
    - 打印到控制台`gofmt hello.go`
    - 写入源文件`gofmt -w hello.go`

7. 值类型 和 引用类型

    - 值类型：int 系列、float系列、bool、string、数组、结构体struct
    - 引用类型：指针、slice切片、map、管道chan、interface等

8. 作用域

    - 局部变量：本方法有效
    - 全局变量：首字母小写，本包下都有效；首字母大写，整个程序下都有效
    
9. Go主线程和Go协程

    - 有独立的栈空间
    - 共享程序堆空间
    - 调度由程序控制
    - 协程是轻量级线程

10. xxx

11. xxx


## 二、GO占位符
1. 普通占位符

    | 占位符 | 说明 | 举例 | 输出|
    | --- | --- | --- | --- |
    | %v  | 相应值的默认格式。          | Printf("%v", people) | {zhangsan}，|
    | %+v | 打印结构体时，会添加字段名    | Printf("%+v", people) | {Name:zhangsan} |
    | %#v | 相应值的Go语法表示          | Printf("#v", people) | main.Human{Name:"zhangsan"} |
    | %T  | 相应值的类型的Go语法表示     | Printf("%T", people) | main.Human |
    | %%  | 字面上的百分号，并非值的占位符 | Printf("%%")         | % |

2. 整数占位符

    | 占位符 | 说明 | 举例 | 输出|
    | --- | --- | --- | --- |
    |%b | 二进制表示                           |  Printf("%b", 5)       | 101 |
    |%c | 相应Unicode码点所表示的字符            |  Printf("%c", 0x4E2D)  | 中 |
    |%d | 十进制表示                           |  Printf("%d", 0x12)    | 18 |
    |%o | 八进制表示                           |  Printf("%d", 10)      | 12 |
    |%q | 单引号围绕的字符字面值，由Go语法安全地转义 | Printf("%q", 0x4E2D)   | '中' |
    |%x | 十六进制表示，字母形式为小写 a-f         | Printf("%x", 13)       | d |
    |%X | 十六进制表示，字母形式为大写 A-F         | Printf("%x", 13)       | D |
    |%U | Unicode格式：U+1234，等同于 "U+%04X"  | Printf("%U", 0x4E2D)   | U+4E2D |

3. 浮点数和复数的组成部分（实部和虚部）

    | 占位符 | 说明 | 举例 | 输出|
    | --- | --- | --- | --- |
    | %b | 无小数部分的，指数为二的幂的科学计数法，  与 strconv.FormatFloat 的 'b' 转换格式一致。例如 -123456p-78|
    | %e | 科学计数法，例如 -1234.456e+78       | Printf("%e", 10.2)    | 1.020000e+01 |
    | %E | 科学计数法，例如 -1234.456E+78       | Printf("%e", 10.2)    | 1.020000E+01 |
    | %f | 有小数点而无指数，例如 123.456        | Printf("%f", 10.2)    | 10.200000 |
    | %g | 根据情况选择 %e 或 %f 以产生更紧凑的（无末尾的0）输出 | Printf("%g", 10.20)  | 10.2 |
    | %G | 根据情况选择 %E 或 %f 以产生更紧凑的（无末尾的0）输出 | Printf("%G", 10.20+2i) | (10.2+2i) |

4. 字符串与字节切片

    | 占位符 | 说明 | 举例 | 输出|
    | --- | --- | --- | --- |
    | %s | 输出字符串表示（string类型或[]byte) | Printf("%s", []byte("Go语言")) | Go语言 |
    | %q | 双引号围绕的字符串，由Go语法安全地转义 | Printf("%q", "Go语言")         | "Go语言" |
    | %x | 十六进制，小写字母，每字节两个字符    | Printf("%x", "golang")         | 676f6c616e67 |
    | %X | 十六进制，大写字母，每字节两个字符    | Printf("%X", "golang")         | 676F6C616E67 |

5. 指针

    | 占位符 | 说明 | 举例 | 输出|
    | --- | --- | --- | --- |
    | %p | 十六进制表示，前缀 0x | Printf("%p", &people) | 0x4f57f0 |

6. 布尔占位符：

    | 占位符 | 说明 | 举例 | 输出|
    | --- | --- | --- | --- |
    | %t  | true 或 false | Printf("%t", true) | true |
