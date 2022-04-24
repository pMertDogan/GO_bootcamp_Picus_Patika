package domain

import (
	"encoding/json"
	"github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/domain/book"
)

/*
//https://stackoverflow.com/questions/49461354/go-vet-composite-literal-uses-unkeyed-fields-with-embedded-types
		responseModel = domain.APIStruct{
			Books : []book.Book{*b},
			Code : 1,
			ErrorMsg: "",
		}

*/


type APIStructBook struct {
	Code int
	ErrorMsg string
	Books book.Books
}

func (r *APIStructBook) Marshal() ([]byte, error) {
	return json.Marshal(r)
}


//Its override default String func maybe change name to ToString
func (r *APIStructBook) String() ([]byte, error) {
	st, err :=  json.Marshal(r)
	return st,err
}

