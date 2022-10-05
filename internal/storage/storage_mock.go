// Code generated by mockery v2.14.0. DO NOT EDIT.

package storage

import (
	mock "github.com/stretchr/testify/mock"
	pagination "github.com/systemli/ticker/internal/api/pagination"
)

// MockTickerStorage is an autogenerated mock type for the TickerStorage type
type MockTickerStorage struct {
	mock.Mock
}

// AddUsersToTicker provides a mock function with given fields: ticker, ids
func (_m *MockTickerStorage) AddUsersToTicker(ticker Ticker, ids []int) error {
	ret := _m.Called(ticker, ids)

	var r0 error
	if rf, ok := ret.Get(0).(func(Ticker, []int) error); ok {
		r0 = rf(ticker, ids)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CountUser provides a mock function with given fields:
func (_m *MockTickerStorage) CountUser() (int, error) {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteMessage provides a mock function with given fields: message
func (_m *MockTickerStorage) DeleteMessage(message Message) error {
	ret := _m.Called(message)

	var r0 error
	if rf, ok := ret.Get(0).(func(Message) error); ok {
		r0 = rf(message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteMessages provides a mock function with given fields: ticker
func (_m *MockTickerStorage) DeleteMessages(ticker Ticker) error {
	ret := _m.Called(ticker)

	var r0 error
	if rf, ok := ret.Get(0).(func(Ticker) error); ok {
		r0 = rf(ticker)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTicker provides a mock function with given fields: ticker
func (_m *MockTickerStorage) DeleteTicker(ticker Ticker) error {
	ret := _m.Called(ticker)

	var r0 error
	if rf, ok := ret.Get(0).(func(Ticker) error); ok {
		r0 = rf(ticker)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUpload provides a mock function with given fields: upload
func (_m *MockTickerStorage) DeleteUpload(upload Upload) error {
	ret := _m.Called(upload)

	var r0 error
	if rf, ok := ret.Get(0).(func(Upload) error); ok {
		r0 = rf(upload)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUploads provides a mock function with given fields: uploads
func (_m *MockTickerStorage) DeleteUploads(uploads []Upload) {
	_m.Called(uploads)
}

// DeleteUploadsByTicker provides a mock function with given fields: ticker
func (_m *MockTickerStorage) DeleteUploadsByTicker(ticker Ticker) error {
	ret := _m.Called(ticker)

	var r0 error
	if rf, ok := ret.Get(0).(func(Ticker) error); ok {
		r0 = rf(ticker)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUser provides a mock function with given fields: user
func (_m *MockTickerStorage) DeleteUser(user User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindMessage provides a mock function with given fields: tickerID, messageID
func (_m *MockTickerStorage) FindMessage(tickerID int, messageID int) (Message, error) {
	ret := _m.Called(tickerID, messageID)

	var r0 Message
	if rf, ok := ret.Get(0).(func(int, int) Message); ok {
		r0 = rf(tickerID, messageID)
	} else {
		r0 = ret.Get(0).(Message)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(tickerID, messageID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindMessagesByTicker provides a mock function with given fields: ticker
func (_m *MockTickerStorage) FindMessagesByTicker(ticker Ticker) ([]Message, error) {
	ret := _m.Called(ticker)

	var r0 []Message
	if rf, ok := ret.Get(0).(func(Ticker) []Message); ok {
		r0 = rf(ticker)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(Ticker) error); ok {
		r1 = rf(ticker)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindMessagesByTickerAndPagination provides a mock function with given fields: ticker, _a1
func (_m *MockTickerStorage) FindMessagesByTickerAndPagination(ticker Ticker, _a1 pagination.Pagination) ([]Message, error) {
	ret := _m.Called(ticker, _a1)

	var r0 []Message
	if rf, ok := ret.Get(0).(func(Ticker, pagination.Pagination) []Message); ok {
		r0 = rf(ticker, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(Ticker, pagination.Pagination) error); ok {
		r1 = rf(ticker, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindSetting provides a mock function with given fields: name
func (_m *MockTickerStorage) FindSetting(name string) (Setting, error) {
	ret := _m.Called(name)

	var r0 Setting
	if rf, ok := ret.Get(0).(func(string) Setting); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(Setting)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTickerByDomain provides a mock function with given fields: domain
func (_m *MockTickerStorage) FindTickerByDomain(domain string) (Ticker, error) {
	ret := _m.Called(domain)

	var r0 Ticker
	if rf, ok := ret.Get(0).(func(string) Ticker); ok {
		r0 = rf(domain)
	} else {
		r0 = ret.Get(0).(Ticker)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTickerByID provides a mock function with given fields: id
func (_m *MockTickerStorage) FindTickerByID(id int) (Ticker, error) {
	ret := _m.Called(id)

	var r0 Ticker
	if rf, ok := ret.Get(0).(func(int) Ticker); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(Ticker)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTickers provides a mock function with given fields:
func (_m *MockTickerStorage) FindTickers() ([]Ticker, error) {
	ret := _m.Called()

	var r0 []Ticker
	if rf, ok := ret.Get(0).(func() []Ticker); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Ticker)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTickersByIDs provides a mock function with given fields: ids
func (_m *MockTickerStorage) FindTickersByIDs(ids []int) ([]Ticker, error) {
	ret := _m.Called(ids)

	var r0 []Ticker
	if rf, ok := ret.Get(0).(func([]int) []Ticker); ok {
		r0 = rf(ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Ticker)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]int) error); ok {
		r1 = rf(ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUploadByUUID provides a mock function with given fields: uuid
func (_m *MockTickerStorage) FindUploadByUUID(uuid string) (Upload, error) {
	ret := _m.Called(uuid)

	var r0 Upload
	if rf, ok := ret.Get(0).(func(string) Upload); ok {
		r0 = rf(uuid)
	} else {
		r0 = ret.Get(0).(Upload)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(uuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUploadsByIDs provides a mock function with given fields: ids
func (_m *MockTickerStorage) FindUploadsByIDs(ids []int) ([]Upload, error) {
	ret := _m.Called(ids)

	var r0 []Upload
	if rf, ok := ret.Get(0).(func([]int) []Upload); ok {
		r0 = rf(ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Upload)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]int) error); ok {
		r1 = rf(ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUploadsByMessage provides a mock function with given fields: message
func (_m *MockTickerStorage) FindUploadsByMessage(message Message) []Upload {
	ret := _m.Called(message)

	var r0 []Upload
	if rf, ok := ret.Get(0).(func(Message) []Upload); ok {
		r0 = rf(message)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Upload)
		}
	}

	return r0
}

// FindUserByEmail provides a mock function with given fields: email
func (_m *MockTickerStorage) FindUserByEmail(email string) (User, error) {
	ret := _m.Called(email)

	var r0 User
	if rf, ok := ret.Get(0).(func(string) User); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUserByID provides a mock function with given fields: id
func (_m *MockTickerStorage) FindUserByID(id int) (User, error) {
	ret := _m.Called(id)

	var r0 User
	if rf, ok := ret.Get(0).(func(int) User); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUsers provides a mock function with given fields:
func (_m *MockTickerStorage) FindUsers() ([]User, error) {
	ret := _m.Called()

	var r0 []User
	if rf, ok := ret.Get(0).(func() []User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUsersByTicker provides a mock function with given fields: ticker
func (_m *MockTickerStorage) FindUsersByTicker(ticker Ticker) ([]User, error) {
	ret := _m.Called(ticker)

	var r0 []User
	if rf, ok := ret.Get(0).(func(Ticker) []User); ok {
		r0 = rf(ticker)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(Ticker) error); ok {
		r1 = rf(ticker)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInactiveSetting provides a mock function with given fields:
func (_m *MockTickerStorage) GetInactiveSetting() Setting {
	ret := _m.Called()

	var r0 Setting
	if rf, ok := ret.Get(0).(func() Setting); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(Setting)
	}

	return r0
}

// GetRefreshIntervalSetting provides a mock function with given fields:
func (_m *MockTickerStorage) GetRefreshIntervalSetting() Setting {
	ret := _m.Called()

	var r0 Setting
	if rf, ok := ret.Get(0).(func() Setting); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(Setting)
	}

	return r0
}

// GetRefreshIntervalSettingValue provides a mock function with given fields:
func (_m *MockTickerStorage) GetRefreshIntervalSettingValue() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// RemoveTickerFromUser provides a mock function with given fields: ticker, user
func (_m *MockTickerStorage) RemoveTickerFromUser(ticker Ticker, user User) error {
	ret := _m.Called(ticker, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(Ticker, User) error); ok {
		r0 = rf(ticker, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveInactiveSetting provides a mock function with given fields: inactiveSettings
func (_m *MockTickerStorage) SaveInactiveSetting(inactiveSettings InactiveSettings) error {
	ret := _m.Called(inactiveSettings)

	var r0 error
	if rf, ok := ret.Get(0).(func(InactiveSettings) error); ok {
		r0 = rf(inactiveSettings)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveMessage provides a mock function with given fields: message
func (_m *MockTickerStorage) SaveMessage(message *Message) error {
	ret := _m.Called(message)

	var r0 error
	if rf, ok := ret.Get(0).(func(*Message) error); ok {
		r0 = rf(message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveRefreshInterval provides a mock function with given fields: refreshInterval
func (_m *MockTickerStorage) SaveRefreshInterval(refreshInterval int) error {
	ret := _m.Called(refreshInterval)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(refreshInterval)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveTicker provides a mock function with given fields: ticker
func (_m *MockTickerStorage) SaveTicker(ticker *Ticker) error {
	ret := _m.Called(ticker)

	var r0 error
	if rf, ok := ret.Get(0).(func(*Ticker) error); ok {
		r0 = rf(ticker)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveUpload provides a mock function with given fields: upload
func (_m *MockTickerStorage) SaveUpload(upload *Upload) error {
	ret := _m.Called(upload)

	var r0 error
	if rf, ok := ret.Get(0).(func(*Upload) error); ok {
		r0 = rf(upload)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveUser provides a mock function with given fields: user
func (_m *MockTickerStorage) SaveUser(user *User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(*User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UploadPath provides a mock function with given fields:
func (_m *MockTickerStorage) UploadPath() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewMockTickerStorage interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockTickerStorage creates a new instance of MockTickerStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockTickerStorage(t mockConstructorTestingTNewMockTickerStorage) *MockTickerStorage {
	mock := &MockTickerStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
