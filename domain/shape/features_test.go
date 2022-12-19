package shape

import (
	"context"
	"fmt"
	"github.com/beorn7/floats"
	"github.com/cucumber/godog"
	"github.com/google/uuid"
	"log"
	"strconv"
	"testing"
	"trackclear.be/example/shape/domain/adapter"
)

type TestContext struct {
	length float32
	width  float32
	radius float32
}

var (
	query    = BDDQueryShape{shapes: make(map[uuid.UUID]BDDShape)}
	factory  = NewFactory()
	provider = adapter.NewInfraBuilder().WithEventStore(adapter.NewInMemoryEventStore()).Build()
)

func iCreateARectangle(ctx context.Context) context.Context {
	testContext := ctx.Value("testContext").(TestContext)
	return makeShapeCommand(ctx, "rectangle", testContext.length, testContext.width)
}

func iCreateACircle(ctx context.Context) context.Context {
	testContext := ctx.Value("testContext").(TestContext)
	return makeShapeCommand(ctx, "circle", testContext.radius)
}

func itAreaIs(ctx context.Context, arg1 string) error {
	id := ctx.Value("id").(uuid.UUID)
	newArea := query.GetById(id).area
	f, _ := strconv.ParseFloat(arg1, 32)
	if !floats.AlmostEqual(float64(newArea), f, 0.01) {
		return fmt.Errorf("expected %f, found %f", f, newArea)
	}
	return nil
}

func lengthOfAndWidthOf(ctx context.Context, arg1 int, arg2 int) (context.Context, error) {
	ctx = context.WithValue(ctx, "testContext", TestContext{length: float32(arg1), width: float32(arg2)})
	return ctx, nil
}

func radiusOf(ctx context.Context, arg1 int) (context.Context, error) {
	ctx = context.WithValue(ctx, "testContext", TestContext{radius: float32(arg1)})
	return ctx, nil
}

func anExisting(ctx context.Context, nature string) context.Context {
	id := query.GetByNature(nature)
	ctx = context.WithValue(ctx, "id", id)
	return ctx
}

func iStretchItBy(ctx context.Context, arg1 int) error {
	id := ctx.Value("id").(uuid.UUID)
	makeStretchCommand(ctx, id, float32(arg1))
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
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
func makeShapeCommand(ctx context.Context, nature string, dimensions ...float32) context.Context {
	command := factory.NewCreationShapeCommand(nature, dimensions...)
	ctx = executeShapeCommand(ctx, command)
	return ctx
}

func makeStretchCommand(ctx context.Context, id uuid.UUID, stretchBy float32) context.Context {
	command := factory.NewStretchShapeCommand(id, stretchBy)
	ctx = executeShapeCommand(ctx, command)
	return ctx
}

func executeShapeCommand(ctx context.Context, command Command) context.Context {
	err := NewShapeCreationCommandHandlerBuilder().
		WithSubscriber(&Subscriber{ctx: &ctx, query: &query}).
		WithInfraProvider(provider).
		Build().
		Execute(command)
	if err != nil {
		log.Fatal(err)
	}
	return ctx
}

type Subscriber struct {
	ctx   *context.Context
	query QueryShapeModel
}

func (s *Subscriber) Update(events []adapter.DomainEvent) {
	for _, e := range events {
		*s.ctx = context.WithValue(*s.ctx, "id", e.AggregateId())
		switch v := e.(type) {
		default:
			panic(fmt.Sprintf("DomainEvent type %T not handled", v))
		case Created:
			shape := BDDShape{id: e.AggregateId(), nature: e.(Created).Nature, area: e.(Created).Area}
			s.query.Save(shape)
		case Stretched:
			shape := s.query.GetById(e.AggregateId())
			shape.area = e.(Stretched).Area
			s.query.Save(shape)
		}
	}
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

func (b BDDQueryShape) GetById(id uuid.UUID) BDDShape {
	if b.shapes[id] == (BDDShape{}) {
		panic(fmt.Sprintf("id %s not found", id))
	}
	return b.shapes[id]
}

func (b BDDQueryShape) GetByNature(nature string) uuid.UUID {
	for _, value := range b.shapes {
		if value.nature == nature {
			return value.id
		}
	}
	panic(fmt.Sprintf("shape of nature %s not found", nature))
}

func (b BDDQueryShape) GetAll() []BDDShape {
	var values []BDDShape
	for _, value := range b.shapes {
		values = append(values, value)
	}
	return values
}
