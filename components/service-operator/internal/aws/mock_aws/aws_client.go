// Code generated by MockGen. DO NOT EDIT.
// Source: internal/aws/aws_client.go

// Package mock_aws is a generated GoMock package.
package mock_aws

import (
	aws "github.com/aws/aws-sdk-go/aws"
	request "github.com/aws/aws-sdk-go/aws/request"
	cloudformation "github.com/aws/aws-sdk-go/service/cloudformation"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAWSClient is a mock of AWSClient interface
type MockAWSClient struct {
	ctrl     *gomock.Controller
	recorder *MockAWSClientMockRecorder
}

// MockAWSClientMockRecorder is the mock recorder for MockAWSClient
type MockAWSClientMockRecorder struct {
	mock *MockAWSClient
}

// NewMockAWSClient creates a new mock instance
func NewMockAWSClient(ctrl *gomock.Controller) *MockAWSClient {
	mock := &MockAWSClient{ctrl: ctrl}
	mock.recorder = &MockAWSClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAWSClient) EXPECT() *MockAWSClientMockRecorder {
	return m.recorder
}

// DescribeStacksWithContext mocks base method
func (m *MockAWSClient) DescribeStacksWithContext(arg0 aws.Context, arg1 *cloudformation.DescribeStacksInput, arg2 ...request.Option) (*cloudformation.DescribeStacksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeStacksWithContext", varargs...)
	ret0, _ := ret[0].(*cloudformation.DescribeStacksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeStacksWithContext indicates an expected call of DescribeStacksWithContext
func (mr *MockAWSClientMockRecorder) DescribeStacksWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeStacksWithContext", reflect.TypeOf((*MockAWSClient)(nil).DescribeStacksWithContext), varargs...)
}

// DescribeStackEventsWithContext mocks base method
func (m *MockAWSClient) DescribeStackEventsWithContext(arg0 aws.Context, arg1 *cloudformation.DescribeStackEventsInput, arg2 ...request.Option) (*cloudformation.DescribeStackEventsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeStackEventsWithContext", varargs...)
	ret0, _ := ret[0].(*cloudformation.DescribeStackEventsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeStackEventsWithContext indicates an expected call of DescribeStackEventsWithContext
func (mr *MockAWSClientMockRecorder) DescribeStackEventsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeStackEventsWithContext", reflect.TypeOf((*MockAWSClient)(nil).DescribeStackEventsWithContext), varargs...)
}

// CreateStackWithContext mocks base method
func (m *MockAWSClient) CreateStackWithContext(arg0 aws.Context, arg1 *cloudformation.CreateStackInput, arg2 ...request.Option) (*cloudformation.CreateStackOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateStackWithContext", varargs...)
	ret0, _ := ret[0].(*cloudformation.CreateStackOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStackWithContext indicates an expected call of CreateStackWithContext
func (mr *MockAWSClientMockRecorder) CreateStackWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStackWithContext", reflect.TypeOf((*MockAWSClient)(nil).CreateStackWithContext), varargs...)
}

// UpdateStackWithContext mocks base method
func (m *MockAWSClient) UpdateStackWithContext(arg0 aws.Context, arg1 *cloudformation.UpdateStackInput, arg2 ...request.Option) (*cloudformation.UpdateStackOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateStackWithContext", varargs...)
	ret0, _ := ret[0].(*cloudformation.UpdateStackOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStackWithContext indicates an expected call of UpdateStackWithContext
func (mr *MockAWSClientMockRecorder) UpdateStackWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStackWithContext", reflect.TypeOf((*MockAWSClient)(nil).UpdateStackWithContext), varargs...)
}

// DeleteStackWithContext mocks base method
func (m *MockAWSClient) DeleteStackWithContext(arg0 aws.Context, arg1 *cloudformation.DeleteStackInput, arg2 ...request.Option) (*cloudformation.DeleteStackOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteStackWithContext", varargs...)
	ret0, _ := ret[0].(*cloudformation.DeleteStackOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteStackWithContext indicates an expected call of DeleteStackWithContext
func (mr *MockAWSClientMockRecorder) DeleteStackWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStackWithContext", reflect.TypeOf((*MockAWSClient)(nil).DeleteStackWithContext), varargs...)
}
