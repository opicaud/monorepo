package test

import (
	"context"
	"fmt"
	"github.com/beorn7/floats"
	"github.com/cucumber/godog"
	"github.com/google/uuid"
	cqrs "github.com/opicaud/monorepo/cqrs/pkg/v2beta1"
	v2beta1 "github.com/opicaud/monorepo/events/eventstore/pkg/v2beta1"
	"github.com/opicaud/monorepo/events/pkg"
	"github.com/opicaud/monorepo/shape-app/domain/internal"
	pkg2 "github.com/opicaud/monorepo/shape-app/domain/pkg"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"log"
	"strconv"
	"testing"
)

type TestContext struct {
	length float32
	width  float32
	radius float32
}

type key int

const testContextKey key = 0
const idKey key = 1

var (
	query    = BDDQueryShape{shapes: make(map[uuid.UUID]BDDShape)}
	factory  = pkg2.New()
	store, _ = v2beta1.NewEventsFrameworkFromConfig("")
)

func iCreateARectangle(ctx context.Context) context.Context {
	testContext := ctx.Value(testContextKey).(TestContext)
	ctx = executeShapeCommand(ctx, factory.NewCreationShapeCommand("rectangle", []float32{testContext.length, testContext.width}...))
	return ctx
}

func iCreateACircle(ctx context.Context) context.Context {
	testContext := ctx.Value(testContextKey).(TestContext)
	ctx = executeShapeCommand(ctx, factory.NewCreationShapeCommand("circle", []float32{testContext.radius}...))
	return ctx
}

func itAreaIs(ctx context.Context, arg1 string) error {
	id := ctx.Value(idKey).(uuid.UUID)
	newArea := query.GetById(id).area
	f, _ := strconv.ParseFloat(arg1, 32)
	if !floats.AlmostEqual(float64(newArea), f, 0.01) {
		return fmt.Errorf("expected %f, found %f", f, newArea)
	}
	return nil
}

func lengthOfAndWidthOf(ctx context.Context, arg1 int, arg2 int) (context.Context, error) {
	ctx = context.WithValue(ctx, testContextKey, TestContext{length: float32(arg1), width: float32(arg2)})
	return ctx, nil
}

func radiusOf(ctx context.Context, arg1 int) (context.Context, error) {
	ctx = context.WithValue(ctx, testContextKey, TestContext{radius: float32(arg1)})
	return ctx, nil
}

func anExisting(ctx context.Context, nature string) context.Context {
	id := query.GetByNature(nature)
	ctx = context.WithValue(ctx, idKey, id)
	return ctx
}

func iStretchItBy(ctx context.Context, arg1 int) error {
	id := ctx.Value(idKey).(uuid.UUID)
	executeShapeCommand(ctx, factory.NewStretchShapeCommand(id, float32(arg1)))
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	initTracing()
	ctx.Step(`^I create a rectangle$`, iCreateARectangle)
	ctx.Step(`^it area is "([^"]*)"$`, itAreaIs)
	ctx.Step(`^length of (\d+) and width of (\d+)$`, lengthOfAndWidthOf)
	ctx.Step(`^I create a circle$`, iCreateACircle)
	ctx.Step(`^radius of (\d+)$`, radiusOf)
	ctx.Step(`^an existing "([^"]*)"$`, anExisting)
	ctx.Step(`^I stretch it by (\d+)$`, iStretchItBy)
}

func TestFeatures(t *testing.T) {

	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t,
			Tags:     "@Done", // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func executeShapeCommand(ctx context.Context, command cqrs.Command[internal.CommandApplier]) context.Context {
	tracer := otel.Tracer("test-scenarios")
	ctx, span := tracer.Start(ctx, "feature")
	defer span.End()
	handler := pkg2.New().NewCommandHandlerBuilder().
		WithSubscriber(&Subscriber{query: &query}).
		WithEventsEmitter(&cqrs.StandardEventsEmitter{}).
		WithEventStore(store).
		Build()
	ctx, err := handler.Execute(ctx, command, factory.NewShapeCommandApplier(store))
	if err != nil {
		log.Fatal(err)
	}
	errTracing := assertTracing(span.SpanContext(), trace.SpanContextFromContext(ctx))
	if errTracing != nil {
		log.Fatal(errTracing)
	}
	return ctx
}

func assertTracing(span trace.SpanContext, fromContext trace.SpanContext) error {
	if !fromContext.IsValid() || !span.IsValid() {
		return fmt.Errorf("tracing is not correct")
	}
	if fromContext.SpanID().String() == span.SpanID().String() {
		return fmt.Errorf("tracing is not correct, span are equals: expected %s, found %s", span.SpanID(), fromContext.SpanID())
	}
	if fromContext.TraceID().String() != span.TraceID().String() {
		return fmt.Errorf("tracing is not correct, traces are not equals: expected %s, found %s", span.TraceID(), fromContext.TraceID())
	}
	return nil
}

func initTracing() {
	provider := sdktrace.NewTracerProvider()
	otel.SetTracerProvider(provider)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
}

type Subscriber struct {
	query QueryShapeModel
}

func (s *Subscriber) Update(ctx context.Context, eventsChn chan []pkg.DomainEvent) context.Context {
	events := <-eventsChn
	for _, e := range events {
		ctx = context.WithValue(ctx, idKey, e.AggregateId())
		switch v := e.(type) {
		default:
			log.Fatal(fmt.Errorf("DomainEvent type %T not handled", v))
		case internal.Created:
			bddShape := BDDShape{id: e.AggregateId(), nature: e.(internal.Created).Nature, area: e.(internal.Created).Area}
			s.query.Save(bddShape)
		case internal.Stretched:
			bddShape := s.query.GetById(e.AggregateId())
			bddShape.area = e.(internal.Stretched).Area
			s.query.Save(bddShape)
		}
	}
	return ctx
}

type BDDShape struct {
	id     uuid.UUID
	nature string
	area   float32
}

type QueryShapeModel interface {
	Save(shape BDDShape)
	GetById(id uuid.UUID) BDDShape
	GetByNature(nature string) uuid.UUID
	GetAll() []BDDShape
}

type BDDQueryShape struct {
	shapes map[uuid.UUID]BDDShape
}

func (b *BDDQueryShape) Save(shape BDDShape) {
	b.shapes[shape.id] = shape
}

func (b *BDDQueryShape) GetById(id uuid.UUID) BDDShape {
	if b.shapes[id] == (BDDShape{}) {
		log.Fatal(fmt.Errorf("id %s not found", id))
	}
	return b.shapes[id]
}

func (b *BDDQueryShape) GetByNature(nature string) uuid.UUID {
	for _, value := range b.shapes {
		if value.nature == nature {
			return value.id
		}
	}
	log.Fatal(fmt.Errorf("shape of nature %s not found", nature))
	return uuid.Nil
}

func (b *BDDQueryShape) GetAll() []BDDShape {
	var values []BDDShape
	for _, value := range b.shapes {
		values = append(values, value)
	}
	return values
}
