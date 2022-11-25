package plugin

import (
	"github.com/opentracing/opentracing-go"
	tracerLog "github.com/opentracing/opentracing-go/log"
	"gorm.io/gorm"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/5/24 11:05
 * @Desc:
 */

const (
	callBackBeforeName = "opentracing:before"
	callBackAfterName  = "opentracing:after"
)

type OpentracingPlugin struct{}

func (op *OpentracingPlugin) Name() string {
	return "opentracingPlugin"
}

func (op *OpentracingPlugin) Initialize(db *gorm.DB) (err error) {
	// 这里的CallBacks和模型的钩子不一样，CallBacks伴随GORM的DB对象整个生命周期，
	// 这里利用CallBacks对GORM框架进行侵入，以达到操作和访问GORM的DB对象的行为
	// Before
	_ = db.Callback().Create().Before("gorm:before_create").Register(callBackBeforeName, before)
	_ = db.Callback().Query().Before("gorm:query").Register(callBackBeforeName, before)
	_ = db.Callback().Delete().Before("gorm:before_delete").Register(callBackBeforeName, before)
	_ = db.Callback().Update().Before("gorm:setup_reflect_value").Register(callBackBeforeName, before)
	_ = db.Callback().Row().Before("gorm:row").Register(callBackBeforeName, before)
	_ = db.Callback().Raw().Before("gorm:raw").Register(callBackBeforeName, before)

	// After
	_ = db.Callback().Create().After("gorm:after_create").Register(callBackAfterName, after)
	_ = db.Callback().Query().After("gorm:after_query").Register(callBackAfterName, after)
	_ = db.Callback().Delete().After("gorm:after_delete").Register(callBackAfterName, after)
	_ = db.Callback().Update().After("gorm:after_update").Register(callBackAfterName, after)
	_ = db.Callback().Row().After("gorm:row").Register(callBackAfterName, after)
	_ = db.Callback().Raw().After("gorm:raw").Register(callBackAfterName, after)
	return
}

// 包内静态变量
const (
	operationName = "gorm"
	gormSpanKey   = "gormSpan"
)

func before(db *gorm.DB) {
	// 在每次SQL操作前从context上下文生成子span
	// 先从父级spans生成子span，这里命名为gorm
	span, _ := opentracing.StartSpanFromContext(db.Statement.Context, operationName)

	// 利用GORM实例去传递span
	db.InstanceSet(gormSpanKey, span)
}

func after(db *gorm.DB) {
	// 在每次SQL操作后从DB实例拿到Span并记录数据
	_span, isExist := db.InstanceGet(gormSpanKey)
	if !isExist {
		return
	}
	span, ok := _span.(opentracing.Span)
	if !ok {
		return
	}
	defer span.Finish()

	if db.Error != nil {
		span.LogFields(tracerLog.Error(db.Error))
	}
	sql := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)
	// 调用span的LogFields方法记录sql
	span.LogFields(tracerLog.String("sql", sql))
}
