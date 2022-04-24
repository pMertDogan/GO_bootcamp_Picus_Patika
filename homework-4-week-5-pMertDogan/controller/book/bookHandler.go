package bookRest

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func BookHandler(r *mux.Router) {
	
	// //*******Without gorilla Mux//can be moved another file
	// //curl -v "http://localhost:8080/statusCheck"
	// http.HandleFunc("/statusCheck", StatusCheck)
	// //curl -v "http://localhost:8080/book?bookID=1"
	// http.HandleFunc("/book", bookHandler)
	// //curl -v "http://localhost:8080/bookWithAuthor?bookID=1"
	// //Maybe we can change it to /book and change the handling method use path variable like /book?author=true
	// http.HandleFunc("/bookWithAuthor", bookRest.GetBookByIdIncludeAuthor)
	// //curl -v "http://localhost:8080/book/find?name=Hob"
	// http.HandleFunc("/book/find", bookRest.FindBookByNameWithoutAuthor)
	// //curl -v "http://localhost:8080/book" POST
	// http.HandleFunc("/book/", bookRest.FindBookByNameWithoutAuthor)
	// http.HandleFunc("/book/enable", bookRest.EnableBook)


	//curl -v "http://localhost:8080/book?bookID=1"
	r.HandleFunc("/book", bookHandler)
	//curl -v "http://localhost:8080/bookWithAuthor?bookID=1"
	//Maybe we can change it to /book and change the handling method use path variable like /book?author=true
	r.HandleFunc("/bookWithAuthor", GetBookByIdIncludeAuthor)
	//curl -v "http://localhost:8080/book/find?name=Hob"
	r.HandleFunc("/book/find",FindBookByNameWithoutAuthor)
	//curl -v "http://localhost:8080/book" POST
	r.HandleFunc("/book/enable", EnableBook)

}



func bookHandler(w http.ResponseWriter, r *http.Request) {
	// var responseModel domain.APIStruct = domain.APIStruct{}
	log.Println(r.Method)
	w.Header().Set("Content-Type", "application/json")

	//redirect to method
	switch r.Method {
	case http.MethodGet:
		log.Println("get book by id")
		GetBookByID(w, r)
		return
	case http.MethodPost:
		bookdID := r.URL.Query().Get("bookID")
		if bookdID != "" {
			UpdateBookQuantity(w, r)
		} else {
			//somethink like
			// bookRest.UpdateBook(w, r)
		}
		return

	case http.MethodDelete:
		SoftDelete(w, r)
		return
	default:
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

}
