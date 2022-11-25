package commands

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CommandTestSuite struct {
	suite.Suite
	ShapeMock MockShape
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(CommandTestSuite))
}

func (suite *CommandTestSuite) BeforeTest(suiteName, testName string) {
	suite.createAMockShape()
}

func (suite *CommandTestSuite) AfterTest(suiteName, testName string) {
	suite.ShapeMock.AssertExpectations(suite.T())
}

func (suite *CommandTestSuite) TestCommandShape() {
	aeraCommand := suite.createAnAreaCommand()
	var err = aeraCommand.Execute()
	assert.NoError(suite.T(), err)

}

func (suite *CommandTestSuite) createAnAreaCommand() Command {
	return AreaCommand{shape: &suite.ShapeMock}
}

func (suite *CommandTestSuite) createAMockShape() {
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
