package commands

import (
	"example2/domain/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type FactoryTestSuite struct {
	suite.Suite
}

func TestFactoryTestSuite(t *testing.T) {
	suite.Run(t, new(FactoryTestSuite))
}

func (suite *FactoryTestSuite) TestCreateACommandFullShape() {
	factory := NewFactoryWithCustomBuilder(&utils.FakeShapeBuilder{})
	command, err := factory.CreateAFullShapeCommand("a-shape", 1, 2)
	assert.NotNil(suite.T(), command)
	assert.NoError(suite.T(), err)
}
