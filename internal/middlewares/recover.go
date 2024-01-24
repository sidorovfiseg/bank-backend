package middlewares

import (
	"fmt"
	"net/http"
)

// TODO unnecessary use of fmt.Sprintf 

func Recover(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func ()  {
			if err := recover(); err != nil {				
				http.Error(w, fmt.Sprintf("PANIC: #{err}"), http.StatusInternalServerError)
			}
		}()

		handler.ServeHTTP(w, r)
	})
}