package repo

import (
	"log"

	model "github.com/grpc/gobus/service/pkg/Model"
	interfaces "github.com/grpc/gobus/service/pkg/repo/Interfaces"
	"gorm.io/gorm"
)

type ProviderRepoImpl struct {
	db *gorm.DB
}

// CreateProvider implements interfaces.ProviderRepoInter.
func (pr *ProviderRepoImpl) CreateProvider(provider *model.Provider) (*model.Provider, error) {
	result := pr.db.Create(provider)
	if result.Error != nil {
		log.Print("Error creating provider - ProviderRepo")
		return nil, result.Error
	}
	return provider, nil
}

// DeleteProviderById implements interfaces.ProviderRepoInter.
func (pr *ProviderRepoImpl) DeleteProviderById(id uint) (*model.Provider, error) {
	var provider *model.Provider
	result := pr.db.Where("id=?", id).Delete(&provider)
	if result.Error != nil {
		log.Print("Error Deleting the provider - ProviderRepo")
		return nil, result.Error
	}
	return provider, nil
}

// EditProvider implements interfaces.ProviderRepoInter.
func (pr *ProviderRepoImpl) EditProvider(provider *model.Provider) (*model.Provider, error) {
	var pvdr model.Provider
	result := pr.db.Where("username=?", provider.Username).First(&pvdr)
	if result.Error != nil {
		log.Print("Error finding the provider")
		return nil, result.Error
	}
	pvdr.Password = provider.Password
	pvdr.Email = provider.Email
	if err := pr.db.Save(&pvdr).Error; err != nil {
		log.Print("Error saving/updating the provider")
		return nil, err
	}
	return &pvdr, nil
}

// FindProviderById implements interfaces.ProviderRepoInter.
func (pr *ProviderRepoImpl) FindProviderById(id uint) (*model.Provider, error) {
	var pvdr model.Provider
	result := pr.db.Where("id=?", id).Find(&pvdr)
	if result.Error != nil {
		log.Print("Error finding the provider by id")
		return nil, result.Error
	}
	return &pvdr, nil
}

// FindProviderByName implements interfaces.ProviderRepoInter.
func (pr *ProviderRepoImpl) FindProviderByName(name string) (*model.Provider, error) {
	var pvdr model.Provider
	result := pr.db.Where("username=?", name).Find(&pvdr)
	if result.Error != nil {
		log.Print("Error finding the user by provider")
		return nil, result.Error
	}
	return &pvdr, nil
}

func NewProviderRepoImpl(DB *gorm.DB) interfaces.ProviderRepoInter {
	return &ProviderRepoImpl{
		db: DB,
	}
}
