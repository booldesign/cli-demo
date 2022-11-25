package hook

import (
	"flag"
	"testing"

	"github.com/sirupsen/logrus"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2022/11/24 12:22
 * @Desc:
 */

var test_webhook = flag.String("webhook", "", "DingTalk Webhook")

func TestNewDingHook(t *testing.T) {
	dh, err := NewDingHook(*test_webhook, nil)
	t.Log(*test_webhook, err)
	if err != nil {
		t.Fatal(err)
	}

	logrus.AddHook(dh)

	logrus.WithFields(logrus.Fields{
		"msgtype": "markdown",
		"markdown": map[string]string{
			"title": "钉钉hook测试!",
			"text":  "#### hook测试~~~  \n> 只是发个测试  \n> ![screenshot](https://www.baidu.com/img/flexible/logo/pc/peak-result.png)\n> *测试结束*",
		},
	}).Error()
}
