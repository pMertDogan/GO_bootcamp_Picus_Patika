package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServerConfig ServerConfig
	JWTConfig    JWTConfig
	DBConfig DBConfig
	Logger   Logger
	Admin	Admin
	StoreData	StoreData
}


type ServerConfig struct {
	AppVersion       string
	Mode             string
	RoutePrefix      string
	Debug            bool
	Port             string
	TimeoutSecs      int64
	ReadTimeoutSecs  int64
	WriteTimeoutSecs int64
}

/*
YAML conversion 
DBConfig:
  Host: localhost
  Port: 5432
  Dbname: basketStore
  Username: postgres
  Password: 123456
*/

type DBConfig struct {
	Host     string
	Port     string
	DBName   string
	Username string
	Password string
}

// Logger config
type Logger struct {
	Development bool
	Encoding    string
	Level       string
}

//JWTConfig:
type JWTConfig struct {
	// SessionTime int
	SecretKey   string
	AccesTokenLifeMinute int64
	RefreshTokenLifeMinute int64
}

type Admin struct{
	Password string
	Email string
}

type StoreData struct{
	Name string
	Description string
	Phone string
	Email string
	Address string
}

// LoadConfig file from given path
func LoadConfig(filename string) (*Config, error) {
	//create new viper instance
	v := viper.New()
	//get config file  from given path
	v.SetConfigName(filename) //name of config file (without extension)
	// viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	// v.AddConfigPath(".")
	v.AddConfigPath("./pkg/config/")
	// add for TestGetAllProductWithPagination
	v.AddConfigPath("../../pkg/config/")
	v.AutomaticEnv()

	//read config file
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	//convert it to type safe	config
	//if its not find the struct in config file struct use init values of the type
	// like string => "" , bool => false , int => 0
	//Its not inclueded in error :/
	//We need addtionals checks for each field
	var c Config
	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	
	return &c, nil

}
