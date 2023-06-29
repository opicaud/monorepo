package main

import (
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net/http"
)

func main() {
	startHttpProxyHealth()
}

func startHttpProxyHealth() {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", "localhost", 50051), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("did not connect")
	}
	defer conn.Close()

	c := grpc_health_v1.NewHealthClient(conn)

	endpoint := runtime.WithHealthzEndpoint(c)
	router := runtime.NewServeMux(endpoint)
	log.Println("Server started")
	log.Println(http.ListenAndServe(":8080", router))
}
