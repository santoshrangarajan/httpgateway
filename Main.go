package main

import (
	"log"
	"net/http"
)

///https://www.alexedwards.net/blog/making-and-using-middleware
////https://github.com/tensor-programming/go-kit-tutorial

func filter1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing filter1")
		next.ServeHTTP(w, r)
		log.Println("Executing filter1 again")
	})
}

func filter2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing filter2")
		if r.URL.Path == "/foo" {
			return
		}
		next.ServeHTTP(w, r)
		log.Println("Executing filter2 again")
	})
}

func final(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing finalHandler")
	w.Write([]byte("OK"))
}

func main() {
	println("httpgateway....")
	mux := http.NewServeMux()

	finalHandler := http.HandlerFunc(final)
	mux.Handle("/", filter1(filter2(finalHandler)))

	log.Println("Listening on :9000...")
	err := http.ListenAndServe(":9000", mux)
	log.Fatal(err)
}

/*

start server
create a service
create filter which intercepts request


*/
