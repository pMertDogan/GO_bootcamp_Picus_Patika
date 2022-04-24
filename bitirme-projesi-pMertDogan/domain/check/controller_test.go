package check

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/database"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/pkg/config"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func Test_statusCheck(t *testing.T) {

	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/check/status", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, `service is runnng`, response["message"])
}

func setupRouter() *gin.Engine {
	cfg, err := config.LoadConfig("config-local")
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	database.ConnectPostgresDB(cfg)
	if err != nil {
		zap.L().Fatal("cannot connect to database")
	}

	r := gin.Default()
	CheckControllerDef(r)
	return r
}

func Test_ready(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/check/ready", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var response map[string]interface{}
	json.Unmarshal([]byte(w.Body.String()), &response)
	fmt.Println(response)
	//print keys
	for k := range response {
		fmt.Println(k)
	}
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `service is runnng`, response["message"])

}
