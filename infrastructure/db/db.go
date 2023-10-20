package db

import (
	"log"
	"path/filepath"
	"runtime"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/epessine/go-products-server/domain/model"
)

var basepath string

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath = filepath.Dir(b)
}

func Connect(env string) *gorm.DB {
	var db *gorm.DB
	var err error
	var path string

	if env != "test" {
		path = basepath + "/database.sqlite"
	} else {
		path = "file::memory:?cache=shared"
	}

	db, err = gorm.Open(sqlite.Open(path), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}

	db.AutoMigrate(&model.Product{})

	return db
}
