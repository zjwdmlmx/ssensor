package global

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

const DBPath = "/var/lib/ssensor/data.db"

func initDB() {
	var err error

	DB, err = gorm.Open("sqlite3", DBPath)

	if err != nil {
		panic("failed to connect to databases")
	}
}

func init() {
	initDB()
	initRedis()
}
