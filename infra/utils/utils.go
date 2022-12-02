package utils

import "github.com/stretchr/testify/mock"

type FakeEventBus struct {
	Mock mock.Mock
}

func (f *FakeEventBus) NotifyAll() {
	f.Mock.Called()
}

func NewFakeEventBus() *FakeEventBus {
	f := FakeEventBus{}
	f.Mock.On("NotifyAll").Return(nil)
	return &f
}
