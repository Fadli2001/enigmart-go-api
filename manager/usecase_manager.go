package manager

import "enigmart-api/usecase"



type UseCaseManager interface {
	ProductUseCase() usecase.ProductUseCase
}

type useCaseManager struct {
	repo RepositoryManager
}

func (u *useCaseManager) ProductUseCase() usecase.ProductUseCase {
	return usecase.NewProductUseCase(u.repo.ProductRepository())
}

func NewUseCaseManager(repo RepositoryManager) UseCaseManager {
	return &useCaseManager{
		repo : repo,
	}
}