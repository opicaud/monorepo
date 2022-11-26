package factories

import (
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
	factory := New()
	command, err := factory.CreateAFullShapeCommand("a-shape", 1, 2)
	assert.NotNil(suite.T(), command)
	assert.NoError(suite.T(), err)
}
