package product

import (
	"errors"
	"strconv"

)

//mocked repository for with singature of IRepoProduct

type ProductRepoMock struct {
	products Products
}

//create interfaces for mocked repository its duck typed
func (p ProductRepoMock)Create(product Product) error {
	//add to products
	p.products = append(p.products, product)
	return nil
}
func (p ProductRepoMock)GetBySkuWithRelations(sku string) (Product, error){
	for _, product := range p.products {
		if product.Sku == sku {
			return product, nil
		}
	}
	return Product{}, errors.New("product not found")
}

func (p ProductRepoMock)GetAllWithPagination(page, pageSize int) (Products, error){
	return p.products, nil
}

func (p ProductRepoMock)SearchProducts(searchText string, page, pageSize int) (Products, error) {
	//return  prouct if product name,description or sku contains search text
	var products Products
	for _, product := range p.products {
		if product.ProductName == searchText || product.Description == searchText || product.Sku == searchText {
			products = append(products, product)
		}
	}
	return products, nil
}

func (p ProductRepoMock)CreateBulkProduct(products Products){
	p.products = append(p.products, products...)
}
func (p ProductRepoMock)Delete(id string) (Product, error){
	//id to int
	idInt, _ := strconv.Atoi(id)
	for i, product := range p.products {
		if int(product.ID) == idInt {
			p.products = append(p.products[:i], p.products[i+1:]...)
			return product, nil
		}
	}
	return Product{}, errors.New("product not found")
}


func (p ProductRepoMock)GetById(id string) (Product, error){
	//id to int
	idInt, _ := strconv.Atoi(id)
	for _, product := range p.products {
		if int(product.ID) == idInt {
			return product, nil
		}
	}
	return Product{}, errors.New("product not found")
}
func (p ProductRepoMock)Migrations(){
	//do nothing
}
func (p ProductRepoMock)Update(id string, patched Product) (*Product, error){
	//id to int
	idInt, _ := strconv.Atoi(id)
	for i, product := range p.products {
		if int(product.ID) == idInt {
			p.products[i] = patched
			return &p.products[i], nil
		}
	}
	return nil, errors.New("product not found")
}
func (p ProductRepoMock)GetProductQuantityById(id int) (int, error)	{
	//seach for product by id
	for _, product := range p.products {
		if int(product.ID) == id {
			return product.StockCount, nil
		}
	}
	return 0, errors.New("product not found")
}							
