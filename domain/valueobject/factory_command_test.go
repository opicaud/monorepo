package valueobject

import (
	"example2/domain/commands"
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
	var r, _ commands.Command = newCreationShapeCommand("a-shape", []float32{1, 2})
	var r2 error = nil
	command, err := r, r2
	assert.NotNil(suite.T(), command)
	assert.NoError(suite.T(), err)
}
