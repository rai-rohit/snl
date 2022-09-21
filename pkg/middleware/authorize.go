package middleware

import (
	"net/http"
	"snakes_ladders/pkg/utils"
)

func Authorize(hf http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if (r.URL.Path == "/addplayers" || r.URL.Path == "/start" ||
			r.URL.Path == "/rolldice" || r.URL.Path == "/positions" ||
			r.URL.Path == "/position" || r.URL.Path == "/snl") &&
			(r.Method == http.MethodGet || r.Method == http.MethodPost) {

			tkn := r.Header.Get("authorization")
			if tkn != "" {
				tkn = tkn[7:]
			}

			_, err := utils.Parse(tkn)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Please login."))
				return
			} else {
				hf.ServeHTTP(w, r)
				return
			}
		}

		hf.ServeHTTP(w, r)
	})
}
