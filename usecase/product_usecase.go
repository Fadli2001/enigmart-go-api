package usecase 

import (
	"enigmart-api/model"
	"enigmart-api/repository"

)

type ProductUseCase interface{
	CreateNewProduct(newProduct *model.Product) model.Product
	GetAllProduct()[]model.Product
	UpdateProduct(newProduct model.Product) model.Product
	DeleteProduct(id string)bool
	GetProductById(id string)model.Product
}

type productUseCase struct{
	repo repository.ProductRepository
}

func (p *productUseCase) CreateNewProduct(newProduct *model.Product) model.Product {
	return p.repo.Insert(newProduct)
} 

func (p *productUseCase) GetAllProduct() []model.Product {
	return p.repo.List()
}

func (p *productUseCase) UpdateProduct(newProduct model.Product) model.Product{
	return p.repo.Update(&newProduct)
}

func (p *productUseCase) DeleteProduct(id string)bool {
	return p.repo.Delete(id)
}

func (p *productUseCase) GetProductById(id string)model.Product{
	return p.repo.Get(id)
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	pc := new(productUseCase)
	pc.repo = repo
	return pc
}