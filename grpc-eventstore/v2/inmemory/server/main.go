package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/opicaud/monorepo/grpc-eventstore/v2/inmemory/server/config"
	pb "github.com/opicaud/monorepo/grpc-eventstore/v2/proto"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
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
	pb.UnimplementedEventStoreServer
	events []*pb.Event
}

func (s *server) Check(ctx context.Context, request *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil

}

func (s *server) Watch(request *grpc_health_v1.HealthCheckRequest, watchServer grpc_health_v1.Health_WatchServer) error {
	return status.Error(codes.Unimplemented, "unimplemented")

}

func (s *server) Save(ctx context.Context, in *pb.Events) (*pb.Response, error) {
	log.Printf("received %s\n", in)
	s.events = append(s.events, in.GetEvent()...)
	return &pb.Response{Message: "", Status: 0, Events: &pb.Events{}}, nil
}
func (s *server) Load(ctx context.Context, in *pb.UUID) (*pb.Response, error) {
	log.Printf("load %s\n", in)
	events := s.search(in.Id)
	r := &pb.Response{Events: &pb.Events{Event: events}}
	return r, nil
}

func (s *server) Remove(ctx context.Context, in *pb.UUID) (*pb.Response, error) {
	log.Printf("WARN: I remove everything, i do not use %s\n", in.Id)
	s.events = []*pb.Event{}
	return &pb.Response{}, nil
}

func (s *server) search(id string) []*pb.Event {
	w := 0
	for _, e := range s.events {
		if e.AggregateId.Id == id {
			s.events[w] = e
			w++
		}
	}
	return s.events[0:w]

}

var (
	port = flag.Int("port", 50052, "The server port")
)

func main() {
	background := context.Background()
	config := config.GetConfigFrom(os.Getenv("OPEN_TELEMETRY_ENABLED"))
	if config.IsTracingEnabled() {
		tp := startTracing(background)
		defer func() {
			if err := tp.Shutdown(background); err != nil {
				log.Printf("Error shutting down tracer provider: %v", err)
			}
		}()
	}
	startGrpc()
}

func startTracing(background context.Context) *sdktrace.TracerProvider {
	exp, err := otlptracegrpc.New(background, otlptracegrpc.WithInsecure())
	if err != nil {
		slog.Error("failed to create trace exporter: %w", err)
	}
	tp := sdktrace.NewTracerProvider(sdktrace.WithBatcher(exp))
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp
}

func startGrpc() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	s := startServer(err)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func startServer(err error) *grpc.Server {
	logger := slog.Default()
	opts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
		logging.WithFieldsFromContext(traceAndSpans),
	}

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.StatsHandler(otelgrpc.NewServerHandler()),
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(InterceptorLogger(logger), opts...),
		))
	srv := &server{}
	pb.RegisterEventStoreServer(s, srv)
	grpc_health_v1.RegisterHealthServer(s, srv)
	return s

}

func InterceptorLogger(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}

func traceAndSpans(ctx context.Context) logging.Fields {
	if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
		return logging.Fields{"traceId", span.TraceID().String(), "spanId", span.SpanID().String()}
	}
	return nil
}
