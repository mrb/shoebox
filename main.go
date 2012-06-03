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
	id, err := noeqc.GenOne()
	if err != nil {
		log.Print(" [GET] 500 Server Error")
		http.Error(w, "Server Error", 500)
		return
	}
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
		log.Print("[POST] 500 Server Error")
		http.Error(w, "Server Error", 500)
		return
	}

	log.Print("[POST] {id: ", stringId, "}, {body: ", stringBody, "}")

	io.WriteString(w, stringId)
}

func GetData(w http.ResponseWriter, req *http.Request) {
	stringId := req.URL.Query().Get(":id")

	data, err := riakc.FetchObject("data", stringId)
	if err != nil {
		log.Print(" [GET] 404 Not Found", string(data))
		http.NotFound(w, req)
		return
	}

	log.Print(" [GET] ", string(data))

	io.WriteString(w, string(data))
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
	m.Get("/data/:id", http.HandlerFunc(GetData))

	http.Handle("/", m)

	log.Print("[INFO] ...starting your shoebox....")

	err = http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
