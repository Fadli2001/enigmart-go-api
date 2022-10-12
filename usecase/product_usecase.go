package usecase

import (
	"enigmart-api/model"
	"enigmart-api/repository"
	"enigmart-api/utils"
)

type ProductUseCase interface{
	CreateNewProduct(newProduct *model.Product) error
	GetAllProduct(page int,totalRows int)([]model.Product,error)
	UpdateProduct(newProduct model.Product) error
	DeleteProduct(id string)error
	GetProductById(id string)(model.Product,error)
}

type productUseCase struct{
	repo repository.ProductRepository
}

func (p *productUseCase) CreateNewProduct(newProduct *model.Product) error {
	newProduct.Id = utils.GenerateId()
	return p.repo.Insert(newProduct)
} 

func (p *productUseCase) GetProductById(id string)(model.Product,error){
	return p.repo.Get(id)
}
func (p *productUseCase) GetAllProduct(page int,totalRows int) ([]model.Product,error) {
	return p.repo.List(page,totalRows )
}

func (p *productUseCase) UpdateProduct(newProduct model.Product) error{
	return p.repo.Update(&newProduct)
}

func (p *productUseCase) DeleteProduct(id string)error {
	return p.repo.Delete(id)
}


func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	pc := new(productUseCase)
	pc.repo = repo
	return pc
}