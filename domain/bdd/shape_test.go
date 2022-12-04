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

var (
	query   = &BDDQueryShape{shapes: make(map[uuid.UUID]BDDShape)}
	factory = aggregate.NewFactory()
)

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

func anExistingRectangle(ctx context.Context) (context.Context, error) {
	shape := query.GetByNature("rectangle")
	ctx = context.WithValue(ctx, "shape", shape)
	return ctx, nil
}

func iStetchItBy(ctx context.Context, arg1 int) error {
	shape := ctx.Value("shape").(BDDShape)
	makeStretchCommand(ctx, shape.id, float32(arg1))
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I create a rectangle$`, iCreateARectangle)
	ctx.Step(`^it area is "([^"]*)"$`, itAreaIs)
	ctx.Step(`^length of (\d+) and width of (\d+)$`, lengthOfAndWidthOf)
	ctx.Step(`^I create a circle$`, iCreateACircle)
	ctx.Step(`^radius of (\d+)$`, radiusOf)
	ctx.Step(`^an existing rectangle$`, anExistingRectangle)
	ctx.Step(`^I stetch it by (\d+)$`, iStetchItBy)
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
func makeShapeCommand(ctx context.Context, nature string, dimensions ...float32) (context.Context, error) {
	command, err := factory.NewCreationShapeCommand(nature, dimensions...)
	ctx = executeShapeCommand(ctx, command)
	return ctx, err
}

func makeStretchCommand(ctx context.Context, id uuid.UUID, stretchBy float32) context.Context {
	command := factory.NewStretchShapeCommand(id, stretchBy)
	ctx = executeShapeCommand(ctx, command)
	return ctx
}

func executeShapeCommand(ctx context.Context, command commands.Command) context.Context {
	aggregate.NewShapeCreationCommandHandlerBuilder().
		WithSubscriber(&Subscriber{ctx: &ctx, query: query}).
		Build().Execute(command)
	return ctx
}

type Subscriber struct {
	ctx   *context.Context
	query QueryShapeModel
}

func (s *Subscriber) Update(events []infra.Event) {
	for _, e := range events {
		*s.ctx = context.WithValue(*s.ctx, "id", e.AggregateId())
		switch v := e.(type) {
		default:
			panic(fmt.Sprintf("Event type %T not handled", v))
		case aggregate.ShapeCreatedEvent:
			shape := BDDShape{id: e.AggregateId(), nature: e.(aggregate.ShapeCreatedEvent).Nature}
			s.query.Save(shape)
		case aggregate.AreaShapeCalculated:
			shape := s.query.GetById(e.AggregateId())
			shape.area = e.(aggregate.AreaShapeCalculated).Area
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
	GetByNature(nature string) BDDShape
	GetAll() []BDDShape
}

type BDDQueryShape struct {
	shapes map[uuid.UUID]BDDShape
}

func (b *BDDQueryShape) Save(shape BDDShape) {
	b.shapes[shape.id] = shape
}

func (b BDDQueryShape) GetById(id uuid.UUID) BDDShape {
	return b.shapes[id]
}

func (b BDDQueryShape) GetByNature(nature string) BDDShape {
	for _, value := range b.shapes {
		if value.nature == nature {
			return value
		}
	}
	panic(fmt.Sprintf("shape of nature %s not found", nature))
}

func (b BDDQueryShape) GetAll() []BDDShape {
	values := []BDDShape{}
	for _, value := range b.shapes {
		values = append(values, value)
	}
	return values
}
