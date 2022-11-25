package crontab

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/booldesign/cli-demo/crontab/internal/logic"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2022/11/24 10:11
 * @Desc:
 */

// Crontab crontab job list
var Crontab = &cli.Command{
	Name:    "crontab",
	Usage:   "crontab job list",
	Aliases: []string{"cron"},
	Action:  runCron,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "指定配置文件",
		},
		&cli.StringFlag{
			Name:    "func",
			Aliases: []string{"f"},
			Usage:   "指定执行的方法",
		},
	},
}

func runCron(ctx *cli.Context) error {

	logrus.Info("启动 crontab 服务...")

	traversal := logic.NewTraversalLogic(nil, nil)

	// 执行指定方法
	if f := ctx.String("func"); f != "" {
		var jobFuncList = map[string]func(){
			"traversalConsole": traversal.Traversal,
		}
		if execFunc, ok := jobFuncList[f]; ok {
			execFunc()
		} else {
			logrus.Error("指定的方法", f, "不存在")
		}

		return nil

	}

	// 定时脚本
	cron := cron.New(cron.WithSeconds())
	if _, err := cron.AddFunc("* * * * * *", traversal.Traversal); err != nil {
		logrus.Error("traversalConsole 执行错误，err:", err)
	}

	cron.Start()

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sig:
		logrus.Info("定时脚本已停止")
		cron.Stop()
	}

	return nil
}
