package product

import (
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain"
	"gorm.io/gorm"
)

//PostreSQL repo
type GormProductRepo struct {
	DB *gorm.DB
}

//Migrate curent values if exist on current DB
func (c *GormProductRepo) Migrations() {
	c.DB.AutoMigrate(&Product{})
	//https://gorm.io/docs/migration.html#content-inner
	//https://gorm.io/docs/migration.html#Auto-Migration
}

//Create single product
func (c *GormProductRepo) Create(product Product) error {
	result := c.DB.Create(&product)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

//Get single product by sku with relations
func (c *GormProductRepo) GetBySkuWithRelations(sku string) (Product, error) {
	var product Product
	//get product by sku with relations
	result := c.DB.Joins("Store").Joins("Category").Where("sku = ?", sku).First(&product)

	if result.Error != nil {
		return product, result.Error
	}

	return product, nil
}

//return all products with relations
func (c *GormProductRepo) GetAllWithPagination(page, pageSize int) (Products, error) {

	var products Products
	// product := Product
	//resturn paginated data
	result := c.DB.Scopes(domain.Paginate(page, pageSize)).Joins("Store").Joins("Category").Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil

}

//return all products with relations
func (c *GormProductRepo) SearchProducts(searchText string, page, pageSize int) (Products, error) {

	//https://www.compose.com/articles/mastering-postgresql-tools-full-text-search-and-phrase-search/
	var products Products

	//ILIKE is case insensitive
	result := c.DB.
		Where("product_name ILIKE ?", "%"+searchText+"%").
		Or("products.description ILIKE ?", "%"+searchText+"%").
		Or("color ILIKE ?", "%"+searchText+"%").
		Or("sku ILIKE ?", "%"+searchText+"%").
		//hardcoded store name search :/
		Or(" \"Store\".\"name\" ILIKE ?", "%"+searchText+"%").
		Or("\"Category\".\"category_name\" ILIKE ?", "%"+searchText+"%").
		Scopes(domain.Paginate(page, pageSize)).
		Joins("Store").Joins("Category").Find(&products).Limit(10)

	// result := c.db.Raw("select * from products  	Join categories ON categories.id = products.category_id Join stores ON stores.id = products.store_id where sku ILIKE ? 	or product_name ILIKE ? 	or products.description ILIKE ?", "%"+searchText+"%", "%"+searchText+"%", "%"+searchText+"%").Scan(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil

}

//create product
func (c *GormProductRepo) CreateBulkProduct(products Products) {
	//categoryName is uniq

	for _, v := range products {
		//https://stackoverflow.com/questions/39333102/how-to-create-or-update-a-record-with-gorm
		//If its not exist just create it else update it
		//SKU is uniq
		//with the help of the unsoped we can detect soft deleted products if we cant detect its try to create it
		if c.DB.Model(&v).Unscoped().Where("sku = ?", v.Sku).Updates(&v).RowsAffected == 0 {
			//zero means not found
			c.DB.Create(&v)
		}
	}

}

//delete product by id
func (c *GormProductRepo) Delete(id string) (Product, error) {
	var product Product
	result := c.DB.Where("id = ?", id).First(&product)

	if result.Error != nil {
		return product, result.Error
	}

	result = c.DB.Delete(&product)

	if result.Error != nil {
		return product, result.Error
	}

	return product, nil
}

//PATCH Product
func (c *GormProductRepo) Update(id string, patched Product) (*Product, error) {

	var old Product
	//get the first product if exist
	result := c.DB.Where("id = ?", id).First(&old)

	if result.Error != nil {
		return nil, result.Error
	}
	//patch the produc

	result = c.DB.Model(&old).Updates(patched).Where("id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &old, nil
}

//Get single product by id
func (c *GormProductRepo) GetById(id string) (Product, error) {
	var product Product
	result := c.DB.Where("id = ?", id).First(&product)

	if result.Error != nil {
		return product, result.Error
	}

	return product, nil
}

//return product quantity by id
func (c *GormProductRepo) GetProductQuantityById(id int) (int, error) {
	var product Product
	result := c.DB.Where("id = ?", id).First(&product)

	if result.Error != nil {
		return 0, result.Error
	}

	return product.StockCount, nil
}
