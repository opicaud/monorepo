package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

type Body struct {
	Action string
	State  string
	Params struct {
		Events []DummyEvent
	}
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	_, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)

}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/event", createEvent).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))

}

func Parse(foo []byte) Body {
	return parseStringBody(string(foo))
}

func parseStringBody(body string) Body {

	b := Body{}
	err := json.Unmarshal([]byte(body), &b)
	if err != nil {
		fmt.Println(err)
	}

	return b

}

func NewDummyEvent(id string, name string, data []byte) DummyEvent {
	parse, _ := uuid.Parse(id)
	return DummyEvent{Id: parse, Name: name, Data: data}
}

type DummyEvent struct {
	Id   uuid.UUID
	Name string
	Data []byte
}

func (d *DummyEvent) AggregateId() uuid.UUID {
	return d.Id
}
