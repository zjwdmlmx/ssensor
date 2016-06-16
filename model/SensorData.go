package model

import "github.com/zjwdmlmx/ssensor/db"

type SensorData struct {
	Id    uint64  `gorm:"primary_key;AUTO_INCREMENT"`
	Uid   string  `gorm:"type:char(64);not null"`
	X     float64 `gorm:"not null"`
	Y     float64 `gorm:"not null"`
	Z     float64 `gorm:"not null"`
	Time  int64   `gorm:"not null"`
	State uint8
}

type sensorDataModel struct{}

var SensorDataModel sensorDataModel

func (sensorDataModel) CreateManay(rows []SensorData) {
	tx := db.DB.Begin()

	for _, row := range rows {
		tx.Create(&row)
	}

	defer func() {
		tx.Commit()
	}()
}

func (sensorDataModel) CreateOne(row *SensorData) {
	db.DB.Create(row)
}
