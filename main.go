package main

import (
	"github.com/bmizerany/noeq.go"
	"github.com/bmizerany/pat"
	"github.com/mrb/riakpbc"

	"fmt"
	"io"
	"log"
	"net/http"
)

var (
	riakc *riakpbc.Conn
	noeqc *noeq.Client
)

func GetId(w http.ResponseWriter, req *http.Request) {
	id, _ := noeqc.GenOne()
	stringId := fmt.Sprintf("%d", id)
	io.WriteString(w, stringId)
}

func main() {
	riakc, err := riakpbc.Dial("127.0.0.1:8087")
	if err != nil {
		log.Print(err)
	}

	log.Print(riakc)

	noeqc, err = noeq.New("", "127.0.0.1:4444")
	if err != nil {
		log.Print(err)
	}

	log.Print(noeqc)

	m := pat.New()
	m.Get("/id/new", http.HandlerFunc(GetId))

	http.Handle("/", m)
	err = http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
