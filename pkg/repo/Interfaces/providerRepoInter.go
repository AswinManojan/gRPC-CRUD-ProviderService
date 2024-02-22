package interfaces

import model "github.com/grpc/gobus/service/pkg/Model"

type ProviderRepoInter interface {
	CreateProvider(provider *model.Provider) (*model.Provider, error)
	EditProvider(provider *model.Provider) (*model.Provider, error)
	DeleteProviderById(id uint) (*model.Provider, error)
	FindProviderById(id uint) (*model.Provider, error)
	FindProviderByName(name string) (*model.Provider, error)
}
