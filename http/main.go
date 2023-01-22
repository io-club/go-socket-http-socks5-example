package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var method string = r.Method

		w.WriteHeader(http.StatusOK)
		w.Header().Add("content-Type", "text/plain")
		w.Write([]byte(method))
	})
	http.ListenAndServe(":8080", nil)
}
