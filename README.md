# Go-times
Go语言常用时间操作。

# List

- 获取现在时间
- 获取现在时间戳
- 获取格式化时间字符串
- 时间戳转格式化字符串
- 字符串转时间
- 字符串转时间戳

[tools](./tools)

# 常见坑
参考自：《Go语言编程之旅》

由于不同地区的时区不同，很可能存在错误，所以在进行时间处理时应该为其指定时区，例如：
```go
location, err := time.LoadLocation("Asia/Shanghai")
if err != nil{
	return time.Now()
}
return time.Now().In(location)
```
无论是格式化还是解析操作都应该指定好时区以防不必要的麻烦
```go
location, _ := time.LoadLocation("Asia/Shanghai")
inputTime := "2029-09-04 12:02:33"
layout := "2006-01-0 15:04:05"
t, _ := time.ParseInLocation(layout, inputTime, location)
dataTime := time.Unix(t.Unix(), 0).In(location).Format(layout)
```

# License
[MIT](./LICENSE)

Copyright (c) 2020 golang collection