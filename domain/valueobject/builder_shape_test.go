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
	_, _, err := newShapeBuilder().createAShape("a-unknown-shape").withDimensions([]float32{2, 3})
	assert.Error(suite.T(), err)
}

func (suite *BuilderAggregateTestSuite) TestCreateARectangleShape() {
	dimensions := []float32{2, 3}
	shape, event, err := newShapeBuilder().createAShape("rectangle").withDimensions(dimensions)
	assert.NoError(suite.T(), err)
	assert.IsType(suite.T(), &rectangle{}, shape)
	assert.Equal(suite.T(), dimensions[0], shape.(*rectangle).length)
	assert.Equal(suite.T(), dimensions[1], shape.(*rectangle).width)
	assert.IsType(suite.T(), event, ShapeCreatedEvent{})
	assert.Equal(suite.T(), "rectangle", event.(ShapeCreatedEvent).nature)
	assert.Equal(suite.T(), dimensions, event.(ShapeCreatedEvent).dimensions)

}

func (suite *BuilderAggregateTestSuite) TestCreateACircleShape() {
	dimensions := []float32{2}
	shape, _, err := newShapeBuilder().createAShape("circle").withDimensions(dimensions)
	assert.NoError(suite.T(), err)
	assert.IsType(suite.T(), &circle{}, shape)
	assert.Equal(suite.T(), dimensions[0], shape.(*circle).radius)

}
