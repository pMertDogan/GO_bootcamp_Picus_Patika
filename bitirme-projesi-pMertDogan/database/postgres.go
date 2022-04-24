package database

import (
	"fmt"

	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/pkg/config"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gormDB *gorm.DB

func ConnectPostgresDB(cfg *config.Config) *gorm.DB {
	//get database connection string from env
	dataSourceName := fmt.Sprintf("host=%s port=%s  dbname=%s user=%s sslmode=disable password=%s",
		cfg.DBConfig.Host,
		cfg.DBConfig.Port,
		cfg.DBConfig.DBName,
		cfg.DBConfig.Username,
		cfg.DBConfig.Password,


	//Same but we used struct to get them from config file
	// viper.GetString("Dtabase.Host"),
	// viper.GetString("Database.Port"),
	// viper.GetString("Database.Dbname"),
	// viper.GetString("Database.Username"),
	// viper.GetString("Database.Password"),
	)

	// fmt.Println("Host is " + cfg.DBConfig.Host)
	// fmt.Println(viper.GetString("database.host"))
	// fmt.Println(viper.GetString("database.port"))
	// fmt.Println(viper.GetString("database.username"))
	// fmt.Println(viper.GetString("database.password"))

	//connect to DB
	gormDB2, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		zap.L().Fatal("cannot open database", zap.Error(err))

	}
	//Maybe there is another way to get same reference to gormDB instance
	//used to share it with funcions like statusCheck
	gormDB = gormDB2

	//get DB from gorm.DB
	sqlDB, err := gormDB.DB()
	if err != nil {
		zap.L().Fatal("cannot get database", zap.Error(err))
	}

	//Check is connectio available
	if err := sqlDB.Ping(); err != nil {
		zap.L().Fatal("cannot ping database", zap.Error(err))
	}

	return gormDB
}

//Status Checker for DB
func StatusCheck() (int, error) {
	sqlDB, err := gormDB.DB()
	if err != nil {
		return 0, err
	}
	//Check is connectio available
	if err := sqlDB.Ping(); err != nil {
		return 0, err
	}

	fmt.Println(sqlDB.Stats())

	return sqlDB.Stats().OpenConnections, nil
}
