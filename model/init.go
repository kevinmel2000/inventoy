package model

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type ModelModule struct{}

type Data struct {
	CreatedAt time.Time `gorm:"null" json:"created_at"`
	CreateBy  string    `gorm:"null" json:"created_by"`
	UpdatedAt time.Time `gorm:"null" json:"updated_at"`
	UpdateBy  string    `gorm:"null" json:"updated_by"`
}

func NewData() Data {
	return Data{
		CreatedAt: time.Now(),
		CreateBy:  "",
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

	if !db.HasTable(&StockBatch{}) {
		db.CreateTable(&StockBatch{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&StockBatch{})
	}

	if !db.HasTable(&OrderItem{}) {
		db.CreateTable(&OrderItem{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&OrderItem{})
	}

	return db
}

func initSQL() *sql.DB {
	db, _ := sql.Open("sqlite3", "./inventory.db")
	return db
}
