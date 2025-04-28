package new

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type Notifier interface {
	SendNotification(message string) error
}

type MockNotifier struct {
	mock.Mock
}

func (m *MockNotifier) SendNotification(message string) error {
	args := m.Called(message)
	return args.Error(0)
}

func NotifyUser(notifier Notifier, message string) error {
	return notifier.SendNotification(message)
}

func TestNotifyUser(t *testing.T) {
	mockNotifier := new(MockNotifier)
	mockNotifier.On("SendNotification", "Hello!").Return(nil)

	err := NotifyUser(mockNotifier, "Hello!")
	mockNotifier.AssertExpectations(t)
	assert.NoError(t, err)
}
