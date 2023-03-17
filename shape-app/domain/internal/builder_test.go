package internal

import (
	"github.com/google/uuid"
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
	_, err := newShapeBuilder().withNature("a-unknown-shape").withId(uuid.New())
	assert.Error(suite.T(), err)
}

func (suite *BuilderAggregateTestSuite) TestCreateARectangleShapeWithId() {
	shape, err := newShapeBuilder().withNature("rectangle").withId(uuid.New())
	assert.NoError(suite.T(), err)
	assert.IsType(suite.T(), &rectangle{}, shape)
	assert.IsType(suite.T(), uuid.UUID{}, shape.(*rectangle).id)
}
