package service

import (
	"errors"
	"log"

	model "github.com/grpc/gobus/service/pkg/Model"
	repointerfaces "github.com/grpc/gobus/service/pkg/repo/Interfaces"
	"github.com/grpc/gobus/service/pkg/service/interfaces"
)

type ProviderServiceImpl struct {
	repo repointerfaces.ProviderRepoInter
}

// AddProvider implements interfaces.ProviderServiceInter.
func (ps *ProviderServiceImpl) CreateProvider(provider *model.Provider) (*model.Provider, error) {
	result, err := ps.repo.CreateProvider(provider)
	if err != nil {
		log.Print("Error while creating provider in Service")
		return nil, err
	}
	return result, nil
}

// DeleteProviderByID implements interfaces.ProviderServiceInter.
func (ps *ProviderServiceImpl) DeleteProviderById(id uint) (*model.Provider, error) {
	result, err := ps.repo.DeleteProviderById(id)
	if err != nil {
		log.Print("Error while deleting provider in Service")
		return nil, err
	}
	return result, nil
}

// EditProviderInfo implements interfaces.ProviderServiceInter.
func (ps *ProviderServiceImpl) EditProvider(provider *model.Provider) (*model.Provider, error) {
	result, err := ps.repo.EditProvider(provider)
	if err != nil {
		log.Print("Error while editing provider in Service")
		return nil, err
	}
	return result, nil
}

// ProviderLogin implements interfaces.ProviderServiceInter.
func (ps *ProviderServiceImpl) ProviderLogin(providerLogin *model.ProviderLogin) (*model.Provider, error) {
	provider, err := ps.repo.FindProviderByName(providerLogin.Username)
	if err != nil {
		log.Print("Error while finding provider in Service")
		return nil, err
	}
	if provider.Password != providerLogin.Password {
		log.Print("Invalid login credentials")
		return nil, errors.New("password Mismatch")
	}
	return provider, nil
}

// ViewProviderByID implements interfaces.ProviderServiceInter.
func (ps *ProviderServiceImpl) FindProviderById(id uint) (*model.Provider, error) {
	result, err := ps.repo.FindProviderById(id)
	if err != nil {
		log.Print("Error while finding provider in Service")
		return nil, err
	}
	return result, nil
}

func NewProviderServiceImpl(REPO repointerfaces.ProviderRepoInter) interfaces.ProviderServiceInter {
	return &ProviderServiceImpl{
		repo: REPO,
	}
}
