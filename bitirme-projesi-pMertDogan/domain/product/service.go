package product

import "go.uber.org/zap"

//create a sigleton of the repo instance
var singleton *IRepoProduct


//initilaze the repo with gorm db
func ProductRepoInit(repo IRepoProduct) IRepoProduct {

if repo == nil {
	zap.L().Fatal("repo cannot be nil")
}

	if singleton == nil {
		singleton = &repo
	}
	return *singleton
}

//Before using this you need initialize the repo
func Repo() IRepoProduct {
	return *singleton
}

// //struct to store repo not used on this case
// type IRepository struct{
// 	repository IRepoProduct
// }

//abstact the repositories
//with the help of this we can create a mock repo for testing
//duck typing
type IRepoProduct interface {
	Create(product Product) error
	GetBySkuWithRelations(sku string) (Product, error)
	GetAllWithPagination(page, pageSize int) (Products, error)
	SearchProducts(searchText string, page, pageSize int) (Products, error)
	CreateBulkProduct(products Products)
	Delete(id string) (Product, error)
	GetById(id string) (Product, error)
	Migrations()
	Update(id string, patched Product) (*Product, error)
	GetProductQuantityById(id int) (int, error)
}
