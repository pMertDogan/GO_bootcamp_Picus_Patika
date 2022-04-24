package logger

import (
	"fmt"

	//"github.com/gin-gonic/gin"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger creates a new logger with the given log level
func NewLogger(config *config.Config) {

	logLevel, err := zapcore.ParseLevel(config.Logger.Level)
	if err != nil {
		panic(fmt.Sprintf("Unknown log level %v", logLevel))
	}


	var cfg zap.Config
	if config.Logger.Development {
		//create config instance for development mode
		cfg = zap.NewDevelopmentConfig()
		//make it colorful
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		//create config instance for production mode
		cfg = zap.NewProductionConfig()
	}
	//build our config with given options and log level
	logger, err := cfg.Build()
	if err != nil {
		logger = zap.NewNop()
	}
	//set the logger to the global logger
	zap.ReplaceGlobals(logger)
}

// Close closes the logger 
func Close() {
	// flushes buffer, if any, and closes the underlying writer.
	defer zap.L().Sync()
	//zap.L is our own global logger
	// => L returns the global Logger, which can be reconfigured with ReplaceGlobals. It's safe for concurrent use.
}

