	//go:build wireinject
	// +build wireinject

	package di

	import (
		"github.com/google/wire"
		http "github.com/SethukumarJ/go-gin-clean-arch/pkg/api"
		handler "github.com/SethukumarJ/go-gin-clean-arch/pkg/api/handler"
		config "github.com/SethukumarJ/go-gin-clean-arch/pkg/config"
		db "github.com/SethukumarJ/go-gin-clean-arch/pkg/db"
		repository "github.com/SethukumarJ/go-gin-clean-arch/pkg/repository"
		usecase "github.com/SethukumarJ/go-gin-clean-arch/pkg/usecase"
	)

	func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
		wire.Build(
			db.ConnectDatabase,
		    repository.NewUserMongoRepository, // added provider function
			usecase.NewUserUseCase,
			handler.NewUserHandler,
			http.NewServerHTTP,
		)
		
		return &http.ServerHTTP{}, nil
	}
