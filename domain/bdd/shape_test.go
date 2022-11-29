package main

import (
	"context"
	"errors"
	"example2/domain/commands/factory"
	"example2/domain/commands/fullshapecommand"
	"example2/domain/utils"
	"fmt"
	"github.com/cucumber/godog"
	"testing"
)

type TestContext struct {
	length float32
	width  float32
}

func iCreateARectangle(ctx context.Context) (context.Context, error) {

	testContext := ctx.Value("testContext").(TestContext)
	factory := factory.NewFactory()
	command, _ := factory.CreateAFullShapeCommand("rectangle", testContext.length, testContext.width)

	repository := utils.FakeRepository{}
	handler := fullshapecommand.NewFullShapeCommandHandler(&repository)
	handler.Execute(command)
	ctx = context.WithValue(ctx, "repository", repository)

	return ctx, nil
}

func itAreaIs(ctx context.Context, arg1 int) error {
	repository := ctx.Value("repository").(utils.FakeRepository)
	shape := repository.Get(0)
	area := shape.GetArea()
	if area != float32(arg1) {
		return errors.New(fmt.Errorf("expected %f, found %f", float32(arg1), area).Error())
	}
	return nil
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
