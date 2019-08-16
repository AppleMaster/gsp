// Code generated by MockGen. DO NOT EDIT.
// Source: ../cloudformation_controller.go

// Package aws_mocks is a generated GoMock package.
package aws_mocks

import (
	context "context"
	reflect "reflect"

	internal "github.com/alphagov/gsp/components/service-operator/internal"
	aws "github.com/alphagov/gsp/components/service-operator/internal/aws"
	cloudformation "github.com/aws/aws-sdk-go/service/cloudformation"
	cloudformation0 "github.com/awslabs/goformation/cloudformation"
	resources "github.com/awslabs/goformation/cloudformation/resources"
	logr "github.com/go-logr/logr"
	gomock "github.com/golang/mock/gomock"
	controller_runtime "sigs.k8s.io/controller-runtime"
)

// MockCloudFormationTemplate is a mock of CloudFormationTemplate interface
type MockCloudFormationTemplate struct {
	ctrl     *gomock.Controller
	recorder *MockCloudFormationTemplateMockRecorder
}

// MockCloudFormationTemplateMockRecorder is the mock recorder for MockCloudFormationTemplate
type MockCloudFormationTemplateMockRecorder struct {
	mock *MockCloudFormationTemplate
}

// NewMockCloudFormationTemplate creates a new mock instance
func NewMockCloudFormationTemplate(ctrl *gomock.Controller) *MockCloudFormationTemplate {
	mock := &MockCloudFormationTemplate{ctrl: ctrl}
	mock.recorder = &MockCloudFormationTemplateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCloudFormationTemplate) EXPECT() *MockCloudFormationTemplateMockRecorder {
	return m.recorder
}

// Template mocks base method
func (m *MockCloudFormationTemplate) Template(arg0 string, arg1 []resources.Tag) *cloudformation0.Template {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Template", arg0, arg1)
	ret0, _ := ret[0].(*cloudformation0.Template)
	return ret0
}

// Template indicates an expected call of Template
func (mr *MockCloudFormationTemplateMockRecorder) Template(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Template", reflect.TypeOf((*MockCloudFormationTemplate)(nil).Template), arg0, arg1)
}

// Parameters mocks base method
func (m *MockCloudFormationTemplate) Parameters() ([]*cloudformation.Parameter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parameters")
	ret0, _ := ret[0].([]*cloudformation.Parameter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Parameters indicates an expected call of Parameters
func (mr *MockCloudFormationTemplateMockRecorder) Parameters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parameters", reflect.TypeOf((*MockCloudFormationTemplate)(nil).Parameters))
}

// ResourceType mocks base method
func (m *MockCloudFormationTemplate) ResourceType() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceType")
	ret0, _ := ret[0].(string)
	return ret0
}

// ResourceType indicates an expected call of ResourceType
func (mr *MockCloudFormationTemplateMockRecorder) ResourceType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceType", reflect.TypeOf((*MockCloudFormationTemplate)(nil).ResourceType))
}

// MockCloudFormationReconciler is a mock of CloudFormationReconciler interface
type MockCloudFormationReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockCloudFormationReconcilerMockRecorder
}

// MockCloudFormationReconcilerMockRecorder is the mock recorder for MockCloudFormationReconciler
type MockCloudFormationReconcilerMockRecorder struct {
	mock *MockCloudFormationReconciler
}

// NewMockCloudFormationReconciler creates a new mock instance
func NewMockCloudFormationReconciler(ctrl *gomock.Controller) *MockCloudFormationReconciler {
	mock := &MockCloudFormationReconciler{ctrl: ctrl}
	mock.recorder = &MockCloudFormationReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCloudFormationReconciler) EXPECT() *MockCloudFormationReconcilerMockRecorder {
	return m.recorder
}

// Reconcile mocks base method
func (m *MockCloudFormationReconciler) Reconcile(arg0 context.Context, arg1 logr.Logger, arg2 controller_runtime.Request, arg3 aws.CloudFormationTemplate, arg4 bool) (internal.Action, aws.StackData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reconcile", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(internal.Action)
	ret1, _ := ret[1].(aws.StackData)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Reconcile indicates an expected call of Reconcile
func (mr *MockCloudFormationReconcilerMockRecorder) Reconcile(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reconcile", reflect.TypeOf((*MockCloudFormationReconciler)(nil).Reconcile), arg0, arg1, arg2, arg3, arg4)
}
