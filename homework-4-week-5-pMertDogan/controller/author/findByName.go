package authorRest

import (
	"net/http"

	"github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/domain"
	"github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/domain/author"
)

func FindAuthorByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var responseModel domain.APIStructAuthor = domain.APIStructAuthor{}
	
	searchText := r.URL.Query().Get("searchText")

	//we dont need it cause when its provided we get here :)
	// if authorID == "" {
	// 	responseModel.ErrorMsg = "please provide book ıd"
	// 	ExitWithError(&responseModel, w, http.StatusBadRequest)
	// 	return

	// }

	a, err := author.Repo().FindByName(searchText)

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