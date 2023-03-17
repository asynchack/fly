package logger

import (
	"context"
	"io"
)

type Logger interface {
	/* jason-comment
	这里的Init()方法，特殊，不是一般意义上的，接收普通数据，然后做实例化，而是接收一个函数的调用！也可以是多个函数的调用！（这些函数调用，返回的还是一个函数）
	传入的一个个函数调用，把参数以闭包形式代入了内部的函数
	所以：Init（）接收的参数是一个函数列表， Init（）内部对这些函数进行调用；然后将其闭包形式代入的参数，赋值给logger，进行实例化



	*/
	Init(SetOptionsFuncs ...SetOptionsFunc) error

	Options() Options                            // 返回logger实例的Options信息
	String() string                              //返回logger实例的名称
	Fields(fields map[string]interface{}) Logger // 填充Options.Fields 字段
	Log(level Level, v ...interface{})
	Logf(level Level, format string, v ...interface{})
}

// 一个logger应该配置哪些参数，这里定义！
type Options struct {
	Level           Level                  // 什么级别上的日志记录
	Output          io.Writer              // 写到哪里？
	Fields          map[string]interface{} // 记录哪些字段？
	CallerSkipCount int                    // 最终记录日志行为是logger包里代码实现， 但一般记录日志发生行是在业务代码包，所以需要根据需要跳出几层函数调用
	Context         context.Context        // 啥用？

}

type SetOptionsFunc func(*Options) // 该函数接收一个Options的指针，从而，直接在Init（）方法中，访问到logger实例的Options地址，作为参数传入给这类函数
// 那么这类函数，就可以把它的闭包值，传到logger实例的Optionss中

func WithLevel(level Level) SetOptionsFunc {
	return func(o *Options) {
		o.Level = level // 这个level来自闭包的WithLevel调用时，传入的实参值
	}
}

func WithOutput(output io.Writer) SetOptionsFunc {
	return func(o *Options) {
		o.Output = output
	}
}

func WithFields(fields map[string]interface{}) SetOptionsFunc {
	return func(o *Options) {
		o.Fields = fields
	}
}

func WithSkipCallerCount(count int) SetOptionsFunc {
	return func(o *Options) {
		o.CallerSkipCount = count
	}
}

type Level int8

const (
	TraceLevel Level = iota - 2
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)
