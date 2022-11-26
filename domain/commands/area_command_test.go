package commands

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ExecutionCommandTestSuite struct {
	suite.Suite
	ShapeMock MockShape
}

func TestCommandTestSuite(t *testing.T) {
	suite.Run(t, new(ExecutionCommandTestSuite))
}

func (suite *ExecutionCommandTestSuite) BeforeTest(suiteName, testName string) {
	suite.createAMockShape()
}

func (suite *ExecutionCommandTestSuite) AfterTest(suiteName, testName string) {
	suite.ShapeMock.AssertExpectations(suite.T())
}

func (suite *ExecutionCommandTestSuite) TestCommandShape() {
	aeraCommand := suite.createAnAreaCommand()
	var err = aeraCommand.Execute()
	assert.NoError(suite.T(), err)

}

func (suite *ExecutionCommandTestSuite) createAnAreaCommand() Command {
	return areaCommand{shape: &suite.ShapeMock}
}

func (suite *ExecutionCommandTestSuite) createAMockShape() {
	suite.ShapeMock = MockShape{}
	suite.ShapeMock.On("Area").Return(nil)
}

type MockShape struct {
	mock.Mock
}

func (c *MockShape) Area() error {
	args := c.Called()
	return args.Error(0)
}
