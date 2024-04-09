package dataBase

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var DB *gorm.DB //数据库

type DBSeting struct {
	Type     string
	User     string
	Password string
	Host     string
	Port     string
	Name     string
	Charset  string
}

// Admin 表结构：饮水机表
type Admin struct {
	ID       uint   `gorm:"primaryKey"` // 主键
	Username string // 用户名
	Password string // 密码
	Status   bool   // 状态
}

// WaterDispenser 表结构：饮水机表
type WaterDispenser struct {
	ID          uint    `gorm:"primaryKey"` // 主键
	Price       float64 // 价格 元/升
	DispenserID string  // 饮水机ID
	Model       string  // 型号
	Location    string  // 安装位置
}

// User 表结构：用户表
type User struct {
	ID       uint    `gorm:"primaryKey"` // 主键
	User     string  // 用户名
	Password string  // 密码
	Card     string  //卡号
	Balance  float64 // 余额
}

// Transaction 表结构：消费记录表
type Transaction struct {
	ID              uint      `gorm:"primaryKey"` // 主键
	User            string    // 用户
	DispenserID     uint      // 饮水机ID
	Amount          float64   // 金额
	TransactionTime time.Time `gorm:"default:CURRENT_TIMESTAMP"` // 消费时间，默认为当前时间
}

// DispenserStatus 表结构：饮水机状态表
type DispenserStatus struct {
	ID          uint      `gorm:"primaryKey"` // 主键
	DispenserID uint      // 饮水机ID
	Status      string    // 状态
	Temperature float64   // 水温
	TDS         float64   // TDS水质
	Flow        bool      // 是否出水状态
	RecordTime  time.Time `gorm:"default:CURRENT_TIMESTAMP"` // 记录时间，默认为当前时间
}

// InitDB 连接数据库
func InitDB() {

	// 创建一个 DBSetting 变量并为其字段赋值
	dbSeting := DBSeting{
		Type:     "mysql",
		User:     "root",
		Password: "123456",
		Host:     "127.0.0.1",
		Port:     "3306",
		Name:     "Cloud",
		Charset:  "utf8mb4",
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		dbSeting.User,
		dbSeting.Password,
		dbSeting.Host,
		dbSeting.Port,
		dbSeting.Name,
		dbSeting.Charset,
	)

	db, err := gorm.Open(dbSeting.Type, dsn)
	if err != nil {
		panic("failed to connect database, err:" + err.Error())
		return
	} else {
		// 打印构建的数据源名称
		fmt.Println("DSN:", dsn)
	}

	// 自动迁移表结构
	db.AutoMigrate(&Admin{}, &WaterDispenser{}, &Transaction{}, &User{}, &DispenserStatus{})

	DB = db
}
