package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2022/11/24 10:15
 * @Desc:
 */

// Daemon service
var Daemon = &cli.Command{
	Name:    "daemon",
	Usage:   "Start daemon server",
	Aliases: []string{"daemon"},
	Action:  runDaemon,
	Flags: []cli.Flag{
		// --lang value, -l value  language for the greeting (default: "english")
		&cli.StringFlag{ // string
			Name:     "lang",        // flag 名称
			Aliases:  []string{"l"}, // 别名
			Value:    "english",     // 默认值
			Usage:    "language for the greeting",
			Required: true, // flag 必须设置
		},
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "指定配置文件",
		},
	},
}

func runDaemon(ctx *cli.Context) error {

	name := "who"
	// 获取arg参数长度
	if ctx.NArg() > 0 {
		name = ctx.Args().Get(0)
	}

	if ctx.String("lang") == "chinese" {
		fmt.Println("你好啊", name)
	} else {
		fmt.Println("Hello", name)
	}

	return nil
}
