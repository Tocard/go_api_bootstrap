package data

import (
	"github.com/jinzhu/gorm"
	"log"
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
		&PoolAdmin{}, &FloatingPartial{},
	)
}

func GetConn() *gorm.DB {
	DB, err := gorm.Open(DB_DRIVER, buildDBPath())
	DB.SingularTable(true)
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
	return path
}
