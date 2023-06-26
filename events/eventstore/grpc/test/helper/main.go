package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/opicaud/monorepo/events/eventstore/pkg"
	pkg2 "github.com/opicaud/monorepo/events/pkg"
	"io"
	"log"
	"net/http"
	"os"
)

type Body struct {
	Action string
	State  string
	Params struct {
		Events []DummyEvent
	}
}

func createEvent(w http.ResponseWriter, r *http.Request) {

	request, err := io.ReadAll(r.Body)
	log.Println("API received: {}", string(request))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	config, err := pkg.NewEventsFrameworkFromConfigV2(os.Getenv("CONFIG"))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	parsedRequest := Parse(request)
	if parsedRequest.Action == "teardown" {
		for _, event := range parsedRequest.Params.Events {
			_ = config.Remove(event.Id)
		}
	} else {
		err = config.Save(convert(parsedRequest.Params.Events)...)
		log.Println(err)
	}
	w.WriteHeader(http.StatusCreated)

}

func healthz(w http.ResponseWriter, request *http.Request) {
	w.WriteHeader(http.StatusOK)
	body := "{\"status\": \"UP\"}"
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(body))
}

func convert(events []DummyEvent) []pkg2.DomainEvent {
	var expected = make([]pkg2.DomainEvent, 0)
	for _, event := range events {
		expected = append(expected, DummyEvent{Id: event.AggregateId(), Named: event.Name(), Datad: event.Data()})
	}
	return expected
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/event", createEvent).Methods("POST")
	router.HandleFunc("/healthz", healthz).Methods("GET")
	log.Println("Server started")
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

func NewDummyEvent(id string, name string, data []byte) pkg2.DomainEvent {
	parse, _ := uuid.Parse(id)
	return &DummyEvent{Id: parse, Named: name, Datad: data}
}

type DummyEvent struct {
	Id    uuid.UUID `json:"id"`
	Named string    `json:"name"`
	Datad []byte    `json:"data"`
}

func (d DummyEvent) AggregateId() uuid.UUID {
	return d.Id
}

func (d DummyEvent) Name() string {
	return d.Named
}

func (d DummyEvent) Data() []byte {
	return d.Datad
}