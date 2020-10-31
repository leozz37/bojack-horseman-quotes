package main

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func GetRandomQuote() string {
	return fmt.Sprintf("%s",
		GetRandomFromFile("data.txt"))
}

func GetRandomFromFile(file string) string {
	rand.Seed(time.Now().UnixNano())
	all, _ := ioutil.ReadFile(file)
	list := bytes.Split(all, []byte("\n"))
	return string(list[rand.Intn(len(list))])
}

func GetQuote(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, GetRandomQuote())
}

func main() {
	port := os.Getenv("PORT")
	log.Print("SERVER STARTED AT PORT: " + port)

	router := mux.NewRouter()
	router.HandleFunc("/", GetQuote).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
