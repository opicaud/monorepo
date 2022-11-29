package main

import (
	"context"
	"example2/domain/commands"
	"example2/domain/commands/factory"
	"example2/domain/commands/fullshapecommand"
	"example2/domain/utils"
	"github.com/cucumber/godog"
	"testing"
)

type TestContext struct {
	length float32
	width  float32
}

func iCreateARectangle(ctx context.Context) (context.Context, error) {
	factory := factory.NewFactory()
	testContext := ctx.Value("testContext").(TestContext)
	command, _ := factory.CreateAFullShapeCommand("rectangle", testContext.length, testContext.width)
	ctx = context.WithValue(ctx, "command", command)
	return ctx, nil
}

func itAreaIs(ctx context.Context, arg1 int) (context.Context, error) {
	repository := &utils.FakeRepository{}
	command, _ := ctx.Value("command").(commands.Command)
	handler := fullshapecommand.NewFullShapeCommandHandler(repository)
	handler.Execute(command)

	return ctx, nil
}

func lengthOfAndWidthOf(ctx context.Context, arg1 int, arg2 int) (context.Context, error) {
	ctx = context.Background()
	testContext := TestContext{length: float32(arg1), width: float32(arg2)}
	ctx = context.WithValue(ctx, "testContext", testContext)
	return ctx, nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I create a rectangle$`, iCreateARectangle)
	ctx.Step(`^it area is (\d+)$`, itAreaIs)
	ctx.Step(`^length of (\d+) and width of (\d+)$`, lengthOfAndWidthOf)
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
