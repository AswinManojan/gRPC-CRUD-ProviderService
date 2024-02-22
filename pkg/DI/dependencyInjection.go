package di

import (
	"github.com/grpc/gobus/service/config"
	db "github.com/grpc/gobus/service/pkg/DB"
	handler "github.com/grpc/gobus/service/pkg/Handler"
	"github.com/grpc/gobus/service/pkg/repo"
	"github.com/grpc/gobus/service/pkg/server"
	"github.com/grpc/gobus/service/pkg/service"
)

func Init() {
	config := config.LoadConfig()
	// fmt.Println(config)
	Db := db.Connect_DB(config)
	providerRepo := repo.NewProviderRepoImpl(Db)
	providerService := service.NewProviderServiceImpl(providerRepo)
	providerHandler := handler.NewProviderHandler(providerService)
	server.NewGRPCServer(config, providerHandler)
}
