package bookRest

import (
	"log"
	"net/http"

	"github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/domain"
	exitError "github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/domain"
	"github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/domain/book"
)

func SoftDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//https://stackoverflow.com/questions/49461354/go-vet-composite-literal-uses-unkeyed-fields-with-embedded-types
	var responseModel domain.APIStructBook = domain.APIStructBook{
		// Books:    nil,
		// Code:     0,
		// ErrorMsg: "",
	}

	bookId := r.URL.Query().Get("bookID")

	log.Println("bookID  " + bookId)

	err := book.Repo().SoftDeleteBook(bookId)

	//check is we can get book without error or not
	if err != nil {
		responseModel.ErrorMsg = "Soft delete error  " + err.Error()
		exitError.ExitWithError(&responseModel, w, http.StatusInternalServerError)
		return
	}

	//get book
	b,err := book.Repo().GetByIDIgnoreSoftDelete(bookId)
	if err != nil {
		responseModel.ErrorMsg = "Book is not exist " + err.Error()
		exitError.ExitWithError(&responseModel, w, http.StatusInternalServerError)
		return
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
