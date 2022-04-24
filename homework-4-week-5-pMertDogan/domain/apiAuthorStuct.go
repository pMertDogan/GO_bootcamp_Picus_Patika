package domain

import (
	"encoding/json"

	"github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/domain/author"
)

type APIStructAuthor struct {
	Code     int
	ErrorMsg string
	Author   author.Author
}

func (r *APIStructAuthor) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

//Its override default String func maybe change name to ToString
func (r *APIStructAuthor) String() ([]byte, error) {
	st, err := json.Marshal(r)
	return st, err
}