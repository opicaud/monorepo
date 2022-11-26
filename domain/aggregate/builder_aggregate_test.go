package aggregate

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type BuilderAggregateTestSuite struct {
	suite.Suite
}

func TestFactoryTestSuite(t *testing.T) {
	suite.Run(t, new(BuilderAggregateTestSuite))
}

func (suite *BuilderAggregateTestSuite) TestCreateAUnknownShape() {
	_, err := NewShapeBuilder().CreateAShape("a-unknown-shape").WithDimensions(0)
	assert.Error(suite.T(), err)
}

func (suite *BuilderAggregateTestSuite) TestCreateARectangleShape() {
	shape, err := NewShapeBuilder().CreateAShape("rectangle").WithDimensions(1, 2)
	assert.NoError(suite.T(), err)
	assert.IsType(suite.T(), &rectangle{}, shape)
}
