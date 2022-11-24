# 演示urfave/cli的使用，cron的使用

## urfave/cli文档
https://cli.urfave.org/v2/getting-started/

## cron定时任务
https://learnku.com/docs/learngo/1.0/cron-ding-shi-ren-wu/13973

## flag设置
```
// --lang value, -l value  language for the greeting (default: "english")
&cli.StringFlag{ // string
    Name:        "lang",        // flag 名称
    Aliases:     []string{"l"}, // 别名
    Value:       "english",     // 默认值
    Usage:       "language for the greeting",
    Destination: &language, // 指定地址，如果没有可以通过 *cli.Context 的 GetString 获取
    Required:    true,      // flag 必须设置
},
```
运行一下

需要先传全局GLOBAL OPTIONS
```
go run main.go daemon --lang chinese weijianwen
或
go run main.go daemon -l chinese weijianwen
```

```
go run main.go crontab
或
go run main.go cron

go run main.go cron -func traversalConsole
```