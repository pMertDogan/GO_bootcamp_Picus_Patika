package basket

import (
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain"
	"go.uber.org/zap"
	"gorm.io/gorm"
	// "gorm.io/gorm/clause"
)

type BasketRepository struct {
	db *gorm.DB
}

//create a sigleton of the repo instance
var singleton *BasketRepository = nil

//initilaze the repo with gorm db
func BasketRepoInit(db *gorm.DB) *BasketRepository {
	if singleton == nil {
		singleton = &BasketRepository{db}
	}
	return singleton
}

//Before using this you need initialize the repo
func Repo() *BasketRepository {
	return singleton
}

//Migrate curent values if exist on current DB
func (c *BasketRepository) Migrations() {
	c.db.AutoMigrate(&Basket{})
	//https://gorm.io/docs/migration.html#content-inner
	//https://gorm.io/docs/migration.html#Auto-Migration
}

//Get All Baskets from DB
func (c *BasketRepository) GetAllBaskets() (Baskets, error) {
	var Baskets Baskets
	var result *gorm.DB
	result = c.db.Find(&Baskets)

	if result.Error != nil {
		return nil, result.Error
	}

	return Baskets, nil

}

//Limit the number of rows to be returned
//We can use above function GetAllBaskets with parameter like limit but we seperated them
func (c *BasketRepository) GetAllBasketsWithLimit(limit int) (Baskets, error) {
	var Baskets Baskets
	var result *gorm.DB
	result = c.db.Find(&Baskets).Limit(limit)

	if result.Error != nil {
		return nil, result.Error
	}

	return Baskets, nil

}

func (c *BasketRepository) GetAllBasketsWithPagination(page, pageSize int) (Baskets, error) {

	var Baskets Baskets
	//resturn paginated data
	result := c.db.Scopes(domain.Paginate(page, pageSize)).Find(&Baskets)

	if result.Error != nil {
		return nil, result.Error
	}

	return Baskets, nil

}

//add items to Basket table
func (c *BasketRepository) CreateOrUpdateBasket(userID, productID, totalQuantity int) error {
	var result *gorm.DB
	//check if product is already in basket
	var basket Basket
	result = c.db.Where("user_id = ? AND product_id = ?", userID, productID).First(&basket)
	//if product is already in basket
	if result.RowsAffected > 0 {
		//update quantity
		result = c.db.Model(&Basket{}).Where("user_id = ? AND product_id = ?", userID, productID).
			Update("total_quantity", basket.TotalQuantity+totalQuantity)
	} else {
		//if product is not in basket
		result = c.db.Create(&Basket{UserID: userID, ProductID: productID, TotalQuantity: totalQuantity})
	}

	if result.Error != nil {
		return result.Error
	}

	return nil

}

//Get baskets by user id include product and user
func (c *BasketRepository) GetBasketsByUserID(userID int) (BasketSToResponseDTO, error) {
	var basket BasketSToResponseDTO
	var result *gorm.DB
	result = c.db.Raw(`
	SELECT 
	baskets.id,
	baskets.created_at,
	baskets.updated_at,
	baskets.deleted_at,
	baskets.user_id,
	baskets.total_quantity,
	baskets.product_id,
	products.sku,
	products.product_name,
	products.description,
	products.color,
	products.price,
	products.stock_count,
	products.category_id,
	products.store_id
FROM baskets
Join products ON products.id = baskets.product_id
WHERE user_id = ?
and baskets.deleted_at is null
`, userID).Scan(&basket)
	// and baskets.deleted_at is null can be added to filter deleted baskets

	zap.L().Debug(result.Statement.SQL.String())
	if result.Error != nil {
		return basket, result.Error
	}

	return basket, nil

}

//Get baskets by user id include product and user
func (c *BasketRepository) GetBasketsByUserIDWithPaginations(userID int, page, pageSize string) (BasketSToResponseDTO, error) {
	var basket BasketSToResponseDTO
	var result *gorm.DB
	
	result = c.db.Raw(`
	SELECT 
	baskets.id,
	baskets.created_at,
	baskets.updated_at,
	baskets.deleted_at,
	baskets.user_id,
	baskets.total_quantity,
	baskets.product_id,
	products.sku,
	products.product_name,
	products.description,
	products.color,
	products.price,
	products.stock_count,
	products.category_id,
	products.store_id
FROM baskets
Join products ON products.id = baskets.product_id
WHERE user_id = ?
and baskets.deleted_at is null
limit ?
offset ?
`, userID, pageSize, page).Scan(&basket)
	// and baskets.deleted_at is null can be added to filter deleted baskets

	zap.L().Debug(result.Statement.SQL.String())

	if result.Error != nil {
		return basket, result.Error
	}

	return basket, nil

}

//get basket by userid and id
func (c *BasketRepository) GetBasketByUserIDAndID(userID, id int) (*Basket, error) {
	var basket Basket
	var result *gorm.DB
	result = c.db.Where("user_id = ? AND id = ?", userID, id).First(&basket)

	if result.Error != nil {
		return nil, result.Error
	}

	return &basket, nil
}

//get baskets by userid and id
func (c *BasketRepository) GetBasketsByUserIDAndBasketIDs(id []int) (Baskets, error) {
	var baskets Baskets
	// SELECT * FROM baskets WHERE id IN (...id);
	result := c.db.Where(id).Joins("Product").Joins("User").Find(&baskets)

	if result.Error != nil {
		return baskets, result.Error
	}

	return baskets, nil

}

//Update basket quantity
func (c *BasketRepository) UpdateBasketQuantity(userID, totalQuantity, id int) error {
	var result *gorm.DB
	result = c.db.Model(&Basket{}).Where("user_id = ? AND id = ?", userID, id).Update("total_quantity", totalQuantity)

	if result.Error != nil {
		return result.Error
	}

	return nil

}

//remove basket by user id
func (c *BasketRepository) RemoveBasketByUserIDBasketID(userID, basketID int) error {
	var result *gorm.DB
	result = c.db.Where("user_id = ? and id = ?", userID, basketID).Delete(&Basket{})

	if result.Error != nil {
		return result.Error
	}

	return nil

}

//get basket by user id and product id
func (c *BasketRepository) GetBasketByUserIDAndProductID(userID, productID int) (*Basket, error) {
	var basket Basket
	var result *gorm.DB
	result = c.db.Where("user_id = ? AND product_id = ?", userID, productID).First(&basket)

	if result.Error != nil {
		return nil, result.Error
	}

	return &basket, nil

}
