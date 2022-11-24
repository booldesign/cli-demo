package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2022/11/24 13:55
 * @Desc:
 */

type LogFormatter struct {
	Message string      `json:"message"`
	TraceId interface{} `json:"traceId,omitempty"`
	UserId  interface{} `json:"userId,omitempty"`
	Caller  string      `json:"caller"`
	File    string      `json:"file"`
}

// Format 格式化日志
func (f *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	f.Message = entry.Message
	if _, ok := entry.Data["traceId"]; ok {
		f.TraceId = entry.Data["traceId"]
	}
	if _, ok := entry.Data["userId"]; ok {
		f.UserId = entry.Data["userId"]
	}
	if entry.HasCaller() {
		f.Caller = entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line)
		if fileVal != "" {
			f.File = fileVal
		}
	}
	jsonData := ""
	d, err := json.Marshal(f)
	if err == nil {
		jsonData = string(d)
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05.000")
	var newLog string
	newLog = fmt.Sprintf("[%s] [%s] %s\n", timestamp, entry.Level, jsonData)

	b.WriteString(newLog)
	return b.Bytes(), nil
}
