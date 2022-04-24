package authorRest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func AuthorHandler(r *mux.Router) {
	

	// r.HandleFunc("/author", authorRest.AuthorHandler)
	authorSlash := r.PathPrefix("/author").Subrouter()

	//"/author/{name}/"
	authorSlash.HandleFunc("/{authorID}", GetAuthorByID)
	//Do not use 
	//authorSlash.HandleFunc("/{authorID}", GetAuthorByID) cause its only accept request with / not handle the author should be author/
	authorSlash.HandleFunc("", getHandler).Methods(http.MethodGet)
	authorSlash.HandleFunc("", getHandler).Methods(http.MethodPost)
	// authorSlash.HandleFunc("/", RootHandler).Methods(http.MethodPost)

	// SAME AS r.HandleFunc("/author/{authorID}", GetAuthorByID)

	// //register mux
	// http.Handle("/", r)

}



func getHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("getHandler")
	

	if searchText := r.URL.Query().Get("searchText"); searchText != "" {
		fmt.Println("searchText  is : " + searchText)
		FindAuthorByName(w, r)
	}


}

func postHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("postHandler")
	

	if searchText := r.URL.Query().Get("searchText"); searchText != "" {
		fmt.Println("searchText  is : " + searchText)
		FindAuthorByName(w, r)
	}


}
