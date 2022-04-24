package authorRest

import (
	"fmt"
	"net/http"

	"github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/domain"
)


//maybe we can make it abstract and its work on both classes
func ExitWithError(model *domain.APIStructAuthor, w http.ResponseWriter, code int) {
	//print error to log
	fmt.Println(model.ErrorMsg)
	//set error code to one
	model.Code = 1
	// get error string
	res, _ := model.String()
	http.Error(w, string(res), code)

	// return

}