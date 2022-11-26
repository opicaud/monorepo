package commands

import (
	"container/list"
	"example2/domain/valueobject"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestHandlerACommand(t *testing.T) {
	fakeRepository := FakeRepository{}
	command := &FakeCommand{}
	command.Mock.On("Execute").Return(nil)
	handler := NewHandlerCommand(fakeRepository)

	err := handler.Handler(command)

	assert.NoError(t, err)
	command.Mock.AssertCalled(t, "Execute")
}

func TestAStandardHandlerACommand(t *testing.T) {
	fakeRepository := FakeRepository{list: list.New()}
	handler := NewHandlerCommand(fakeRepository)
	assert.IsType(t, FakeRepository{}, handler.(*standardCommandHandler).repository)
}

func (f FakeRepository) Save(shape valueobject.Shape) error {
	f.list.PushFront(shape)
	return nil
}

func (f FakeRepository) AssertContains(t *testing.T, shape valueobject.Shape) bool {
	return assert.Contains(t, f.list, shape)
}

type FakeCommand struct {
	Mock mock.Mock
}
type FakeRepository struct {
	list *list.List
}

func (f *FakeCommand) Execute() error {
	args := f.Mock.Called()
	return args.Error(0)
}
