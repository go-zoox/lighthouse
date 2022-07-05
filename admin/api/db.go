package api

import (
	"os"
	"time"

	"github.com/go-zoox/fs"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

type DNS struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	// gorm.Model
	// 记录类型
	Type string `json:"type"`
	// 记录域名
	Host string `json:"host"`
	// 记录值
	Value string `json:"value"`
	//
	IsActive bool `json:"is_active"`
	// 记录备注
	Description string `json:"description"`
}

func GetDB() *gorm.DB {
	if Db == nil {
		pwd, _ := os.Getwd()
		db, err := gorm.Open(sqlite.Open(fs.JoinPath(pwd, "data/sqlite3.db")), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		Db = db

		db.AutoMigrate(&DNS{})
	}

	return Db
}

func DBList(page int, pageSize int) (int64, []DNS) {
	db := GetDB()

	var data []DNS
	var total int64

	db.Find(&data).Count(&total)

	db.
		Order("updated_at desc").
		Find(&data).
		Offset(int((page - 1) * pageSize)).
		Limit(int(pageSize))

	return total, data
}

func DBCreate(dns *DNS) {
	db := GetDB()

	db.Create(&dns)
}

func DBRetrieve(id int) *DNS {
	db := GetDB()

	var dns DNS
	if err := db.First(&dns, id).Error; err != nil {
		return nil
	}

	return &dns
}

func DBUpdate(dns *DNS) {
	db := GetDB()

	db.Save(&dns)
}

func DBDelete(id int) {
	db := GetDB()

	var dns DNS
	if err := db.First(&dns, id).Error; err != nil {
		return
	}

	db.Delete(&dns)
}

func DBAll() []DNS {
	db := GetDB()

	var data []DNS

	db.
		Order("updated_at desc").
		Find(&data)

	return data
}
