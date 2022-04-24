package models


//go install booksAPI 
//booksAPI <command>

import (
	"encoding/json"
	"fmt"
	"strings"
)



//define books
type Books []Book

//aka fromJson
func UnmarshalBooks(data []byte) (Books, error) {
	var r Books
	err := json.Unmarshal(data, &r)
	return r, err
}

//aka toJson
func (r *Books) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Book struct {
	ID            string `json:"ID"`
	BookName      string `json:"BookName"`
	NumberOfPages int    `json:"NumberOfPages"`
	StockCount    int    `json:"StockCount"`
	Price         int    `json:"Price"`
	ISBN          string `json:"ISBN"`
	StockCode     string `json:"StockCode"`
	Author        Author `json:"Author"`
}

type Author struct {
	AuthorID string `json:"AuthorID"`
	Name     string `json:"Name"`
}

//convert book params to string to print
func (v Book) ToString() string {

	return "name: " + v.BookName + " id: " + v.ID + " stockCount: " + fmt.Sprint(v.StockCount) + " stockCode: " + fmt.Sprint(v.StockCode) + " price: " + fmt.Sprint(v.Price) + "₺"
}


// check is bookname contains searchText
func (b Book) IsNameContains(searchText string) bool {
	//contains  is caseSensitive
	return strings.Contains(strings.ToLower(b.BookName), strings.ToLower(searchText))
}
