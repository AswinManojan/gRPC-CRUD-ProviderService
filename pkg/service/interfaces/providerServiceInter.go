package interfaces

import model "github.com/grpc/gobus/service/pkg/Model"

type ProviderServiceInter interface {
	ProviderLogin(providerLogin *model.ProviderLogin) (*model.Provider, error)
	FindProviderById(id uint) (*model.Provider, error)
	EditProvider(provider *model.Provider) (*model.Provider, error)
	DeleteProviderById(id uint) (*model.Provider, error)
	CreateProvider(provider *model.Provider) (*model.Provider, error)
}
