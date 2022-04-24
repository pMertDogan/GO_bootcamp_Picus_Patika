package bookRest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/domain/book"
)

type BookAndAuthorAPIStruct struct {
	Code     int             
	ErrorMSG string          
	Book     book.BookAndAuthor 
}

func (r *BookAndAuthorAPIStruct) String() ([]byte, error) {
	st, err :=  json.Marshal(r)
	return st,err
}

func (r *BookAndAuthorAPIStruct) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

//curl -v "http://localhost:8080/bookWithAuthor?bookID=1"
//get book by id
func GetBookByIdIncludeAuthor(w http.ResponseWriter, r *http.Request) {

	var bookAndAuthorAPI BookAndAuthorAPIStruct = BookAndAuthorAPIStruct{}
	w.Header().Set("Content-Type", "application/json")

	bookdID := r.URL.Query().Get("bookID")
	log.Println("bookID: " + bookdID)

	if bookdID == "" {
		bookAndAuthorAPI.ErrorMSG = "bookID is required"
		exitWithError(bookAndAuthorAPI, w)
		return
	}

	b, err := book.Repo().GetBookByIdIncludeAuthor(bookdID)

	// {2 The Unix Programming Environment Rob Pike 1}
	if err != nil {
		bookAndAuthorAPI.ErrorMSG = "Book is not exist record not found"
		exitWithError(bookAndAuthorAPI, w)
		return
	}

	//its
	bookAndAuthorAPI.Book = b
	v, err := bookAndAuthorAPI.Marshal()
	if err != nil {
		bookAndAuthorAPI.ErrorMSG = "Internal server error"
		exitWithError(bookAndAuthorAPI, w)
		return
	}
	w.WriteHeader(http.StatusOK)
	
	w.Write([]byte(v))
}

func exitWithError(bookAndAuthorAPI BookAndAuthorAPIStruct, w http.ResponseWriter) {
	bookAndAuthorAPI.Code = 1
	// bookAndAuthorAPI.errorMSG = "query failed"
	res, err := json.Marshal(bookAndAuthorAPI)
	if err != nil {
		fmt.Println(err)
	}
	http.Error(w, string(res), http.StatusBadGateway)
	return
}

/*
PS C:\Projeler\homework-4-week-5-pMertDogan> curl -v "http://localhost:8080/bookWithAuthor?bookID=1"
VERBOSE: GET with 0-byte payload
VERBOSE: received 65-byte response of content type application/json


StatusCode        : 200
StatusDescription : OK
Content           : {"ID":1,"BookName":"Hobbit","Name":"J.R.R. Tolkien","AuthorID":0}
RawContent        : HTTP/1.1 200 OK
                    Content-Length: 65
                    Content-Type: application/json
                    Date: Tue, 29 Mar 2022 17:08:02 GMT

                    {"ID":1,"BookName":"Hobbit","Name":"J.R.R. Tolkien","AuthorID":0}
Forms             : {}
Headers           : {[Content-Length, 65], [Content-Type, application/json], [Date, Tue, 29 Mar 2022 17:08:02 GMT]}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        : System.__ComObject
RawContentLength  : 65

*/
