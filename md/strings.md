## string
+ string 的地层实现
+ string 和 []byte 的转换实现

### 运行时底层结构
```go

type stringStruct struct {
    str unsafe.Pointer // 字符串内容存储地址
    len int // 字符串长度
}

```

### 声明/初始化
```go
var s string // 声明一个 string 类型变量，默认值为空 
s := "" // 初始化一个值为空的 string 类型变量
```

#### 追加元素
```go
var s  = "hello"
s += " word"
```


#### 字符串转[]byte黑科技
```go
func stringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}
```

#### ps
+ string 底层结构 len 记录的是字节长度而不是字符长度
+ 获取 string 字符使用 len([]rune(s))
+ 字符串 range 按字符遍历，下标返回字符第一个字节位置
+ 字符串 s[10] 取得是第11个字节
+ 字符串是常量，值不可更改，字符串变量赋值操作实际是让变量指向新的地址


#### 参考
源码位置 runtime/string.go