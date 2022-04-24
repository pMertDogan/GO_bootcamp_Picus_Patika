package admin

import (
	"net/http"
	// "os"

	"github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/domain"
	"github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/domain/author"
	"github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/domain/book"
)

func InitDatabase(w http.ResponseWriter, r *http.Request) {
	domain.InitDB(author.Repo(),  book.Repo())
	// os.Exit(1)
}