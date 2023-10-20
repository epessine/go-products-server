package main

import (
	"os"

	"gorm.io/gorm"

	"github.com/epessine/go-products-server/application/grpc"
	"github.com/epessine/go-products-server/infrastructure/db"
)

var database *gorm.DB

func main() {
	database = db.Connect(os.Getenv("ENV"))
	grpc.Start(database, 50051)
}
