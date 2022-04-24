package main

import (
	"flag"
	"log"

	"github.com/joho/godotenv"
	controller "github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/controller"
	database "github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/data/database"
	"github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/domain"
	"github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/domain/author"
	"github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/domain/book"
)

var resetAppDB = flag.Bool("init", false, "Migrate tables and save default values thats readed by json files")
var dropTable = flag.Bool("clear", false, "Drop authors and books tables for clear SQL data")

//init env and parse flags
func init() {
	//Load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file. \n ERROR : " + err.Error())
	}
	//parse reset flag or any other
	// flag.Parse()

}

//https://documenter.getpostman.com/view/11892665/UVyq1HoD
func main() {
	//init database
	db, err := database.ConnectPostgresDB()
	if err != nil {
		log.Fatal("cannot connect to database")
	}

	//init book repo
	bookRepository := book.BookRepoInit(db)
	authorRepository := author.AuthorRepoInit(db)

	//check if the user add  droptable flag
	if *dropTable {
		//drop old tables to clean all
		domain.DropTables(authorRepository, bookRepository)
	}

	//check is user add reset  flag
	if *resetAppDB {
		//migrate and json to SQL
		domain.InitDB(authorRepository, bookRepository)
	}

	//setup the endpoint controllers
	controller.SetupControllers()
	public.PulicControllerDef()
}


