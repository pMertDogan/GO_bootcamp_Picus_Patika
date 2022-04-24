package bookRest

import (
	"log"
	"net/http"

	"github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/domain"
	"github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/domain/book"
)

//curl -v "http://localhost:8080/book?bookID=1"
//get book by id
func GetBookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var responseModel domain.APIStructBook = domain.APIStructBook{}

	bookdID := r.URL.Query().Get("bookID")
	log.Println("bookID: " + bookdID)

	//exit flow
	if bookdID == "" {
		responseModel.ErrorMsg = "bookID is required"
		domain.ExitWithError(&responseModel, w, http.StatusBadRequest)
		return
	}

	//get book
	b, err := book.Repo().GetByID(bookdID)
	// b := book.Repo().GetByIdWithAuthorName(bookdID)
	// fmt.Println(b)
	if err != nil {
		responseModel.ErrorMsg = "Book is not exist " + err.Error()
		domain.ExitWithError(&responseModel, w, http.StatusBadRequest)
		return
	}

	responseModel.Books = []book.Book{*b}
	//convert struct to json
	json, err := responseModel.Marshal()

	//maybe we dont need it ?
	if err != nil {
		//exit with error flow
		responseModel.ErrorMsg = "Marshall failed \n" + err.Error()
		domain.ExitWithError(&responseModel, w, http.StatusInternalServerError)
		return
		//return to avoid below command should not be forgetted ! Maybe we need use another logic
		//alternative we can use if else :)

	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(json))
}

/*

PS C:\Projeler\homework-4-week-5-pMertDogan> curl -v "http://localhost:8080/book?bookID=1"
VERBOSE: GET with 0-byte payload
VERBOSE: received 369-byte response of content type application/json


StatusCode        : 200
StatusDescription : OK
Content           : {"CreatedAt":"2022-03-25T23:46:42.744335+03:00","UpdatedAt":"2022-03-25T23:46:42.744335+03:00","DeletedAt":null,"ID":"1","AuthorID":"0","BookName":"Hobbit","NumberOfPages":665,"Sto
                    ckCount":14,"Price":...
RawContent        : HTTP/1.1 200 OK
                    Content-Length: 369
                    Content-Type: application/json
                    Date: Tue, 29 Mar 2022 17:06:58 GMT

                    {"CreatedAt":"2022-03-25T23:46:42.744335+03:00","UpdatedAt":"2022-03-25T23:46:42.744335+03:...
Forms             : {}
Headers           : {[Content-Length, 369], [Content-Type, application/json], [Date, Tue, 29 Mar 2022 17:06:58 GMT]}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        : System.__ComObject
RawContentLength  : 369

*/
