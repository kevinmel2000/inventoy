package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type ModelModule struct{}

type Data struct {
	CreateTime time.Time `gorm:"null" json:"create_time"`
	CreateBy   string    `gorm:"null" json:"create_by"`
	UpdateTime time.Time `gorm:"null" json:"update_time"`
	UpdateBy   string    `gorm:"null" json:"update_by"`
}

func NewData() Data {
	return Data{
		CreateTime: time.Now(),
		CreateBy:   "",
	}
}

func initDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./inventory.db")
	db.LogMode(true)
	if err != nil {
		fmt.Errorf("Could not open db: %v", err)
	}

	if !db.HasTable(&Item{}) {
		db.CreateTable(&Item{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Item{})
	}

	if !db.HasTable(&InboundItem{}) {
		db.CreateTable(&InboundItem{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&InboundItem{})
	}

	if !db.HasTable(&OutboundItem{}) {
		db.CreateTable(&OutboundItem{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&OutboundItem{})
	}

	return db
}
