package data

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"strings"
	"time"

	// import the entire database drivers
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Model Override gorm Model to change "id" on json restitution.
type Model struct {
	ID        uint       `gorm:"primary_key" json:"id" index:"id" important:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

var (
	DB_HOST       = ""
	DB_PATH       = "/data/database.sqlite"
	DB_USER       = ""
	DB_PASSWORD   = ""
	DB_DRIVER     = "sqlite3"
	DEBUG         = false
	numberOfTries = 30
)

// Migrate auto migrate data.
func Migrate() {
	log.Println("Migrate data")
	db := GetConn()
	defer db.Close()
	db.AutoMigrate(
		&Farmer{},
	)
}

func GetConn() *gorm.DB {
	DB, err := gorm.Open(DB_DRIVER, buildDBPath())
	for err != nil {
		time.Sleep(1 * time.Second)
		DB, err = gorm.Open(DB_DRIVER, buildDBPath())
		if err == nil || numberOfTries == 0 {
			log.Println(err)
			break
		}
		numberOfTries--
	}

	if err != nil {
		log.Println("fatal ", err)
	}
	if DEBUG {
		return DB.Debug()
	}
	return DB
}

func buildDBPath() string {
	path := DB_PATH
	switch strings.ToUpper(DB_DRIVER) {
	case "MYSQL":
		path = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local",
			DB_USER,
			DB_PASSWORD,
			DB_HOST,
			DB_PATH,
		)
	}
	return path
}
