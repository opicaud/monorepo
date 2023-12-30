package cli

import (
	"context"
	"github.com/opicaud/monorepo/shape-app/api/proto"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"time"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

type Client struct{}
type Features interface {
	CreateShape(address string, request *proto.ShapeRequest) (*proto.Message, error)
}

func (f Client) CreateShape(address string, request *proto.ShapeRequest) (*proto.Message, error) {
	provider := sdktrace.NewTracerProvider()
	otel.SetTracerProvider(provider)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	tracer := otel.Tracer("grpc-client-shape")
	ctx, span := tracer.Start(context.Background(), "CreateShape")
	slog.Info("context", "trace", span.SpanContext().TraceID(), "span", span.SpanContext().SpanID())
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		logger.Error("did not connect: %v", err)
		return nil, err
	}
	defer conn.Close()

	c := proto.NewShapesClient(conn)

	ctx, cancel := context.WithTimeout(ctx, time.Second)

	defer span.End()
	defer cancel()
	r, err := c.Create(ctx, request)

	if err != nil {
		return nil, err
	}
	return r.GetMessage(), nil
}
