package valueobject

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
	_, err := newShapeBuilder().CreateAShape("a-unknown-shape").WithDimensions([]float32{2, 3})
	assert.Error(suite.T(), err)
}

func (suite *BuilderAggregateTestSuite) TestCreateARectangleShape() {
	dimensions := []float32{2, 3}
	shape, err := newShapeBuilder().CreateAShape("rectangle").WithDimensions(dimensions)
	assert.NoError(suite.T(), err)
	assert.IsType(suite.T(), &rectangle{}, shape)
	assert.Equal(suite.T(), dimensions[0], shape.(*rectangle).length)
	assert.Equal(suite.T(), dimensions[1], shape.(*rectangle).width)

}

func (suite *BuilderAggregateTestSuite) TestCreateACircleShape() {
	dimensions := []float32{2}
	shape, err := newShapeBuilder().CreateAShape("circle").WithDimensions(dimensions)
	assert.NoError(suite.T(), err)
	assert.IsType(suite.T(), &circle{}, shape)
	assert.Equal(suite.T(), dimensions[0], shape.(*circle).radius)

}
