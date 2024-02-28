package dataBase

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB //数据库

type DBSetting struct {
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
	User     string `gorm:"varchar(10);not null"`
	Password string `gorm:"not null"`
}

type Data struct {
	gorm.Model
	Sunlit      float64 `gorm:"not null" json:"sunlit"`
	Temperature float64 `gorm:"not null" json:"temperature"`
	Personnel   string  `gorm:"varchar(15);not null" json:"personnel"`
	CreatedTime string  `gorm:"not null" json:"created_time"`
	People      uint16  `gorm:"not null" json:"people"`
	Day         uint16  `gorm:"not null" json:"day"`
	//A           uint16  `gorm:"not null" json:"a"`
	//B           uint16  `gorm:"not null" json:"b"`
	//C           uint16  `gorm:"not null" json:"c"`
	//D           uint16  `gorm:"not null" json:"d"`
}

type Status struct {
	gorm.Model
	Lamp   int    `gorm:"not null"`
	Loud   int    `gorm:"not null"`
	TemH   string `gorm:"not null"`
	TemL   string `gorm:"not null"`
	LightH string `gorm:"not null"`
	LightL string `gorm:"not null"`
	Time   string `gorm:"not null"`
	ST     string `gorm:"not null"`
	//People int `gorm:"not null"`
}

//type Time struct {
//	gorm.Model
//	Time int `gorm:"not null"`
//	//People int `gorm:"not null"`
//}

//type File struct {
//	gorm.Model
//	Filename string `gorm:"column:filename"` // 文件名
//	Content  []byte `gorm:"column:content"`  // 文件内容
//}

// InitDB 连接数据库
func InitDB() {

	// 创建一个 DBSetting 变量并为其字段赋值
	dbSettings := DBSetting{
		Type:     "mysql",
		User:     "root",
		Password: "123456",
		Host:     "127.0.0.1",
		Port:     "3306",
		Name:     "database",
		Charset:  "utf8",
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		dbSettings.User,
		dbSettings.Password,
		dbSettings.Host,
		dbSettings.Port,
		dbSettings.Name,
		dbSettings.Charset,
	)

	db, err := gorm.Open(dbSettings.Type, dsn)
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
