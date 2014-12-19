package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/fabioberger/dockerize-tutorial/config"
	"github.com/fabioberger/dockerize-tutorial/models"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		posts := models.GetAllPosts()
		fmt.Println(posts)
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	n := negroni.Classic()
	n.UseHandler(mux)

	config.Init()
	models.Init()

	n.Run(":4000")
}
