// Code generated by mockery v1.1.0. DO NOT EDIT.

package common

import mock "github.com/stretchr/testify/mock"

// mockJobTraceIsMaskingURLParams is an autogenerated mock type for the jobTraceIsMaskingURLParams type
type mockJobTraceIsMaskingURLParams struct {
	mock.Mock
}

// IsMaskingURLParams provides a mock function with given fields:
func (_m *mockJobTraceIsMaskingURLParams) IsMaskingURLParams() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
