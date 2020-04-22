// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mattermost/mattermost-plugin-mscalendar/server/remote (interfaces: Client)

// Package mock_remote is a generated GoMock package.
package mock_remote

import (
	gomock "github.com/golang/mock/gomock"
	remote "github.com/mattermost/mattermost-plugin-mscalendar/server/remote"
	url "net/url"
	reflect "reflect"
	time "time"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// AcceptEvent mocks base method
func (m *MockClient) AcceptEvent(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptEvent", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AcceptEvent indicates an expected call of AcceptEvent
func (mr *MockClientMockRecorder) AcceptEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptEvent", reflect.TypeOf((*MockClient)(nil).AcceptEvent), arg0, arg1)
}

// CallFormPost mocks base method
func (m *MockClient) CallFormPost(arg0, arg1 string, arg2 url.Values, arg3 interface{}) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallFormPost", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CallFormPost indicates an expected call of CallFormPost
func (mr *MockClientMockRecorder) CallFormPost(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallFormPost", reflect.TypeOf((*MockClient)(nil).CallFormPost), arg0, arg1, arg2, arg3)
}

// CallJSON mocks base method
func (m *MockClient) CallJSON(arg0, arg1 string, arg2, arg3 interface{}) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallJSON", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CallJSON indicates an expected call of CallJSON
func (mr *MockClientMockRecorder) CallJSON(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallJSON", reflect.TypeOf((*MockClient)(nil).CallJSON), arg0, arg1, arg2, arg3)
}

// CreateCalendar mocks base method
func (m *MockClient) CreateCalendar(arg0 string, arg1 *remote.Calendar) (*remote.Calendar, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCalendar", arg0, arg1)
	ret0, _ := ret[0].(*remote.Calendar)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCalendar indicates an expected call of CreateCalendar
func (mr *MockClientMockRecorder) CreateCalendar(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCalendar", reflect.TypeOf((*MockClient)(nil).CreateCalendar), arg0, arg1)
}

// CreateEvent mocks base method
func (m *MockClient) CreateEvent(arg0 string, arg1 *remote.Event) (*remote.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEvent", arg0, arg1)
	ret0, _ := ret[0].(*remote.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEvent indicates an expected call of CreateEvent
func (mr *MockClientMockRecorder) CreateEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEvent", reflect.TypeOf((*MockClient)(nil).CreateEvent), arg0, arg1)
}

// CreateMySubscription mocks base method
func (m *MockClient) CreateMySubscription(arg0 string) (*remote.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMySubscription", arg0)
	ret0, _ := ret[0].(*remote.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMySubscription indicates an expected call of CreateMySubscription
func (mr *MockClientMockRecorder) CreateMySubscription(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMySubscription", reflect.TypeOf((*MockClient)(nil).CreateMySubscription), arg0)
}

// DeclineEvent mocks base method
func (m *MockClient) DeclineEvent(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeclineEvent", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeclineEvent indicates an expected call of DeclineEvent
func (mr *MockClientMockRecorder) DeclineEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeclineEvent", reflect.TypeOf((*MockClient)(nil).DeclineEvent), arg0, arg1)
}

// DeleteCalendar mocks base method
func (m *MockClient) DeleteCalendar(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCalendar", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCalendar indicates an expected call of DeleteCalendar
func (mr *MockClientMockRecorder) DeleteCalendar(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCalendar", reflect.TypeOf((*MockClient)(nil).DeleteCalendar), arg0, arg1)
}

// DeleteSubscription mocks base method
func (m *MockClient) DeleteSubscription(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSubscription", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSubscription indicates an expected call of DeleteSubscription
func (mr *MockClientMockRecorder) DeleteSubscription(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSubscription", reflect.TypeOf((*MockClient)(nil).DeleteSubscription), arg0)
}

// DoBatchViewCalendarRequests mocks base method
func (m *MockClient) DoBatchViewCalendarRequests(arg0 []*remote.ViewCalendarParams) ([]*remote.ViewCalendarResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoBatchViewCalendarRequests", arg0)
	ret0, _ := ret[0].([]*remote.ViewCalendarResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoBatchViewCalendarRequests indicates an expected call of DoBatchViewCalendarRequests
func (mr *MockClientMockRecorder) DoBatchViewCalendarRequests(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoBatchViewCalendarRequests", reflect.TypeOf((*MockClient)(nil).DoBatchViewCalendarRequests), arg0)
}

// FindMeetingTimes mocks base method
func (m *MockClient) FindMeetingTimes(arg0 string, arg1 *remote.FindMeetingTimesParameters) (*remote.MeetingTimeSuggestionResults, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindMeetingTimes", arg0, arg1)
	ret0, _ := ret[0].(*remote.MeetingTimeSuggestionResults)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindMeetingTimes indicates an expected call of FindMeetingTimes
func (mr *MockClientMockRecorder) FindMeetingTimes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindMeetingTimes", reflect.TypeOf((*MockClient)(nil).FindMeetingTimes), arg0, arg1)
}

// GetCalendars mocks base method
func (m *MockClient) GetCalendars(arg0 string) ([]*remote.Calendar, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCalendars", arg0)
	ret0, _ := ret[0].([]*remote.Calendar)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCalendars indicates an expected call of GetCalendars
func (mr *MockClientMockRecorder) GetCalendars(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCalendars", reflect.TypeOf((*MockClient)(nil).GetCalendars), arg0)
}

// GetDefaultCalendarView mocks base method
func (m *MockClient) GetDefaultCalendarView(arg0 string, arg1, arg2 time.Time) ([]*remote.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDefaultCalendarView", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*remote.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDefaultCalendarView indicates an expected call of GetDefaultCalendarView
func (mr *MockClientMockRecorder) GetDefaultCalendarView(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDefaultCalendarView", reflect.TypeOf((*MockClient)(nil).GetDefaultCalendarView), arg0, arg1, arg2)
}

// GetEvent mocks base method
func (m *MockClient) GetEvent(arg0, arg1 string) (*remote.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEvent", arg0, arg1)
	ret0, _ := ret[0].(*remote.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEvent indicates an expected call of GetEvent
func (mr *MockClientMockRecorder) GetEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEvent", reflect.TypeOf((*MockClient)(nil).GetEvent), arg0, arg1)
}

// GetMailboxSettings mocks base method
func (m *MockClient) GetMailboxSettings(arg0 string) (*remote.MailboxSettings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMailboxSettings", arg0)
	ret0, _ := ret[0].(*remote.MailboxSettings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMailboxSettings indicates an expected call of GetMailboxSettings
func (mr *MockClientMockRecorder) GetMailboxSettings(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMailboxSettings", reflect.TypeOf((*MockClient)(nil).GetMailboxSettings), arg0)
}

// GetMe mocks base method
func (m *MockClient) GetMe() (*remote.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMe")
	ret0, _ := ret[0].(*remote.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMe indicates an expected call of GetMe
func (mr *MockClientMockRecorder) GetMe() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMe", reflect.TypeOf((*MockClient)(nil).GetMe))
}

// GetNotificationData mocks base method
func (m *MockClient) GetNotificationData(arg0 *remote.Notification) (*remote.Notification, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNotificationData", arg0)
	ret0, _ := ret[0].(*remote.Notification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNotificationData indicates an expected call of GetNotificationData
func (mr *MockClientMockRecorder) GetNotificationData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNotificationData", reflect.TypeOf((*MockClient)(nil).GetNotificationData), arg0)
}

// GetSchedule mocks base method
func (m *MockClient) GetSchedule(arg0 []*remote.ScheduleUserInfo, arg1, arg2 *remote.DateTime, arg3 int) ([]*remote.ScheduleInformation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSchedule", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*remote.ScheduleInformation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSchedule indicates an expected call of GetSchedule
func (mr *MockClientMockRecorder) GetSchedule(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSchedule", reflect.TypeOf((*MockClient)(nil).GetSchedule), arg0, arg1, arg2, arg3)
}

// GetSuperuserToken mocks base method
func (m *MockClient) GetSuperuserToken() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSuperuserToken")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSuperuserToken indicates an expected call of GetSuperuserToken
func (mr *MockClientMockRecorder) GetSuperuserToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSuperuserToken", reflect.TypeOf((*MockClient)(nil).GetSuperuserToken))
}

// ListSubscriptions mocks base method
func (m *MockClient) ListSubscriptions() ([]*remote.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSubscriptions")
	ret0, _ := ret[0].([]*remote.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSubscriptions indicates an expected call of ListSubscriptions
func (mr *MockClientMockRecorder) ListSubscriptions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSubscriptions", reflect.TypeOf((*MockClient)(nil).ListSubscriptions))
}

// RenewSubscription mocks base method
func (m *MockClient) RenewSubscription(arg0 string) (*remote.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenewSubscription", arg0)
	ret0, _ := ret[0].(*remote.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RenewSubscription indicates an expected call of RenewSubscription
func (mr *MockClientMockRecorder) RenewSubscription(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenewSubscription", reflect.TypeOf((*MockClient)(nil).RenewSubscription), arg0)
}

// TentativelyAcceptEvent mocks base method
func (m *MockClient) TentativelyAcceptEvent(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TentativelyAcceptEvent", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// TentativelyAcceptEvent indicates an expected call of TentativelyAcceptEvent
func (mr *MockClientMockRecorder) TentativelyAcceptEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TentativelyAcceptEvent", reflect.TypeOf((*MockClient)(nil).TentativelyAcceptEvent), arg0, arg1)
}
