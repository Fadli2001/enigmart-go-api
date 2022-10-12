package manager

import "enigmart-api/repository"

type RepositoryManager interface {
	ProductRepository() repository.ProductRepository
}

type repositoryManager struct {
	infra InfraManager
}

func (i *repositoryManager) ProductRepository() repository.ProductRepository {
	return repository.NewProductRepository(i.infra.SqlDB())
}

func NewRepositoryManager(infra InfraManager) RepositoryManager {
	return &repositoryManager{
		infra: infra,
	}
}