package daemon

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/booldesign/cli-demo/daemon/internal/config"
	"github.com/booldesign/cli-demo/daemon/internal/server"
	"github.com/booldesign/cli-demo/daemon/internal/svc"
	"github.com/booldesign/cli-demo/library/database/mysql"
	"github.com/booldesign/cli-demo/library/database/redis"
	"github.com/booldesign/cli-demo/library/log"
	"github.com/booldesign/cli-demo/library/util"
	"github.com/booldesign/cli-demo/library/viperconf"
	model "github.com/booldesign/cli-demo/model/usercenter"

	"github.com/bwmarrin/snowflake"
	"github.com/sirupsen/logrus"
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
		&cli.StringFlag{
			Name:    "database",     // flag 名称
			Aliases: []string{"db"}, // 别名
			Usage:   "数据库地址",   // Usage
		},
		&cli.StringFlag{
			Name:    "redis",
			Aliases: []string{"r"},
			Usage:   "redis地址",
		},
	},
}

func runDaemon(ctx *cli.Context) error {
	c := Init(ctx)

	db, err := mysql.GetGormDB(ctx.Context, "user")
	if err != nil {
		return err
	}

	r := redis.GetRedis("user")

	svcCtx := svc.NewServiceContext(c, db, r)
	svr := server.NewUsercenterServer(svcCtx)

	node, err := snowflake.NewNode(1)
	if err != nil {
		return err
	}
	// Generate a snowflake ID.
	leaseId := node.Generate()
	lockKey := "insertLogLock"
	cancel, err := redis.TryLock(ctx.Context, r, lockKey, 30, 5, leaseId)
	if err != nil {
		logrus.Fatal("脚本", err)
	}
	go watchSignal()

	if err = svr.InsertLog(ctx.Context, &model.LogModel{
		Event:      "INSERT",
		CreateTime: time.Now(),
		Data:       `{"id":"1","name":"测试1","type":"备注"}`,
	}); err != nil {
		logrus.Error("insert log err ", err)
	}

	_, err = redis.UnLock(ctx.Context, r, cancel, lockKey, leaseId)
	if err != nil {
		logrus.Error("signal 释放锁err", err)
	}
	return nil
}

func Init(ctx *cli.Context) config.Config {
	var c config.Config
	viperconf.LoadConfig("./daemon/etc/app.yaml", &c)

	if ctx.String("db") != "" {
		c.DB.Url = ctx.String("db")
	}

	w, err := util.CreateFile(c.App.Log.Path + ctx.App.Name)
	if err != nil {
		panic(err)
	}
	log.LoadLogrus(c.App.Log.Level, w)

	// 连接数据库
	err = mysql.Assign("user", c.DB)
	if err != nil {
		panic(err)
	}

	// 连接redis
	redis.LoadRedis("user", c.Cache)

	return c
}

func watchSignal() {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sig:
	}

	logrus.Info("杀进程，触发signal")
}
