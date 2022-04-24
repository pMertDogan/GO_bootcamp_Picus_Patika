package admin

import (
	"fmt"
	"net/http"
	"strings"
)

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("AdminMiddleware")
		token := r.Header.Get("Authorization")
		if strings.HasPrefix(r.URL.Path, "/admin/") {
			if token != "PleaseTypeTokenHashMashasdasHere" {
				next.ServeHTTP(w, r)
				fmt.Println("token is valid")
			} else {
				fmt.Println("token is not valid")
				http.Error(w, "Token not found or invalid", http.StatusUnauthorized)
			}
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
