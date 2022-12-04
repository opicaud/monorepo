package main

import (
	"context"
	"example2/domain/aggregate"
	"example2/domain/commands"
	"example2/infra"
	"fmt"
	"github.com/beorn7/floats"
	"github.com/cucumber/godog"
	"github.com/google/uuid"
	"strconv"
	"testing"
)

type TestContext struct {
	length float32
	width  float32
	radius float32
}

func iCreateARectangle(ctx context.Context) (context.Context, error) {
	testContext := ctx.Value("testContext").(TestContext)
	return makeShapeCommand(ctx, "rectangle", testContext.length, testContext.width)
}

func iCreateACircle(ctx context.Context) (context.Context, error) {
	testContext := ctx.Value("testContext").(TestContext)
	return makeShapeCommand(ctx, "circle", testContext.radius)

}

func itAreaIs(ctx context.Context, arg1 string) error {
	id := ctx.Value("id").(uuid.UUID)
	newArea := ctx.Value("query").(QueryShapeModel).Get(id).area
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

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I create a rectangle$`, iCreateARectangle)
	ctx.Step(`^it area is "([^"]*)"$`, itAreaIs)
	ctx.Step(`^length of (\d+) and width of (\d+)$`, lengthOfAndWidthOf)
	ctx.Step(`^I create a circle$`, iCreateACircle)
	ctx.Step(`^radius of (\d+)$`, radiusOf)

}

func TestFeatures(t *testing.T) {

	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
func makeShapeCommand(ctx context.Context, nature string, dimensions ...float32) (context.Context, error) {
	factory := aggregate.NewFactory()
	command, err := factory.NewCreationShapeCommand(nature, dimensions...)

	ctx = executeShapeCommand(ctx, command)

	return ctx, err
}

func executeShapeCommand(ctx context.Context, command commands.Command) context.Context {
	aggregate.NewShapeCreationCommandHandlerBuilder().
		WithSubscriber(&Subscriber{ctx: &ctx, query: &BDDQueryShape{shapes: make(map[uuid.UUID]BDDShape)}}).
		Build().Execute(command)
	return ctx
}

type Subscriber struct {
	ctx   *context.Context
	query QueryShapeModel
}

func (s *Subscriber) Update(events []infra.Event) {
	for _, e := range events {
		switch v := e.(type) {
		default:
			panic(fmt.Sprintf("Event type %T not handled", v))
		case aggregate.ShapeCreatedEvent:
			shape := BDDShape{id: e.AggregateId()}
			s.query.Save(shape)
		case aggregate.AreaShapeCalculated:
			shape := s.query.Get(e.AggregateId())
			shape.area = e.(aggregate.AreaShapeCalculated).Area
			s.query.Save(shape)
			*s.ctx = context.WithValue(*s.ctx, "id", e.AggregateId())
		}
	}
	*s.ctx = context.WithValue(*s.ctx, "query", s.query)
}

type BDDShape struct {
	id   uuid.UUID
	area float32
}

type QueryShapeModel interface {
	Save(shape BDDShape)
	Get(id uuid.UUID) BDDShape
	GetAll() []BDDShape
}

type BDDQueryShape struct {
	shapes map[uuid.UUID]BDDShape
}

func (b *BDDQueryShape) Save(shape BDDShape) {
	b.shapes[shape.id] = shape
}

func (b BDDQueryShape) Get(id uuid.UUID) BDDShape {
	return b.shapes[id]
}

func (b BDDQueryShape) GetAll() []BDDShape {
	values := []BDDShape{}
	for _, value := range b.shapes {
		values = append(values, value)
	}
	return values
}
