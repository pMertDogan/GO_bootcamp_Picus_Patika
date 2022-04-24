package bookRest

import (
	// "encoding/json"
	"log"

	"net/http"

	"github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/domain"
	"github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/domain/book"
)

//update book quantity
func UpdateBookQuantity(w http.ResponseWriter, r *http.Request) {

	var responseModel domain.APIStructBook = domain.APIStructBook{}

	w.Header().Set("Content-Type", "application/json")

	// var xBook book.Book
	// w.Header().Set("Content-Type", "application/json")
	// dec := json.NewDecoder(r.Body)
	// err := dec.Decode(&xBook)
	// dec.DisallowUnknownFields()

	
	// if err != nil {
	// 	responseModel.ErrorMsg = "We cant parse json" + err.Error()
	// 	domain.ExitWithError(&responseModel, w, http.StatusBadRequest)
	// 	// http.Error(w, "We cant parse json" + err.Error(), http.StatusBadRequest)
	// 	return
	// }

	bookdID := r.URL.Query().Get("bookID")
	log.Println("bookID: " + bookdID)

	//check is bookID provided
	if bookdID == "" {
		responseModel.ErrorMsg = "bookID is required"
		domain.ExitWithError(&responseModel, w, http.StatusBadRequest)
		return
	}

	//check is quantity provided
	quantity := r.URL.Query().Get("quantity")
	log.Println("quantity: " + quantity)

	if quantity == "" {
		responseModel.ErrorMsg = "quantity is required"
		domain.ExitWithError(&responseModel, w, http.StatusBadRequest)
		return
	}

	err := book.Repo().UpdateBookQuantity(bookdID, quantity)

	if err != nil {
		responseModel.ErrorMsg = "unable update book"
		domain.ExitWithError(&responseModel, w, http.StatusBadRequest)
		return
	}
	//get book
	GetBookByID(w,r)

}
