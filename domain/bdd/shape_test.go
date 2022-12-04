package main

import (
	"context"
	"example2/domain/commands"
	"example2/domain/valueobject"
	"fmt"
	"github.com/beorn7/floats"
	"github.com/cucumber/godog"
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
	repository := ctx.Value("repository").(valueobject.InMemoryRepository)
	area := repository.Get(0).GetArea()
	f, _ := strconv.ParseFloat(arg1, 32)
	if !floats.AlmostEqual(float64(area), f, 0.01) {
		return fmt.Errorf("expected %f, found %f", f, area)
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
	factory := valueobject.NewFactory()
	command, err := factory.NewCreationShapeCommand(nature, dimensions...)

	ctx = executeShapeCommand(ctx, command)

	return ctx, err
}
func executeShapeCommand(ctx context.Context, command commands.Command) context.Context {
	repository := valueobject.InMemoryRepository{}
	valueobject.
		NewShapeCreationCommandHandler(&repository).
		Execute(command)
	ctx = context.WithValue(ctx, "repository", repository)
	return ctx
}
