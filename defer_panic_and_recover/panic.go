package main

import (
	"net/http"
)

// a go app panics when it doesnt know how to continue
// execution; a true exception

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go"))
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}

}
