package domain

import (
	"fmt"

	"github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/data/database"
	"github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/domain/author"
	"github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/domain/book"
)

//drop old tables to clean all
func DropTables(authorRepository *author.AuthorRepository, bookRepository *book.BookRepository) error {
	err := database.DropTables(authorRepository, bookRepository)
	if err != nil {
		return fmt.Errorf("cannot drop tables" + err.Error())
	}
	// os.Exit(1)
	return nil
}

// migrate and save json data to sql
func InitDB(authorRepository *author.AuthorRepository, bookRepository *book.BookRepository) {
	fmt.Println("reset start")
	//create our DB struct on SQL
	database.MigrateDatabase(authorRepository, bookRepository)

	//store local data to sql
	database.ReadFilesAndSaveThemToDB(authorRepository, bookRepository)
	fmt.Println("reset end")
	fmt.Println("tip run main with -reset")
	// os.Exit(0)
}