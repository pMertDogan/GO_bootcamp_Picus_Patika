package category

import (

	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain"
	"gorm.io/gorm"
	// "gorm.io/gorm/clause"
)

type CategoryRepository struct {
	db *gorm.DB
}

//create a sigleton of the repo instance
var singleton *CategoryRepository = nil

//initilaze the repo with gorm db
func CategoryRepoInit(db *gorm.DB) *CategoryRepository {
	if singleton == nil {
		singleton = &CategoryRepository{db}
	}
	return singleton
}

//Before using this you need initialize the repo
func Repo() *CategoryRepository {
	return singleton
}

//Migrate curent values if exist on current DB
func (c *CategoryRepository) Migrations() {
	c.db.AutoMigrate(&Category{})
	//https://gorm.io/docs/migration.html#content-inner
	//https://gorm.io/docs/migration.html#Auto-Migration
}

//Save All Categories to SQL
//Maybe we need optimize this function for larger imports.
//https://gorm.io/docs/create.html#Batch-Insert
// Will help us but we need make it work same as our logic
//create or update
// c.db.Clauses(clause.OnConflict{
// 	UpdateAll: true,
//   }).Create(&categories)
// c.db.FirstOrCreate(&categories)

func (c *CategoryRepository) CreateCategories(categories Categorys) {
	//categoryName is uniq

	for _, v := range categories {
		//https://stackoverflow.com/questions/39333102/how-to-create-or-update-a-record-with-gorm
		//If its not exist just create it else update it
		if c.db.Model(&v).Unscoped().Where("category_name = ?", v.CategoryName).Updates(&v).RowsAffected == 0 {
			c.db.Create(&v)
			//zero means not found
		}
	}

	/*
		Maybe we can use it, doing same thing
		https://gorm.io/docs/advanced_query.html#FirstOrCreate can be used
				// User not found, initialize it with give conditions and Assign attributes
			db.Where(User{Name: "non_existing"}).Assign(User{Age: 20}).FirstOrCreate(&user)
			// SELECT * FROM users WHERE name = 'non_existing' ORDER BY id LIMIT 1;
			// INSERT INTO "users" (name, age) VALUES ("non_existing", 20);
			// user -> User{ID: 112, Name: "non_existing", Age: 20}

			// Found user with `name` = `jinzhu`, update it with Assign attributes
			db.Where(User{Name: "jinzhu"}).Assign(User{Age: 20}).FirstOrCreate(&user)
			// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;
			// UPDATE users SET age=20 WHERE id = 111;
			// user -> User{ID: 111, Name: "jinzhu", Age: 20}

	*/

}

//Get All Categories from DB
func (c *CategoryRepository) GetAllCategories() (Categorys, error) {
	var categories Categorys
	var result *gorm.DB
	result = c.db.Find(&categories)

	if result.Error != nil {
		return nil, result.Error
	}

	return categories, nil

}

//Limit the number of rows to be returned
//We can use above function GetAllCategories with parameter like limit but we seperated them
func (c *CategoryRepository) GetAllCategoriesWithLimit(limit int) (Categorys, error) {
	var categories Categorys
	var result *gorm.DB
	result = c.db.Find(&categories).Limit(limit)

	if result.Error != nil {
		return nil, result.Error
	}

	return categories, nil

}
//Get Categories by ID from DB and return it as Category struct type and error if exist 
func (c *CategoryRepository) GetAllCategoriesWithPagination(page, pageSize int) (Categorys, error) {

	var categories Categorys
	//resturn paginated data
	result := c.db.Scopes(domain.Paginate(page, pageSize)).Find(&categories)

	if result.Error != nil {
		return nil, result.Error
	}

	return categories, nil

}
