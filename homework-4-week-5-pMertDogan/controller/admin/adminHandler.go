package admin

import (

	"github.com/gorilla/mux"
)

func AdminHandler(r *mux.Router) {
	// create mux
	// r.Use(AdminMiddleware)			 
	r.HandleFunc("/admin/drop", DropTables)
	r.HandleFunc("/admin/init", InitDatabase)

}

