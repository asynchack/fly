package database 


func Setup(Driver string) {
	if Driver == "mysql" {
		db := new(mysql)
		db.setup()
	} 
	if Driver == "pgsql" {
		db := new(pgsql)
		db.setup()
	}
}


type mysql struct {}
type pgsql struct {}

func (m *mysql) setup() {
	// 根据config.DatabaseConfig.Driver 和 Source，可以初始化得到*gorm.DB实例，
	// 其实也是传参给gorm包去实例化的 
	// 实例化之后，就可以将数据库连接实例，赋值给global.DB变量了 
	
}

func (p *pgsql) setup() {
	// 同理mysql 
	// 至于fiy中还抽象了一层interface，目前没感觉什么好处，先不做接口抽象
}

// 另外可以把获取source字符串，抽成一个函数，getdriver字符串也抽一些，甚至open也抽一下 
// 不抽也行，抽的好处，暂时不理解