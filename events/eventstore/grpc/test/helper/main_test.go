package main

import (
	"github.com/opicaud/monorepo/events/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldPplop(t *testing.T) {
	var expected = make([]pkg.DomainEvent, 0)
	expected = append(expected, NewDummyEvent("00000000-0000-0000-0000-000000000000", "SHAPE_CREATED", []byte(`{"Nature":"square","Dimensions":[2,3],"Id":"00000000-0000-0000-0000-000000000000","Area":1}`)))

	var request = `{"action": "setup","params": {"events":[{"aggregateId": {"id":"00000000-0000-0000-0000-000000000000"}, "name": "SHAPE_CREATED", "data": "eyJOYXR1cmUiOiJzcXVhcmUiLCJEaW1lbnNpb25zIjpbMiwzXSwiSWQiOiIwMDAwMDAwMC0wMDAwLTAwMDAtMDAwMC0wMDAwMDAwMDAwMDAiLCJBcmVhIjoxfQ=="}]},"state":"a state"}`
	parsedRequest := parseStringBody(request)

	assert.Equal(t, "setup", parsedRequest.Action)
	assert.Equal(t, "a state", parsedRequest.State)

	assert.Len(t, parsedRequest.Params.Events, 1)
	assert.Equal(t, expected[0].Name(), parsedRequest.Params.Events[0].Name())
	assert.Equal(t, expected[0].AggregateId(), parsedRequest.Params.Events[0].AggregateId())
	assert.Equal(t, expected[0].Data(), parsedRequest.Params.Events[0].Data())

}
