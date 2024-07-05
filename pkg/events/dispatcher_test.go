package events

import (
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	name    string
	payload any
}

func (e *TestEvent) GetName() string {
	return e.name
}

func (e *TestEvent) GetPayload() any {
	return e.payload
}

func (e *TestEvent) SetPayload(payload any) {
	e.payload = payload
}

func (e *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

type EventHandlerMock struct {
	mock.Mock
}

func (m *EventHandlerMock) Handle(event EventInterface) {
	m.Called(event)
}

type TestEventHandler struct {
	id uint
}

func (m *TestEventHandler) Handle(event EventInterface) {}

type EventDispatcherTestSuite struct {
	suite.Suite
	event      *TestEvent
	event2     *TestEvent
	handler    *TestEventHandler
	handler2   *TestEventHandler
	handler3   *TestEventHandler
	dispatcher *Dispatcher
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}

func (s *EventDispatcherTestSuite) SetupTest() {
	s.dispatcher = NewEventDispatcher()
	s.handler = &TestEventHandler{id: 1}
	s.handler2 = &TestEventHandler{id: 2}
	s.handler3 = &TestEventHandler{id: 3}
	s.event = &TestEvent{name: "event", payload: "event"}
	s.event2 = &TestEvent{name: "event2", payload: "event2"}
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	err := s.dispatcher.Register(s.event.GetName(), s.handler)
	s.Nil(err)
	s.Equal(1, len(s.dispatcher.handlers[s.event.GetName()]))

	err = s.dispatcher.Register(s.event.GetName(), s.handler2)
	s.Nil(err)
	s.Equal(2, len(s.dispatcher.handlers[s.event.GetName()]))

	s.Equal(s.handler, s.dispatcher.handlers[s.event.GetName()][0])
	s.Equal(s.handler2, s.dispatcher.handlers[s.event.GetName()][1])
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Register_WithSameHandler() {
	err := s.dispatcher.Register(s.event.GetName(), s.handler)
	s.Nil(err)
	s.Equal(1, len(s.dispatcher.handlers[s.event.GetName()]))

	err = s.dispatcher.Register(s.event.GetName(), s.handler)
	s.Equal(ErrHandlerAlreadyRegistered, err)
	s.Equal(1, len(s.dispatcher.handlers[s.event.GetName()]))
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Dispatch() {
	eh := &EventHandlerMock{}
	eh.On("Handle", mock.Anything).Times(1)

	err := s.dispatcher.Register(s.event.GetName(), eh)
	s.Nil(err)
	err = s.dispatcher.Dispatch(s.event)
	s.Nil(err)

	eh.AssertExpectations(s.T())
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Remove() {
	err := s.dispatcher.Register(s.event.GetName(), s.handler)
	s.Nil(err)
	err = s.dispatcher.Register(s.event.GetName(), s.handler2)
	s.Nil(err)
	err = s.dispatcher.Register(s.event.GetName(), s.handler3)
	s.Nil(err)

	err = s.dispatcher.Remove(s.event.GetName(), s.handler3)
	s.Nil(err)
	s.Equal(2, len(s.dispatcher.handlers[s.event.GetName()]))
	s.False(s.dispatcher.Has(s.event.GetName(), s.handler3))
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Has() {
	err := s.dispatcher.Register(s.event.GetName(), s.handler)
	s.Nil(err)
	err = s.dispatcher.Register(s.event.GetName(), s.handler2)
	s.Nil(err)

	s.True(s.dispatcher.Has(s.event.GetName(), s.handler))
	s.True(s.dispatcher.Has(s.event.GetName(), s.handler2))
	s.False(s.dispatcher.Has(s.event2.GetName(), s.handler3))
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Clear() {
	// event 1
	err := s.dispatcher.Register(s.event.GetName(), s.handler)
	s.Nil(err)

	err = s.dispatcher.Register(s.event.GetName(), s.handler2)
	s.Nil(err)

	// event 2
	err = s.dispatcher.Register(s.event2.GetName(), s.handler3)
	s.Nil(err)

	s.dispatcher.Clear()
	s.Equal(0, len(s.dispatcher.handlers))
}
