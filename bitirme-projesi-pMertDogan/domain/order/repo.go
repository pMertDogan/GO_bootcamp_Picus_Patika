package order

import (

	"encoding/json"

	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/basket"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/product"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrderRepository struct {
	db *gorm.DB
}

//create a sigleton of the repo instance
var singleton *OrderRepository = nil

//initilaze the repo with gorm db
func OrderRepoInit(db *gorm.DB) *OrderRepository {
	if singleton == nil {
		singleton = &OrderRepository{db}
	}
	return singleton
}

//Before using this you need initialize the repo
func Repo() *OrderRepository {
	return singleton
}

//Migrate curent values if exist on current DB
func (c *OrderRepository) Migrations() {
	c.db.AutoMigrate(&Order{})
	//https://gorm.io/docs/migration.html#content-inner
	//https://gorm.io/docs/migration.html#Auto-Migration
}

/*
	1.Open Transaction and lock product rows
	2.Check if current quantity of each product is enough for order
	3.If not enough quantity then rollback transaction and return error
	4.If enough quantity then update product quantity
	5.Update add order to orders table
	6.Delete completed basket items
	7.Commit transaction
*/
func (c *OrderRepository) CompleteOrder(baskets basket.Baskets, comment, shipingAddress, billingAddress string) error {

	//create producIDArray From basket
	productIdQuantityMap := baskets.GenerateProductIDTotalQuantityMap()

	//create productQuantityMap keys as array

	productIDs := make([]int, len(productIdQuantityMap))

	//this one create a slice of productIDs from keys of the map
	i := 0
	for k := range productIdQuantityMap {
		productIDs[i] = k
		i++
	}
	zap.L().Info("productIDs", zap.Any("productIDs", productIDs))

	//https://www.postgresql.org/docs/9.4/explicit-locking.html
	//lock product rows using productIDs array
	//With this lock we can avoid race conditions , false updates
	result := c.db.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id in ?", productIDs).Find(&product.Products{})

	if result.Error != nil {
		return errors.Wrap(result.Error, "Error while locking product rows")
	}

	//Open Transaction
	//If one of them fail then rollback transaction
	//So we secure our data with transaction
	//If all operations are valid we will commit it and unlock product rows
	tx := c.db.Begin()
	var buyedProducts product.Products
	//check if current quantity of each product is enough for order
	for productID, quantity := range productIdQuantityMap {
		var product product.Product
		result = tx.Where("products.id = ?", productID).Joins("Category").Joins("Store").First(&product)

		if result.Error != nil {
			tx.Rollback()
			return errors.Wrap(result.Error, "Error while finding product")
		}
		if product.StockCount < quantity {
			//we can not complete order
			tx.Rollback()
			return errors.New("Not enough quantity")
		} else {
			//update product quantity and save it.If error then rollback transaction will help us on this update
			product.StockCount = product.StockCount - quantity
			//we add them to our buyedProducts array
			buyedProducts = append(buyedProducts, product)
			tx.Save(&product)
		}
	}

	var orders Orders
	var order Order

	order.UserID = baskets[0].UserID
	order.Comment = comment
	order.ShippingAdress = shipingAddress
	order.BillingAddress = billingAddress
	// for each buyedProducts we add order to orders table
	for _, product := range buyedProducts {
		//generate product snapshot from product
		v, err := json.Marshal(product)
		if err != nil {
			tx.Rollback()
			return errors.Wrap(err, "Error while marshalling product")
		}
		order.ProductSnapshot = v
		order.ProductID = int(product.ID)
		order.Quantity = productIdQuantityMap[int(product.ID)]
		orders = append(orders, order)
		//Add quantity information to order
	}

	//add orders to orders table
	tx.Save(&orders)

	//delete completed basket items
	tx.Delete(&baskets)

	//end transaction
	return tx.Commit().Error
}

//get orders of customer with pagination
func (c *OrderRepository) GetOrders(userID int, page, pageSize int) (Orders, error) {
	var orders Orders
	result := c.db.Where("user_id = ?", userID).
		Scopes(domain.Paginate(page, pageSize)).
		Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}

/*
Get user check is user has order with requested order id
*/
func (c *OrderRepository) HasOrder(userID, orderID int) (*Order,error) {
	var order Order
	result := c.db.Where("user_id = ? AND id = ?", userID, orderID).First(&order)
	if result.Error != nil {
		return nil,result.Error
	}
	return &order,nil
}

/*
Cancel order if order created time is not older than 14 days.
All operations will be done in transaction
After cancel order we will delete order from orders table
We will add stock count to product table
*/
func (c *OrderRepository) CancelOrder(order Order) error {

	//start transaction
	tx := c.db.Begin()

	// lock product rows
	result := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", order.ProductID).Find(&product.Products{})
	if result.Error != nil {
		tx.Rollback()
		return  errors.Wrap(result.Error, "Error while locking product rows")
	}

	//lock order rows
	result = tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", order.ID).Find(&Orders{})
	if result.Error != nil {
		tx.Rollback()
		return errors.Wrap(result.Error, "Error while locking order rows")
	}

	
	//delete order
	result = tx.Delete(&order)

	if result.Error != nil {
		return  result.Error
	}
	zap.L().Debug("order deleted", zap.Any("order", order))

	//add deleted order quantity to product table
	//convert order.Quantity to string
	// quantity := strconv.Itoa(order.Quantity)
	//convert order.ProductID to string
	// productID := strconv.Itoa(order.ProductID)

	// result =	tx.Raw("UPDATE products SET stock_count = stock_count + " + quantity+ " WHERE id = "+ productID)
	// Raw SQL {"SQL : ": "UPDATE products SET stock_count = stock_count + $1 WHERE id = $2"}

	result = tx.Exec("UPDATE products SET stock_count = stock_count + ? WHERE id = ?", order.Quantity, order.ProductID)

	zap.L().Debug("Raw SQL", zap.Any("SQL : ", 	result.Statement.SQL.String()))
	if result.Error != nil {
		return  result.Error
	}
	zap.L().Debug("product stock count updated")


	//end transaction
	return tx.Commit().Error


}
