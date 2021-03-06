package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rubiagatra/cloud-native-go/api"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", api.EchoHandleFunc)
	http.HandleFunc("/api/hello", api.HelloHandleFunc)

	http.HandleFunc("/api/manga", api.ListOfMangaHandleFunc)
	http.HandleFunc("/api/manga/", api.MangaHandleFunc)
	log.Println("Server run at port ", port())
	http.ListenAndServe(port(), nil)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to Cloud Native Go")
}
