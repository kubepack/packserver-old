package main

import (
	"fmt"
	"net/http"
	"html"
	"log"
	"io/ioutil"
)

func main()  {
	fmt.Println("Server Started...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello request")
		fmt.Fprintf(w, "Hello %q", html.EscapeString(r.URL.Path))
		resp, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(string(resp))
	})


	log.Fatal(http.ListenAndServe(":8080", nil))
}
