package main

import (
	models "booksAPI/models"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var usageInformation = ` Usage: go run ./main.go [command...] [additionalParams...]

Commads:
	-list command 
		go run main.go list

	-search command
		go run main.go search <bookName>
		go run main.go search Return
		go run main.go search Lord 
		go run main.go search RiNg

	-get command
		go run main.go get <bookID>
		go run main.go get 5

	-delete command
	go run main.go delete <bookID>
	go run main.go delete 5

	-buy command
	go run main.go buy <bookID> <quantity>
	go run main.go buy 5 2
	
	-reset command
	go run main.go reset
`

// lazy :)
// to get json output to use it quicktype
// https://go.dev/play/p/hfMRFkXU4PV

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    books, err := UnmarshalBooks(bytes)
//    bytes, err = books.Marshal()

const JsonLocation = "bookLiblary.json"
const JsonLocationCopy = "bookLiblary_copy.json"

func main() {

	if len(os.Args) < 2 {
		printUsageAndExit("Unsupported")
	}

	//available books , mocked with test data
	bookLiblary := getBooks()

	//get command
	firstQuery := os.Args[1]
	//understand command and redirect to handler. Lower case to make it not case sensitive
	switch strings.ToLower(os.Args[1]) {
	//if the request is list
	case "list":
		//print all books
		listHandler(bookLiblary)
	case "search":
		//is book name provided?

		searchHandler(bookLiblary)
	case "get":
		getHandler(bookLiblary)
	case "delete":

		deleteHandler(bookLiblary)

	case "buy":
		buyHandler(bookLiblary)

	case "reset":
		resetApp()
	default:
		printUsageAndExit(firstQuery + " is not supported.")

	}

}

//remove element using index
func remove(s []models.Book, i int) []models.Book {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

//buy book from liblary if there is enought book in stock
func buyHandler(bookLiblary []models.Book) {
	// arg lenght should be 4
	if len(os.Args) != 4 {
		printUsageAndExit("invalid query")
	}
	// os.Args[0] is not needed, os.Args[1] is command aka "buy"
	id := os.Args[2]
	quantity, err := strconv.Atoi(os.Args[3])

	//if its not convert string value to int
	//example
	//buy 2 bookName is false
	//buy 2 3 is correct
	if err != nil {
		printUsageAndExit("please type int value as for <quantity> like 'buy <id> <quantity>")
	}
	//for each book in liblary check is avaiable then check is quantity is okey
	for i, book := range bookLiblary {

		if book.ID == id {

			if book.StockCount >= quantity {
				book.StockCount = book.StockCount - quantity
				fmt.Println("Great you succesfuly ordered !")
				fmt.Println(book.ToString())
				//change current lib
				bookLiblary[i] = book
				//store updated one
				storeUpdatedBooks(bookLiblary)
				os.Exit(0)
			} else {
				book.ToString()
				printUsageAndExit("Sory , we dont have enought stock")

			}
		}
	}
	printUsageAndExit("book id is not correct , verify book id is exist")
}

// delete book by id if exist
func deleteHandler(bookLiblary []models.Book) {

	// check is command correct
	// must be equal to 3  	delete		 <id>
	//                		os.Args[2]   [3]
	if len(os.Args) != 3 {
		printUsageAndExit("is unsupported command")
	}
	//get id of the book
	queryID := os.Args[2]

	for i, book := range bookLiblary {

		if book.ID == queryID {
			//delete book from books

			storeUpdatedBooks(remove(bookLiblary, i))
			//
			fmt.Println("Book succesuly removed, book information : " + book.ToString())
			os.Exit(0)
		}
	}
	printUsageAndExit(" we dont have a book with ID: " + queryID)
}

//search by word
func searchHandler(bookLiblary models.Books) {
	// [1] = app tep location , [2] = query , [2:] book name
	if len(os.Args) == 3 {
		//join array with space

		searchText := strings.Join(os.Args[2:], " ")

		//to display not available meesage at the end of the search
		notAvailableFlag := true

		//iterate over array
		for _, book := range bookLiblary {
			
			if book.IsNameContains(searchText) {
				fmt.Println(book.ToString())
				notAvailableFlag = false
			}
		}

		if notAvailableFlag {
			fmt.Println("The book '" + searchText + "' is not available. You can get all book name with 'list' command")
		}
		os.Exit(0)
	} else {
		//looks like its unsupported request
		//go run .\main.go search sasdas asdsa
		printUsageAndExit("")
	}
}

//print all books
func listHandler(bookLiblary []models.Book) {
	//just print all of the books
	for _, b := range bookLiblary {
		fmt.Println(b.ToString())
	}

	os.Exit(0)
}

//get book by id
func getHandler(bookLiblary []models.Book) {

	// must be equal to 3  	delete		 <id>
	//                		os.Args[2]   [3]

	if len(os.Args) != 3 {
		printUsageAndExit("unsupported parameters")
	}
	queryID := os.Args[2]

	for _, book := range bookLiblary {

		if book.ID == queryID {
			fmt.Println(book.ToString())
			os.Exit(0)
		}
	}
	printUsageAndExit(" we dont have a book with ID: " + queryID)

}

//If there is an error occured like invalid param .. just print usage information with optional error msg
func printUsageAndExit(optionalText string) {
	fmt.Println(optionalText)
	fmt.Println(usageInformation)
	os.Exit(1)

}

// Get books from source , currenlty its just return dummy values
func getBooks() []models.Book {
	//read our data from dumy json file
	dat, err := os.ReadFile(JsonLocation)
	if err != nil {
		// check json file :)
		printUsageAndExit("Unable read source json file")
	}

	//    books, err := UnmarshalBooks(bytes)
	//    bytes, err = books.Marshal()
	//convert json to to books
	books, err := models.UnmarshalBooks(dat)

	if err != nil {
		fmt.Print(string(dat))

		printUsageAndExit("Unable convert readed bytes to books")
	}
	return books

}

//its just save the books to file to make changes persist
func storeUpdatedBooks(booksToStore models.Books) {
	//convert to JSON
	bytes, err := booksToStore.Marshal()
	if err != nil {
		printUsageAndExit("Unable convert ")
	}
	//store it
	os.WriteFile(JsonLocation, bytes, 0644)

}

//simple method to reset app.
//Some operations is changes our bookLiblary.json . We use this method to restored it to original
func resetApp() {
	dat, err := os.ReadFile(JsonLocationCopy)
	if err != nil {
		// check json file :)
		printUsageAndExit("Unable read source json file")
	}

	os.WriteFile(JsonLocation, []byte(dat), 0644)
	fmt.Println("App Resetted Succesfuly!")
	os.Exit(0)
}
