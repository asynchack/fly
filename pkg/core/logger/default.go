package logger

import "sync"

/* jason-comment
1、定义数据结构defaultLogger、及其方法，实现Logger接口的6个方法；
2、定义NewLogger函数，是defaultLogger的构造器

3、init（）函数中，调用NewLogger函数，实例化一个默认的defaultLogge的实例， */

// logger需要哪些数据？读写锁？ 和一个Options类型的字段即可
type defaultLogger struct {
	sync.RWMutex // why？要有读写锁？
	opts         Options
}

func (l *defaultLogger) Init(SetOptionsFuncs ...SetOptionsFunc) {
	for _, fc := range SetOptionsFuncs {
		fc(&l.opts) // fc 是func(*Options)的函数，将l的opts字段传递给它，然后该函数内部会对l.opts的字段进行赋值-填充
	}

}

func (l *defaultLogger) Fields(fields map[string]interface{}) Logger {
	// 为什么要上锁？
	l.Lock()
	l.opts.Fields = copyFields(fields)
	l.Unlock()
	return l

}

func copyFields(fields map[string]interface{}) {

}

func (l *defaultLogger) String() string {
	return "default"
}

func (l *defaultLogger) Options() Options {
	return l.opts

}

func (l *defaultLogger) Log(level Level, v ...interface{}) {

}

func (l *defaultLogger) Logf(level Level, format string, v ...interface{}) {

}

/* jason-comment

Options() Options                            // 返回logger实例的Options信息
String() string                              //返回logger实例的名称
Fields(fields map[string]interface{}) Logger // 填充Options.Fields 字段
Log(level Level, v ...interface{})
Logf(level Level, format string, v ...interface{})
*/
