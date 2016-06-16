package model

import (
	"time"

	"github.com/zjwdmlmx/ssensor/db"
)

type HistoryData struct {
	Id        uint64    `gorm:"primary_key"`
	Uid       string    `gorm:"type:char(64);not null"`
	Time      time.Time `gorm:"not null"`
	Longitude float64   `gorm:"not null"`
	Latitude  float64   `gorm:"not null"`
	State     uint8     `gorm:"not null"`
}

type historyDataModel struct{}

var HistoryDataModel historyDataModel

func (historyDataModel) CreateManay(rows []HistoryData) {
	tx := db.DB.Begin()

	for _, row := range rows {
		db.DB.Create(&row)
	}

	defer func() {
		tx.Commit()
	}()
}

func (historyDataModel) CreateOne(row *HistoryData) {
	db.DB.Create(row)
}
