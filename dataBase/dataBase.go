package dataBase

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
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

type User struct {
	gorm.Model
	User     string  `gorm:"varchar(10);not null"`
	Password string  `gorm:"not null"`
	CardID   string  `gorm:"varchar(8);not null"`
	Money    float64 `gorm:"not null; check:money >= 0"`
}

type Data struct {
	gorm.Model
	Temperature float64 `gorm:"not null" json:"temperature"`
	Tds         float64 `gorm:"not null" json:"Tds"`
	//Personnel   string  `gorm:"varchar(15);not null" json:"personnel"`
	CreatedTime string `gorm:"not null" json:"created_time"`
	//People      uint16  `gorm:"not null" json:"people"`
	//Day         uint16  `gorm:"not null" json:"day"`
}

type Status struct {
	gorm.Model
	Lamp int `gorm:"not null"`
	//Loud   int    `gorm:"not null"`
	TemH string `gorm:"not null"`
	TemL string `gorm:"not null"`
	TdsH string `gorm:"not null"`
	//LightL string `gorm:"not null"`
	Time string `gorm:"not null"`
	//ST     string `gorm:"not null"`
	//People int `gorm:"not null"`
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

	db.AutoMigrate(&User{}, &Data{}, &Status{})

	DB = db
}
