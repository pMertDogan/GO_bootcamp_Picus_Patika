package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/controller/admin"
	authorRest "github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/controller/author"
	bookRest "github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/controller/book"
)

//setup http controllers
func SetupControllers() {
	r := mux.NewRouter()

	handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})

	//set handlers whitelist if needed
	// handlers.AllowedMethods([]string{"POST", "GET", "PUT", "PATCH"})


	//curl -v "http://localhost:8080/statusCheck"
	r.HandleFunc("/statusCheck", StatusCheck)

	//Bookhandler is a function that will handle all the request for book
	bookRest.BookHandler(r)

	// http.HandleFunc("/book", bookRest.UpdateBookWithID) NOT IMPLEMENTED YET
	//******Middleware
	// r.Use(admin.AdminLogger)
	//add admin middleware
	r.Use(admin.AdminMiddleware)

	//**** Author Handles it with Gorilla Mux
	//setup admin routes
	authorRest.AuthorHandler(r)

	//*****setup admin routes
	admin.AdminHandler(r)

	//register gorilla mux
	http.Handle("/", r)

	// srv := &http.Server{
	// 	Addr:         "0.0.0.0:8080",
	// 	WriteTimeout: time.Second * 15,
	// 	ReadTimeout:  time.Second * 15,
	// 	IdleTimeout:  time.Second * 60,
	// 	Handler:      r,
	// }
	// log.Println(srv.ListenAndServe())

	// log.Println(http.ListenAndServeTLS(":8080", "certFile", "keyFile", nil))
	log.Println(http.ListenAndServe(":8080", r))

}
