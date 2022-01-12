package core

import (
	"context"
)

type (
	Property struct {
		Name      string `value:"name"`
		Dir       string `value:"dir"`
		MaxAge    int    `value:"max-age"` // 存活日期，单位天
		DebugMode bool   `value:"debug"`
		Std       bool   `value:"std"`
		File      bool   `value:"file"`
		TraceName string `value:"trace-Name"` // 会话追踪名称
	}

	Log interface {
		Named(string) Log
		Skip(skip int) Log
		Logger() interface{}
		Trace(ctx context.Context) context.Context

		Debug(ctx context.Context, msg string, keyAndValues ...interface{})
		Debugw(ctx context.Context, msg string, keyAndValues ...map[string]interface{})
		Info(ctx context.Context, msg string, keyAndValues ...interface{})
		Infow(ctx context.Context, msg string, keyAndValues ...map[string]interface{})
		Warn(ctx context.Context, msg string, keyAndValues ...interface{})
		Warnw(ctx context.Context, msg string, keyAndValues ...map[string]interface{})
		Error(ctx context.Context, msg string, keyAndValues ...interface{})
		Errorw(ctx context.Context, msg string, keyAndValues ...map[string]interface{})
	}
)
