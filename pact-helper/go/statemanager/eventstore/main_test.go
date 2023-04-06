package main

import (
	"github.com/opicaud/monorepo/events/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldPplop(t *testing.T) {
	var expected = make([]pkg.DomainEvent, 0)
	expected = append(expected, NewDummyEvent("363c80b0-d2ae-11ed-afa1-0242ac120002", "name", []byte("plop")))

	var request = `{"action": "setup","params": {"events": [{"id": "363c80b0-d2ae-11ed-afa1-0242ac120002", "name": "name", "data": "cGxvcA=="}]},"state": "expected State"}`
	parsedRequest := parseStringBody(request)

	assert.Equal(t, "setup", parsedRequest.Action)
	assert.Equal(t, "expected State", parsedRequest.State)

	assert.Len(t, parsedRequest.Params.Events, 1)
	assert.Equal(t, expected[0].Name(), parsedRequest.Params.Events[0].Name())
	assert.Equal(t, expected[0].AggregateId(), parsedRequest.Params.Events[0].AggregateId())
	assert.Equal(t, expected[0].Data(), parsedRequest.Params.Events[0].Data())

}
