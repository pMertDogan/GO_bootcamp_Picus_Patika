package product

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
)




func TestGetAllProductWithPagination(t *testing.T) {
	router := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/product/search?search=Mi", nil)
	router.ServeHTTP(w, req)
	//create struct for test
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Mi")
}




