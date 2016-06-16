package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("sqlite3", "/var/lib/ssensor/data.db")

	if err != nil {
		panic("failed to connect to databases")
	}
}
