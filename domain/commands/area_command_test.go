package commands

import (
	mock "example2/domain/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ExecutionCommandTestSuite struct {
	suite.Suite
	mockShape mock.MockShape
}

func TestCommandTestSuite(t *testing.T) {
	suite.Run(t, new(ExecutionCommandTestSuite))
}

func (suite *ExecutionCommandTestSuite) BeforeTest(suiteName, testName string) {
	suite.mockShape = mock.CreateAMockShape()
}

func (suite *ExecutionCommandTestSuite) AfterTest(suiteName, testName string) {
	suite.mockShape.AssertExpectations(suite.T())
}

func (suite *ExecutionCommandTestSuite) TestCommandShape() {
	aeraCommand := suite.createAnAreaCommand()
	var err = aeraCommand.Execute()
	assert.NoError(suite.T(), err)

}

func (suite *ExecutionCommandTestSuite) createAnAreaCommand() Command {
	return areaCommand{shape: &suite.mockShape}
}
