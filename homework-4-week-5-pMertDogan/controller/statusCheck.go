package check

import (
	"fmt"
	"github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/data/database"
	"net/http"
)

//curl -v "http://localhost:8080/statusCheck"
//To check our service is running
func StatusCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//Check is database	conected with gorm built-in function
	err := database.StatusCheck()
	if err != nil {
		w.Write([]byte("There is a problem with DB connection"))
		//Example
		//failed to connect to `host=localhost user=postgres database=bookLiblary`: dial error (dial tcp 127.0.0.1:5432:
		//connectex: No connection could be made because the target machine actively refused it.)Service is running
		fmt.Println(err.Error())
		//change header to 500
		w.WriteHeader(http.StatusInternalServerError)
		//There is a problem with DB connectionService is running
	}

	w.Write([]byte("Service is running"))
}
