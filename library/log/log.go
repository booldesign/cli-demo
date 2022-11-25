package log

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2022/11/24 11:10
 * @Desc:
 */

/*
logrus 支持更多的日志级别：
	Panic：记录日志，然后panic。
	Fatal：致命错误，出现错误时程序无法正常运转。输出日志后，程序退出；
	Error：错误日志，需要查看原因；
	Warn：警告信息，提醒程序员注意；
	Info：关键操作，核心流程的日志；
	Debug：一般程序中输出的调试信息；
	Trace：很细粒度的信息，一般用不到
*/

// LoadLogrus 加载日志库
func LoadLogrus(level string, w io.Writer) {
	logrus.SetFormatter(&LogFormatter{})
	logrus.SetReportCaller(true) // 打印 log 产生的位置
	//设置日志级别
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		logLevel = logrus.WarnLevel
	}
	logrus.SetLevel(logLevel) // debug

	//设置日志输出多端
	out := []io.Writer{os.Stdout}
	if w != nil {
		out = append(out, w)
	}
	logrus.SetOutput(io.MultiWriter(out...))

	// Hook
	/*var dingHookUrl = "https://oapi.dingtalk.com/robot/send?access_token=xxxxxx"
	dh, err := hook.NewDingHook(dingHookUrl, nil)
	if err == nil {
		logrus.AddHook(dh)
	}*/

}
