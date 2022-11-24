package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"

	"github.com/booldesign/cli-demo/cmd"
	"github.com/booldesign/cli-demo/library/log"
	"github.com/booldesign/cli-demo/library/util"

	"github.com/urfave/cli/v2"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2022/11/24 09:42
 * @Desc:
 */

const (
	APPName    = "cli demo"
	APPVersion = "v0.0.1"
	APPUsage   = "对cli的使用演示"
)

func main() {

	w, err := util.CreateFile("./logs/" + APPName)
	if err != nil {
		panic(err)
	}
	log.LoadLogrus(w)

	app := &cli.App{
		Name:     APPName,
		Version:  APPVersion,
		Usage:    APPUsage,
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "weijianwen",
				Email: "booldesign@163.com",
			},
		},
		Copyright: "Copyright (c) 2022 BoolDesign",
		Commands: []*cli.Command{
			cmd.Crontab,
			cmd.Daemon,
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
