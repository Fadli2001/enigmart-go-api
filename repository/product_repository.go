package repository

import (
	"enigmart-api/model"
	"enigmart-api/utils"

	"github.com/jmoiron/sqlx"
)


type ProductRepository interface {
	Insert(newProduct *model.Product) error
	List(page int,totalRows int) ([]model.Product,error)
	Update(newProduct *model.Product) error
	Get(id string)(model.Product,error)
	Delete(id string)error
}

type productRepository struct {
	db *sqlx.DB
}

func (p *productRepository) Insert(newProduct *model.Product)error{
	_,err := p.db.NamedExec(utils.INSERT_PRODUCT,newProduct)
	if err != nil {
		return err
	}
	return nil
}

func (p *productRepository) Get(id string)(model.Product,error) {
	var product model.Product
	err := p.db.Get(&product,utils.SELECT_PRODUCT_ID,id)
	if err != nil {
		return model.Product{},err
	}
	return product,nil
}

func (p *productRepository) List(page int,totalRows int) ([]model.Product,error) {
	limit := totalRows
	offset := limit * (page - 1)
	var products []model.Product
	err := p.db.Select(&products,utils.SELECT_PRODUCT_LIST,limit,offset)	
	if err != nil {
		return nil,err
	}	
	return products,nil
}

func (p *productRepository) Update(product *model.Product) error{
	_,err := p.db.NamedExec(utils.UPDATE_PRODUCT,product)
	if err != nil {
		return err
	}
	return nil	
}

func (p *productRepository) Delete(id string)error {
	_,err := p.db.Exec(utils.DELETE_PRODUCT,id)
	if err != nil {
		return err
	}
	return nil
}

func NewProductRepository(db *sqlx.DB) ProductRepository {
	repo := new(productRepository)
	repo.db = db
	return repo
}

/**
 * git init
 * git add .
 * git commit -m 'initial commit'
 * git branch 1-ConnectDB
 * git checkout 1-ConnectDb
 */