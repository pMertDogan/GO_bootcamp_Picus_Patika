package bookRest

import (
	"log"
	"net/http"

	"github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/domain"
	exitError "github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/domain"
	"github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/domain/book"
)

/*
ITS CASE SENSIITIVE
PS C:\Projeler\homework-4-week-5-pMertDogan> curl -v "http://localhost:8080/book/find?name=Hob"
VERBOSE: GET with 0-byte payload
VERBOSE: received 369-byte response of content type application/json


StatusCode        : 200
StatusDescription : OK
Content           : {"CreatedAt":"2022-03-25T23:46:42.744335+03:00","UpdatedAt":"2022-03-25T23:46:42.744335+03:00","DeletedAt":null,"ID":"1","AuthorID":"0","BookName":"Hobbit","NumberOfPages":665,"Sto
                    ckCount":14,"Price":...
RawContent        : HTTP/1.1 200 OK
                    Content-Length: 369
                    Content-Type: application/json
                    Date: Tue, 29 Mar 2022 17:23:15 GMT

                    {"CreatedAt":"2022-03-25T23:46:42.744335+03:00","UpdatedAt":"2022-03-25T23:46:42.744335+03:...
Forms             : {}
Headers           : {[Content-Length, 369], [Content-Type, application/json], [Date, Tue, 29 Mar 2022 17:23:15 GMT]}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        : System.__ComObject
RawContentLength  : 369



PS C:\Projeler\homework-4-week-5-pMertDogan>
*/
//get book by searched word
func FindBookByNameWithoutAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//https://stackoverflow.com/questions/49461354/go-vet-composite-literal-uses-unkeyed-fields-with-embedded-types
	var responseModel domain.APIStructBook = domain.APIStructBook{
		// Books:    nil,
		// Code:     0,
		// ErrorMsg: "",
	}

	name := r.URL.Query().Get("name")

	log.Println("Find Requested  name: " + name)

	b, err := book.Repo().FindByName(name)
	//check is we can get book without error or not
	if err != nil {
		responseModel.ErrorMsg = "Book is not exist " + err.Error()
		//we can type domain.ExitWithError(&responseModel, w, http.StatusInternalServerError)
		exitError.ExitWithError(&responseModel, w, http.StatusInternalServerError)
		return
		// res ,_ := responseModel.String()
		// http.Error(w, string(res), http.StatusBadRequest)
		// return
		// w.Write([]byte())
	}
	responseModel.Books = []book.Book{*b}
	//convert struct to json
	json, err := responseModel.Marshal()

	//maybe we dont need it ?
	if err != nil {
		//exit with error flow
		responseModel.ErrorMsg = "Internal server error \n" + err.Error()
		exitError.ExitWithError(&responseModel, w, http.StatusInternalServerError)
		return
		//return to avoid below command should not be forgetted ! Maybe we need use another logic
		//alternative we can use if else :)

	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(json))
}
