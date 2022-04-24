package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//available books , mocked with test data
	bookLiblary := []string{"Hobbit", "World War Z", "Kolpaçino", "Ayla", "Lord of the Ring: The Return of the King"}
	//can be list or search
	firstQuery := os.Args[1]
	//logic time . lower case to support Search seArch search..
	switch strings.ToLower(os.Args[1]) {
	//if the request is list
	case "list":
		//print all books
		for _, v := range bookLiblary {
			fmt.Println(v)
		}
		//close app. 0 means succes
		os.Exit(0)

	//if the query is search
	case "search":
		//is book name provided?
		if len(os.Args) >= 3 {
			// [1] = app tep location , [2] = query , [2:] book name
			//join array with space
			requestedBook := strings.Join(os.Args[2:], " ")
			//iterate over array
			for _, v := range bookLiblary {
				//is book exist?
				if strings.ToLower(v) == strings.ToLower(requestedBook) {
					//let's make our book lover happy :)
					fmt.Println(requestedBook + " is available")
					//to exit from for loop and switch case
					os.Exit(0)

				}
			}
			//print book is not avaiable
			fmt.Println("The book '" + requestedBook + "' is not available. You can get all book name with 'list' command")
			os.Exit(0)
		} else {
			//looks like its unsupported request
			fmt.Println("Supported querys are [list, search <bookName>]")
			os.Exit(0)
		}
	default:
		//inform user about supported commands
		fmt.Println(firstQuery + " is not supported. Supported querys are [list, search <bookName>] ")
	}

}
