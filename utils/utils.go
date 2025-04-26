package utils

import "net/http"

func GenerateHandler(html []byte) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "text/html")
		w.Write(html)
	}
}
