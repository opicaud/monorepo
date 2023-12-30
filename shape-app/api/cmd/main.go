package main

import (
	"context"
	"flag"
	"fmt"
	pkg3 "github.com/opicaud/monorepo/events/eventstore/pkg"
	v2beta "github.com/opicaud/monorepo/events/pkg/v2beta"
	pb "github.com/opicaud/monorepo/shape-app/api/proto"
	pkg2 "github.com/opicaud/monorepo/shape-app/domain/pkg"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
)

type server struct {
	pb.UnimplementedShapesServer
}

var eventStore, errConfig = pkg3.NewEventsFrameworkFromConfigV2(os.Getenv("CONFIG"))

func (s *server) Create(ctx context.Context, in *pb.ShapeRequest) (*pb.Response, error) {

	slog.Info("context", "traceId", trace.SpanFromContext(ctx).SpanContext().TraceID(), "spanId", trace.SpanFromContext(ctx).SpanContext().SpanID())
	factory := pkg2.New()
	var command = factory.NewCreationShapeCommand(in.Shapes.Shape, in.Shapes.Dimensions...)
	if errConfig != nil {
		panic(errConfig)
	}
	handler := factory.NewCommandHandlerBuilder().
		WithEventStore(eventStore).
		WithEventsEmitter(&v2beta.StandardEventsEmitter{}).
		Build()
	err := handler.Execute(command, factory.NewShapeCommandApplier(eventStore))
	if err != nil {
		return nil, err
	}
	message := pb.Message{
		Code: uint32(codes.OK),
	}
	response := pb.Response{
		Message: &message,
	}
	return &response, nil
}

func (s *server) Check(ctx context.Context, request *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	if request.GetService() == "eventstore" {
		return checkHealth(eventStore.GetHealthClient(), request)
	}
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil

}

func (s *server) Watch(request *grpc_health_v1.HealthCheckRequest, watchServer grpc_health_v1.Health_WatchServer) error {
	return status.Error(codes.Unimplemented, "unimplemented")

}

func checkHealth(eventStoreHealthClient grpc_health_v1.HealthClient, request *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return eventStoreHealthClient.Check(context.Background(), request)
}

func main() {
	tp := initTracerProvider()
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()
	startGrpc()
}

func initTracerProvider() *sdktrace.TracerProvider {
	tp := sdktrace.NewTracerProvider()
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp
}

func startGrpc() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	s := startServer(err)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func startServer(err error) *grpc.Server {
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.StatsHandler(otelgrpc.NewServerHandler()))
	srv := &server{}
	pb.RegisterShapesServer(s, srv)
	grpc_health_v1.RegisterHealthServer(s, srv)
	return s

}
