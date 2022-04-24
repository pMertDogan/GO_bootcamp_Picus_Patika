package store

import (
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/pkg/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type StoreRepository struct {
	db *gorm.DB
}

//create a sigleton of the repo instance
var singleton *StoreRepository = nil

//initilaze the repo with gorm db
func StoreRepoInit(db *gorm.DB) *StoreRepository {
	if singleton == nil {
		singleton = &StoreRepository{db}
	}
	return singleton
}

//Before using this you need initialize the repo
func Repo() *StoreRepository {
	return singleton
}

//Migrate curent values if exist on current DB
func (c *StoreRepository) Migrations() {
	c.db.AutoMigrate(&Store{})
	//https://gorm.io/docs/migration.html#content-inner
	//https://gorm.io/docs/migration.html#Auto-Migration
}

func (c *StoreRepository) Create(Store Store) error {
	result := c.db.Create(&Store)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

//Create initial store if not exist
func (c *StoreRepository) CreateStoreIfNotExist(cfg *config.Config) error {

	var store Store
	store.Name = cfg.StoreData.Name
	store.Address = cfg.StoreData.Address
	store.Phone = cfg.StoreData.Phone
	store.Email = cfg.StoreData.Email
	store.Description = cfg.StoreData.Description

	//if there is no store is created , create it
	if result := c.db.First(&store); result.RowsAffected == 0 {
		//We can use this one direclty without if check
		c.db.FirstOrCreate(&store, store)
		zap.L().Info("store  created")
	}

	return nil
}
