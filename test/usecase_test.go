package test

import (
	"sync"
	"testing"

	"github.com/EBANNO/DEMO/test/mocks"
	"github.com/stretchr/testify/mock"
)

func Test_Process(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	// Create mocks
	entityRepositoryMock := new(mocks.EntityRepository)
	entityRepositoryMock.On("Create").Return(nil).Run(func(args mock.Arguments) {
		wg.Done()
	}).Once()

	// Invoke method
	usecase := New(entityRepositoryMock)
	usecase.Process()
	wg.Wait()

	// Check results
	entityRepositoryMock.AssertNumberOfCalls(t, "Create", 1)
}
