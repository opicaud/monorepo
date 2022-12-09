package shape

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type FactoryTestSuite struct {
	suite.Suite
}

func TestFactoryTestSuite2(t *testing.T) {
	suite.Run(t, new(FactoryTestSuite))
}

func (suite *FactoryTestSuite) TestCreateACommandFullShape() {
	var command = newCreationShapeCommand("a-shape", []float32{1, 2})
	assert.NotNil(suite.T(), command)
}
