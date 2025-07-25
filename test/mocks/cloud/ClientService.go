// Code generated by mockery v2.52.3. DO NOT EDIT.

package cloud

import (
	clientcloud "github.com/metal-stack/ontap-go/api/client/cloud"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"
)

// ClientService is an autogenerated mock type for the ClientService type
type ClientService struct {
	mock.Mock
}

// CloudTargetCollectionGet provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) CloudTargetCollectionGet(params *clientcloud.CloudTargetCollectionGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientcloud.ClientOption) (*clientcloud.CloudTargetCollectionGetOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CloudTargetCollectionGet")
	}

	var r0 *clientcloud.CloudTargetCollectionGetOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientcloud.CloudTargetCollectionGetParams, runtime.ClientAuthInfoWriter, ...clientcloud.ClientOption) (*clientcloud.CloudTargetCollectionGetOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientcloud.CloudTargetCollectionGetParams, runtime.ClientAuthInfoWriter, ...clientcloud.ClientOption) *clientcloud.CloudTargetCollectionGetOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientcloud.CloudTargetCollectionGetOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientcloud.CloudTargetCollectionGetParams, runtime.ClientAuthInfoWriter, ...clientcloud.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CloudTargetCreate provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) CloudTargetCreate(params *clientcloud.CloudTargetCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientcloud.ClientOption) (*clientcloud.CloudTargetCreateCreated, *clientcloud.CloudTargetCreateAccepted, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CloudTargetCreate")
	}

	var r0 *clientcloud.CloudTargetCreateCreated
	var r1 *clientcloud.CloudTargetCreateAccepted
	var r2 error
	if rf, ok := ret.Get(0).(func(*clientcloud.CloudTargetCreateParams, runtime.ClientAuthInfoWriter, ...clientcloud.ClientOption) (*clientcloud.CloudTargetCreateCreated, *clientcloud.CloudTargetCreateAccepted, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientcloud.CloudTargetCreateParams, runtime.ClientAuthInfoWriter, ...clientcloud.ClientOption) *clientcloud.CloudTargetCreateCreated); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientcloud.CloudTargetCreateCreated)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientcloud.CloudTargetCreateParams, runtime.ClientAuthInfoWriter, ...clientcloud.ClientOption) *clientcloud.CloudTargetCreateAccepted); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*clientcloud.CloudTargetCreateAccepted)
		}
	}

	if rf, ok := ret.Get(2).(func(*clientcloud.CloudTargetCreateParams, runtime.ClientAuthInfoWriter, ...clientcloud.ClientOption) error); ok {
		r2 = rf(params, authInfo, opts...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CloudTargetDelete provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) CloudTargetDelete(params *clientcloud.CloudTargetDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientcloud.ClientOption) (*clientcloud.CloudTargetDeleteOK, *clientcloud.CloudTargetDeleteAccepted, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CloudTargetDelete")
	}

	var r0 *clientcloud.CloudTargetDeleteOK
	var r1 *clientcloud.CloudTargetDeleteAccepted
	var r2 error
	if rf, ok := ret.Get(0).(func(*clientcloud.CloudTargetDeleteParams, runtime.ClientAuthInfoWriter, ...clientcloud.ClientOption) (*clientcloud.CloudTargetDeleteOK, *clientcloud.CloudTargetDeleteAccepted, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientcloud.CloudTargetDeleteParams, runtime.ClientAuthInfoWriter, ...clientcloud.ClientOption) *clientcloud.CloudTargetDeleteOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientcloud.CloudTargetDeleteOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientcloud.CloudTargetDeleteParams, runtime.ClientAuthInfoWriter, ...clientcloud.ClientOption) *clientcloud.CloudTargetDeleteAccepted); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*clientcloud.CloudTargetDeleteAccepted)
		}
	}

	if rf, ok := ret.Get(2).(func(*clientcloud.CloudTargetDeleteParams, runtime.ClientAuthInfoWriter, ...clientcloud.ClientOption) error); ok {
		r2 = rf(params, authInfo, opts...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CloudTargetGet provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) CloudTargetGet(params *clientcloud.CloudTargetGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientcloud.ClientOption) (*clientcloud.CloudTargetGetOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CloudTargetGet")
	}

	var r0 *clientcloud.CloudTargetGetOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientcloud.CloudTargetGetParams, runtime.ClientAuthInfoWriter, ...clientcloud.ClientOption) (*clientcloud.CloudTargetGetOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientcloud.CloudTargetGetParams, runtime.ClientAuthInfoWriter, ...clientcloud.ClientOption) *clientcloud.CloudTargetGetOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientcloud.CloudTargetGetOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientcloud.CloudTargetGetParams, runtime.ClientAuthInfoWriter, ...clientcloud.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CloudTargetModify provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) CloudTargetModify(params *clientcloud.CloudTargetModifyParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientcloud.ClientOption) (*clientcloud.CloudTargetModifyOK, *clientcloud.CloudTargetModifyAccepted, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CloudTargetModify")
	}

	var r0 *clientcloud.CloudTargetModifyOK
	var r1 *clientcloud.CloudTargetModifyAccepted
	var r2 error
	if rf, ok := ret.Get(0).(func(*clientcloud.CloudTargetModifyParams, runtime.ClientAuthInfoWriter, ...clientcloud.ClientOption) (*clientcloud.CloudTargetModifyOK, *clientcloud.CloudTargetModifyAccepted, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientcloud.CloudTargetModifyParams, runtime.ClientAuthInfoWriter, ...clientcloud.ClientOption) *clientcloud.CloudTargetModifyOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientcloud.CloudTargetModifyOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientcloud.CloudTargetModifyParams, runtime.ClientAuthInfoWriter, ...clientcloud.ClientOption) *clientcloud.CloudTargetModifyAccepted); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*clientcloud.CloudTargetModifyAccepted)
		}
	}

	if rf, ok := ret.Get(2).(func(*clientcloud.CloudTargetModifyParams, runtime.ClientAuthInfoWriter, ...clientcloud.ClientOption) error); ok {
		r2 = rf(params, authInfo, opts...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SetTransport provides a mock function with given fields: transport
func (_m *ClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}

// NewClientService creates a new instance of ClientService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClientService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ClientService {
	mock := &ClientService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
