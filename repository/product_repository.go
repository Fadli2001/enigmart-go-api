package repository


import(
	"enigmart-api/model"

)


type ProductRepository interface {
	Insert(newProduct *model.Product) model.Product
	List() []model.Product
	Update(newProduct *model.Product) model.Product
	Get(id string)model.Product
	Delete(id string)bool
}

type productRepository struct {
	db []model.Product
}

func (p *productRepository) Insert(newProduct *model.Product)model.Product{
	p.db = append(p.db, *newProduct)
	return *newProduct
}

func (p *productRepository) List() []model.Product {
	return p.db
}

func (p *productRepository) Update(newProduct *model.Product)model.Product{
	var productRes model.Product
	for i:=0 ; i < len(p.db); i++ {
		if p.db[i].Id == newProduct.Id {
			p.db[i] = *newProduct
			productRes = p.db[i]
		}
	}
	return productRes
}

func (p *productRepository) Get(id string) model.Product {
	var product model.Product
	for _, item := range p.db {
		if item.Id == id {
			product = item
		}
	}
	return product
}

func (p *productRepository) Delete(id string)bool {
	var products []model.Product
	result := false
	for i := 0; i < len(p.db); i++ {
		if p.db[i].Id == id {
			products = append(p.db[:i], p.db[i+1:]...)
			result = true
		}
	}
	p.db = products
	return result
}

func NewProductRepository() ProductRepository {
	repo := new(productRepository)
	return repo
}