package migration

import (
	"fly/tools"
	"fmt"
	"path/filepath"
	"sort"
	"strconv"
	"sync"

	"github.com/spf13/cast"
	"gorm.io/gorm"
)

/* jason-comment
将迁移这个操作，抽象到一个类上：一次操作对应一个类的实例
	一次操作需要哪些内容呢？定义结构体Migrate中
		要操作的数据库连接实例 db *grom.DB
		要操作哪些表 versions map[int]func(db *gorm.DB, version string) map中的key记录表的名称，value指向实际操作这种表的动作（函数)
		锁？为什么？有并发吗？mutex sync.Mutex
			答：SetVersion方法时，因为时多个go文件init（）在使用（估计时并发）所以要加锁控制！


	一次操作需要哪些动作呢？ 定义Migrate的方法
		需要触发迁移吧：Migrate.Migrate()方法
		迁移前需要知道操作哪个数据库吧：把数据库连接实例赋给自己的字段，Migrate.SetDB(db *gorm.DB) ，与之对应来个Migrate.GetDB() *gorm.DB
		迁移中需要知道哪些表吧：得有方法注册表、和其表迁移函数到versions字段吧：Migrate.SetVersion(version int, f func(db *gorm.DB, version string))
			什么时候注册？另一个包version中定义表和其迁移函数，并在init（）中注册，这样只要该migration被导入，进而就先导入了version包，进行执行了其init（）函数，那么migration包被导入完成后，就可以直接SetDB（）然后Migrate（）了



*/

type Migrate struct {
	db       *gorm.DB
	versions map[int]func(db *gorm.DB, version string) error // int做key，方便排序
	mutex    sync.Mutex
}

// 实例化话一个
var MigrateInstance = &Migrate{
	versions: make(map[int]func(db *gorm.DB, version string) error, 0),
}

func (m *Migrate) SetDB(db *gorm.DB) {
	m.db = db
}

func (m *Migrate) GetDB() *gorm.DB {
	return m.db
}

func (m *Migrate) Migrate() {
	/* jason-comment
	先排序；针对versions中的key，确保表初始化，有序执行；
	借助一个中间的slice排序， 排序后遍历，对应取map中的函数，然后执行
	*/
	versions := make([]int, 0)
	for key, _ := range m.versions {
		versions = append(versions, key)

	}

	if !sort.IntsAreSorted(versions) {
		sort.Ints(versions) // 升序排列

	}
	var err error
	var count int64
	for _, version := range versions {
		// 执行前，先查表，判断此表是否已经执行过
		err = m.db.Debug().Table("sys_migration").Where("version = ?", version).Count(&count).Error
		if err != nil {
			fmt.Println(tools.Yellow("查询sys_migration失败-在迁移期间"))

		}
		if count > 0 {
			fmt.Println(tools.Yellow(fmt.Sprintf("该表已经迁移过了: %d", version)))
			continue

		}
		m.versions[version](m.db.Debug(), strconv.Itoa(version))
	}
}

func (m *Migrate) SetVersions(version int, f func(db *gorm.DB, version string) error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.versions[version] = f

}

func ConvertFileName2Int(fileName string) int {
	/* jason-comment
	把go文件，名称转为int类型

	*/
	s := filepath.Base(fileName)
	return cast.ToInt(s[:13])
}
