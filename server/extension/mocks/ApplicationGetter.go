// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package mocks

import (
	"github.com/argoproj/argo-cd/v3/pkg/apis/application/v1alpha1"
	mock "github.com/stretchr/testify/mock"
)

// NewApplicationGetter creates a new instance of ApplicationGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewApplicationGetter(t interface {
	mock.TestingT
	Cleanup(func())
}) *ApplicationGetter {
	mock := &ApplicationGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// ApplicationGetter is an autogenerated mock type for the ApplicationGetter type
type ApplicationGetter struct {
	mock.Mock
}

type ApplicationGetter_Expecter struct {
	mock *mock.Mock
}

func (_m *ApplicationGetter) EXPECT() *ApplicationGetter_Expecter {
	return &ApplicationGetter_Expecter{mock: &_m.Mock}
}

// Get provides a mock function for the type ApplicationGetter
func (_mock *ApplicationGetter) Get(ns string, name string) (*v1alpha1.Application, error) {
	ret := _mock.Called(ns, name)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *v1alpha1.Application
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(string, string) (*v1alpha1.Application, error)); ok {
		return returnFunc(ns, name)
	}
	if returnFunc, ok := ret.Get(0).(func(string, string) *v1alpha1.Application); ok {
		r0 = returnFunc(ns, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.Application)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = returnFunc(ns, name)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// ApplicationGetter_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type ApplicationGetter_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ns string
//   - name string
func (_e *ApplicationGetter_Expecter) Get(ns interface{}, name interface{}) *ApplicationGetter_Get_Call {
	return &ApplicationGetter_Get_Call{Call: _e.mock.On("Get", ns, name)}
}

func (_c *ApplicationGetter_Get_Call) Run(run func(ns string, name string)) *ApplicationGetter_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 string
		if args[0] != nil {
			arg0 = args[0].(string)
		}
		var arg1 string
		if args[1] != nil {
			arg1 = args[1].(string)
		}
		run(
			arg0,
			arg1,
		)
	})
	return _c
}

func (_c *ApplicationGetter_Get_Call) Return(application *v1alpha1.Application, err error) *ApplicationGetter_Get_Call {
	_c.Call.Return(application, err)
	return _c
}

func (_c *ApplicationGetter_Get_Call) RunAndReturn(run func(ns string, name string) (*v1alpha1.Application, error)) *ApplicationGetter_Get_Call {
	_c.Call.Return(run)
	return _c
}
