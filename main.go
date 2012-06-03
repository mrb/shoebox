package main

import (
	"github.com/bmizerany/noeq.go"
	"github.com/bmizerany/pat"
	"github.com/mrb/riakpbc"

	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	riakc *riakpbc.Conn
	noeqc *noeq.Client
	err   error
)

func GetId(w http.ResponseWriter, req *http.Request) {
	id, _ := noeqc.GenOne()
	stringId := fmt.Sprintf("%d", id)

	log.Print(" [GET] {id: ", stringId, "}")

	io.WriteString(w, stringId)
}

func PostData(w http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	id, _ := noeqc.GenOne()

	stringId := fmt.Sprintf("%d", id)
	stringBody := fmt.Sprintf("%s", body)

	_, err := riakc.StoreObject("data", stringId, stringBody)

	if err != nil {
		log.Print(err)
	}

	log.Print("[POST] {id: ", stringId, "}, {body: ", stringBody, "}")

	io.WriteString(w, stringId)
}

func main() {
	riakc, err = riakpbc.New("127.0.0.1:8087", 1e8, 1e8)
	if err != nil {
		log.Print(err)
		return
	}

	noeqc, err = noeq.New("", "127.0.0.1:4444")
	if err != nil {
		log.Print(err)
		return
	}

	err = riakc.Dial()
	if err != nil {
		log.Print(err)
		return
	}

	m := pat.New()
	m.Get("/id/new", http.HandlerFunc(GetId))
	m.Post("/data", http.HandlerFunc(PostData))

	http.Handle("/", m)

	log.Print("[INFO] ...starting your shoebox....")

	err = http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
