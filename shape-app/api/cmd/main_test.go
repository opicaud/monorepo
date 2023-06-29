package main

import (
	"context"
	"fmt"
	"github.com/opicaud/monorepo/shape-app/domain/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"testing"
	"time"
)

func TestHealthCheckWithEventStoreRunning(t *testing.T) {
	store := test.NewFakeInMemoryEventStore()
	mockHealthClient := store.MockedHealthClient.Mock
	mockHealthClient.On("Check", mock.Anything).Return(
		&grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil)
	request := grpc_health_v1.HealthCheckRequest{Service: "eventstore"}

	healthCheckResponse, err := checkHealth(store.GetHealthClient(), &request)

	assert.NoError(t, err)
	assert.Equal(t, grpc_health_v1.HealthCheckResponse_SERVING, healthCheckResponse.Status)
	mockHealthClient.AssertExpectations(t)

}

func TestHealthCheck(t *testing.T) {
	result := make(chan int, 1)
	go start(result)
	port := <-result
	close(result)
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", "localhost", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("did not connect")
	}
	defer conn.Close()

	c := grpc_health_v1.NewHealthClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	check, err := c.Check(ctx, &grpc_health_v1.HealthCheckRequest{})
	assert.NoError(t, err)
	assert.Equal(t, grpc_health_v1.HealthCheckResponse_SERVING, check.Status)
}

func start(result chan int) {
	lis, err := net.Listen("tcp", ":0")
	s := startServer(err)
	result <- lis.Addr().(*net.TCPAddr).Port
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
