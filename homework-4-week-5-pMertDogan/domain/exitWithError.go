package domain

import (
	"fmt"
	"net/http"
	// "github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/domain"
)

//dont forget call return to exit flow
func ExitWithError(model *APIStructBook, w http.ResponseWriter, code int) {
	//print error to log
	fmt.Println(model.ErrorMsg)
	//set error code to one
	model.Code = 1
	// get error string
	res, _ := model.String()
	http.Error(w, string(res), code)

	// return

}

/*
http://localhost:8080/book?bookID=63
{
"Code": 1,
"ErrorMsg": "Book is not exist record not found",
"Books": null
}
*/

/*
http://localhost:8080/book?bookID=2

{
"Code": 0,
"ErrorMsg": "",
"Books": [
{
"CreatedAt": "2022-03-25T23:46:42.744838+03:00",
"UpdatedAt": "2022-03-25T23:46:42.744838+03:00",
"DeletedAt": null,
"ID": "2",
"AuthorID": "1",
"BookName": "The Unix Programming Environment",
"NumberOfPages": 375,
"StockCount": 55,
"Price": 11,
"Isbn": "ISBN1",
"StockCode": "3456",
"Author": {
"ID": 0,
"CreatedAt": "0001-01-01T00:00:00Z",
"UpdatedAt": "0001-01-01T00:00:00Z",
"DeletedAt": null,
"AuthorID": "",
"Name": ""
}
}
]
}


*/
