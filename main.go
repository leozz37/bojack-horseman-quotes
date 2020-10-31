package main

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func NewRandomXingamento() string {
	return fmt.Sprintf("%s",
		getRandomFromFile("data.txt"),)
}

func getRandomFromFile(file string) string {
	rand.Seed(time.Now().UnixNano())
	all, _ := ioutil.ReadFile(file)
	list := bytes.Split(all, []byte("\n"))
	return string(list[rand.Intn(len(list))])
}

func GetWord(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, NewRandomXingamento())
}

func main() {
	router := mux.NewRouter()
	//port := os.Getenv("PORT")
	router.HandleFunc("/", GetWord).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
