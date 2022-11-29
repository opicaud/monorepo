package main

import (
	"context"
	"errors"
	"example2/domain/commands/factory"
	"example2/domain/commands/shapecreationcommand"
	"example2/domain/utils"
	"fmt"
	"github.com/cucumber/godog"
	"strconv"
	"testing"
)

type TestContext struct {
	length float32
	width  float32
}

func iCreateARectangle(ctx context.Context) (context.Context, error) {

	testContext := ctx.Value("testContext").(TestContext)
	command, _ := factory.
		NewFactory().
		NewCreationShapeCommand("rectangle", testContext.length, testContext.width)

	repository := utils.FakeRepository{}
	shapecreationcommand.
		NewShapeCreationCommandHandler(&repository).
		Execute(command)
	ctx = context.WithValue(ctx, "repository", repository)

	return ctx, nil
}

func itAreaIs(ctx context.Context, arg1 string) error {
	repository := ctx.Value("repository").(utils.FakeRepository)
	area := repository.Get(0).GetArea()
	f, _ := strconv.ParseFloat(arg1, 32)
	if area != float32(f) {
		return errors.New(fmt.Errorf("expected %f, found %f", f, area).Error())
	}
	return nil
}

func lengthOfAndWidthOf(ctx context.Context, arg1 int, arg2 int) (context.Context, error) {
	ctx = context.Background()
	ctx = context.WithValue(ctx, "testContext", TestContext{length: float32(arg1), width: float32(arg2)})
	return ctx, nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I create a rectangle$`, iCreateARectangle)
	ctx.Step(`^it area is "([^"]*)"$`, itAreaIs)
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
