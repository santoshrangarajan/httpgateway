package main

import (
	"log"
	"net/http"
)

var db DbService

///https://www.alexedwards.net/blog/making-and-using-middleware

func init() {
	db = dbService{}
}

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

func dblist(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing dblist")
	db.list()
	w.Write([]byte("OK"))
}

/*func dbHandlerFunc(db interface{}) func(http.ResponseWriter, *http.Request) {
	if db == nil {
		panic("nil dbService session!")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing dblist")
		db.open()
		w.Write([]byte("OK"))
	}
}*/

func main() {
	println("httpgateway....")

	mux := http.NewServeMux()

	finalHandler := http.HandlerFunc(final)
	dbHandler := http.HandlerFunc(dblist)

	mux.Handle("/", filter1(filter2(finalHandler)))
	mux.Handle("/list", dbHandler)

	//mux.Handler("/list", dbHandler)

	log.Println("Listening on :9000...")
	err := http.ListenAndServe(":9000", mux)
	log.Fatal(err)
}

/*

start server
create a service
create filter which intercepts request


*/
