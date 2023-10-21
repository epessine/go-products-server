package factory

import (
	"gorm.io/gorm"

	"github.com/epessine/go-products-server/application/usecase"
	"github.com/epessine/go-products-server/infrastructure/repository"
)

func ProductUseCase(database *gorm.DB) usecase.Product {
	repository := repository.Product{DB: database}
	usecase := usecase.Product{Repository: repository}

	return usecase
}