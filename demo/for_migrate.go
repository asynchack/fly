func run()  {
	// fly migrate -c config/settings.yml要执行的函数 


	if !generate {
		// 非生成，那么就是直接初始化数据表 
		// 1、先读取配置文件，加载配置信息 
		config.Setup(configFilePath)
		// 2、初始化得到数据库实例，成功后，赋值给global.DB实例，之后可全局任意包引用！
		database.Setup(config.DatabaseConfig.Driver) // 有一个global.DB可直接用了，是*gorm.DB类型

		// 3、migration包，的Migration结构体拿到global.DB实例后，做后续的migrate工作 
		migration.Migration.SetDB(global.DB) // 之后Migration的db就指向全局的globa.DB了

		// 4、调用数据库表初始化函数 
		if err := initDB(); err != nil {
			panic("migrate failed!")
		}


	}
}

func initDB() error {
	// global.DB 这里可以做一些基础配置，就是对*gorm.DB的配置 配置一些全局参数之类  
	migration.Migration.Migrate() // 实际开始迁移 

	// 先初始化一个version表，记录所有migrate的表，后面migrate过程中，如果已经有version记录了，表示该version对应的表已经初始化过，无需再初始化！
}


func (m *Migration)Migrate()  {
	// 前面应该还有对k排序，确保表初始化顺序有序
	for k, v := range m.versions {
		// 并且：迁移前先查数据库（有个表记录了已经初始化的表-用version记录）如果查到，那么就不必初始化该表了
		v(m.db, k) // 这里的v是func(db *gorm.DB, version string) 这里进行实际的表初始化工作 
	} 
}


func (m *Migration) SetVersion(version string, f func(db *gorm.DB, version string)){
	// 该函数被各个包中init（）被调用，把本包中的表初始化函数，注册到Migration的versions字段中 
	// 实际初始化表时，为什么还需要一个version字段呢？答：每次初始化完成之后，就在数据库的表versions中记录一条数据，表面该version对应的数据表初始化过了


}

type Migration struct {
	db *gorm.DB
	versions map[int]func(db *gorm.DB, version string) 

}