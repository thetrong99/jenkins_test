package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "The Trong 1999 \n Thai Binh \n")
	})
	http.ListenAndServe(":8463", nil)
}
