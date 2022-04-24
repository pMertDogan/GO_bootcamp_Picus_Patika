package product

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {

	router := SetupRouter()
	w := httptest.NewRecorder()

	payload := strings.NewReader(`
    "sku" : "t10",
    "productName" : "Xaiomi Mi8",
    "description" : "Xiaomi Mi 8 is new flagship",
    "color" : "Black",
    "price" : "3000",
    "stockCount" : "50",
    "categoryID" : "76",
    "storeID" : "1"`)

	req, _ := http.NewRequest("POST", "/product/", payload)
	req.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQGxvY2FsaG9zdC5jb20iLCJleHAiOjE2NDk0NDk3NDgsImlhdCI6MTY0OTQ0OTE0OCwiaXNBZG1pbiI6dHJ1ZSwiaXNJdEFjY2VzVG9rZW4iOnRydWUsInVzZXJJZCI6MX0.VAap49WE3JejrJqHjX3aECLg0pV8-COE2HAO77jZxp8")

	router.ServeHTTP(w, req)
	// fmt.Println(w.Body.String())
	// assert.Equal(t, 401, w.Code)
	// assert.Contains(t, w.Body.String(), "Please provide token")

	assert.Equal(t, 403, w.Code)
	assert.Contains(t, w.Body.String(), "Token is either expired or not active yet")

}
