package authorRest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/domain"
	"github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/domain/author"
)

func GetAuthorByID(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")

	var responseModel domain.APIStructAuthor = domain.APIStructAuthor{}
	//get author with vars of the mux
	vars := mux.Vars(r)

	fmt.Println("author id is : " + vars["authorID"])
	authorID := vars["authorID"]

	//we dont need it cause when its provided we get here :)
	// if authorID == "" {
	// 	responseModel.ErrorMsg = "please provide book ıd"
	// 	ExitWithError(&responseModel, w, http.StatusBadRequest)
	// 	return

	// }

	a, err := author.Repo().GetByID(authorID)

	if err != nil {
		responseModel.ErrorMsg = "unable get author internal error"
		ExitWithError(&responseModel, w, http.StatusBadRequest)
		return
	}

	responseModel.Author = *a

	json, _ := responseModel.Marshal()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(json))

}
