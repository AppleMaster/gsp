// Code generated by counterfeiter. DO NOT EDIT.
package awsfakes

import (
	"context"
	"sync"

	"github.com/alphagov/gsp/components/service-operator/internal/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

type FakeAWSClient struct {
	CreateStackWithContextStub        func(context.Context, *cloudformation.CreateStackInput, ...request.Option) (*cloudformation.CreateStackOutput, error)
	createStackWithContextMutex       sync.RWMutex
	createStackWithContextArgsForCall []struct {
		arg1 context.Context
		arg2 *cloudformation.CreateStackInput
		arg3 []request.Option
	}
	createStackWithContextReturns struct {
		result1 *cloudformation.CreateStackOutput
		result2 error
	}
	createStackWithContextReturnsOnCall map[int]struct {
		result1 *cloudformation.CreateStackOutput
		result2 error
	}
	DeleteStackWithContextStub        func(context.Context, *cloudformation.DeleteStackInput, ...request.Option) (*cloudformation.DeleteStackOutput, error)
	deleteStackWithContextMutex       sync.RWMutex
	deleteStackWithContextArgsForCall []struct {
		arg1 context.Context
		arg2 *cloudformation.DeleteStackInput
		arg3 []request.Option
	}
	deleteStackWithContextReturns struct {
		result1 *cloudformation.DeleteStackOutput
		result2 error
	}
	deleteStackWithContextReturnsOnCall map[int]struct {
		result1 *cloudformation.DeleteStackOutput
		result2 error
	}
	DescribeStackEventsWithContextStub        func(context.Context, *cloudformation.DescribeStackEventsInput, ...request.Option) (*cloudformation.DescribeStackEventsOutput, error)
	describeStackEventsWithContextMutex       sync.RWMutex
	describeStackEventsWithContextArgsForCall []struct {
		arg1 context.Context
		arg2 *cloudformation.DescribeStackEventsInput
		arg3 []request.Option
	}
	describeStackEventsWithContextReturns struct {
		result1 *cloudformation.DescribeStackEventsOutput
		result2 error
	}
	describeStackEventsWithContextReturnsOnCall map[int]struct {
		result1 *cloudformation.DescribeStackEventsOutput
		result2 error
	}
	DescribeStacksWithContextStub        func(context.Context, *cloudformation.DescribeStacksInput, ...request.Option) (*cloudformation.DescribeStacksOutput, error)
	describeStacksWithContextMutex       sync.RWMutex
	describeStacksWithContextArgsForCall []struct {
		arg1 context.Context
		arg2 *cloudformation.DescribeStacksInput
		arg3 []request.Option
	}
	describeStacksWithContextReturns struct {
		result1 *cloudformation.DescribeStacksOutput
		result2 error
	}
	describeStacksWithContextReturnsOnCall map[int]struct {
		result1 *cloudformation.DescribeStacksOutput
		result2 error
	}
	UpdateStackWithContextStub        func(context.Context, *cloudformation.UpdateStackInput, ...request.Option) (*cloudformation.UpdateStackOutput, error)
	updateStackWithContextMutex       sync.RWMutex
	updateStackWithContextArgsForCall []struct {
		arg1 context.Context
		arg2 *cloudformation.UpdateStackInput
		arg3 []request.Option
	}
	updateStackWithContextReturns struct {
		result1 *cloudformation.UpdateStackOutput
		result2 error
	}
	updateStackWithContextReturnsOnCall map[int]struct {
		result1 *cloudformation.UpdateStackOutput
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAWSClient) CreateStackWithContext(arg1 context.Context, arg2 *cloudformation.CreateStackInput, arg3 ...request.Option) (*cloudformation.CreateStackOutput, error) {
	fake.createStackWithContextMutex.Lock()
	ret, specificReturn := fake.createStackWithContextReturnsOnCall[len(fake.createStackWithContextArgsForCall)]
	fake.createStackWithContextArgsForCall = append(fake.createStackWithContextArgsForCall, struct {
		arg1 context.Context
		arg2 *cloudformation.CreateStackInput
		arg3 []request.Option
	}{arg1, arg2, arg3})
	fake.recordInvocation("CreateStackWithContext", []interface{}{arg1, arg2, arg3})
	fake.createStackWithContextMutex.Unlock()
	if fake.CreateStackWithContextStub != nil {
		return fake.CreateStackWithContextStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createStackWithContextReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAWSClient) CreateStackWithContextCallCount() int {
	fake.createStackWithContextMutex.RLock()
	defer fake.createStackWithContextMutex.RUnlock()
	return len(fake.createStackWithContextArgsForCall)
}

func (fake *FakeAWSClient) CreateStackWithContextCalls(stub func(context.Context, *cloudformation.CreateStackInput, ...request.Option) (*cloudformation.CreateStackOutput, error)) {
	fake.createStackWithContextMutex.Lock()
	defer fake.createStackWithContextMutex.Unlock()
	fake.CreateStackWithContextStub = stub
}

func (fake *FakeAWSClient) CreateStackWithContextArgsForCall(i int) (context.Context, *cloudformation.CreateStackInput, []request.Option) {
	fake.createStackWithContextMutex.RLock()
	defer fake.createStackWithContextMutex.RUnlock()
	argsForCall := fake.createStackWithContextArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeAWSClient) CreateStackWithContextReturns(result1 *cloudformation.CreateStackOutput, result2 error) {
	fake.createStackWithContextMutex.Lock()
	defer fake.createStackWithContextMutex.Unlock()
	fake.CreateStackWithContextStub = nil
	fake.createStackWithContextReturns = struct {
		result1 *cloudformation.CreateStackOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeAWSClient) CreateStackWithContextReturnsOnCall(i int, result1 *cloudformation.CreateStackOutput, result2 error) {
	fake.createStackWithContextMutex.Lock()
	defer fake.createStackWithContextMutex.Unlock()
	fake.CreateStackWithContextStub = nil
	if fake.createStackWithContextReturnsOnCall == nil {
		fake.createStackWithContextReturnsOnCall = make(map[int]struct {
			result1 *cloudformation.CreateStackOutput
			result2 error
		})
	}
	fake.createStackWithContextReturnsOnCall[i] = struct {
		result1 *cloudformation.CreateStackOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeAWSClient) DeleteStackWithContext(arg1 context.Context, arg2 *cloudformation.DeleteStackInput, arg3 ...request.Option) (*cloudformation.DeleteStackOutput, error) {
	fake.deleteStackWithContextMutex.Lock()
	ret, specificReturn := fake.deleteStackWithContextReturnsOnCall[len(fake.deleteStackWithContextArgsForCall)]
	fake.deleteStackWithContextArgsForCall = append(fake.deleteStackWithContextArgsForCall, struct {
		arg1 context.Context
		arg2 *cloudformation.DeleteStackInput
		arg3 []request.Option
	}{arg1, arg2, arg3})
	fake.recordInvocation("DeleteStackWithContext", []interface{}{arg1, arg2, arg3})
	fake.deleteStackWithContextMutex.Unlock()
	if fake.DeleteStackWithContextStub != nil {
		return fake.DeleteStackWithContextStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.deleteStackWithContextReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAWSClient) DeleteStackWithContextCallCount() int {
	fake.deleteStackWithContextMutex.RLock()
	defer fake.deleteStackWithContextMutex.RUnlock()
	return len(fake.deleteStackWithContextArgsForCall)
}

func (fake *FakeAWSClient) DeleteStackWithContextCalls(stub func(context.Context, *cloudformation.DeleteStackInput, ...request.Option) (*cloudformation.DeleteStackOutput, error)) {
	fake.deleteStackWithContextMutex.Lock()
	defer fake.deleteStackWithContextMutex.Unlock()
	fake.DeleteStackWithContextStub = stub
}

func (fake *FakeAWSClient) DeleteStackWithContextArgsForCall(i int) (context.Context, *cloudformation.DeleteStackInput, []request.Option) {
	fake.deleteStackWithContextMutex.RLock()
	defer fake.deleteStackWithContextMutex.RUnlock()
	argsForCall := fake.deleteStackWithContextArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeAWSClient) DeleteStackWithContextReturns(result1 *cloudformation.DeleteStackOutput, result2 error) {
	fake.deleteStackWithContextMutex.Lock()
	defer fake.deleteStackWithContextMutex.Unlock()
	fake.DeleteStackWithContextStub = nil
	fake.deleteStackWithContextReturns = struct {
		result1 *cloudformation.DeleteStackOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeAWSClient) DeleteStackWithContextReturnsOnCall(i int, result1 *cloudformation.DeleteStackOutput, result2 error) {
	fake.deleteStackWithContextMutex.Lock()
	defer fake.deleteStackWithContextMutex.Unlock()
	fake.DeleteStackWithContextStub = nil
	if fake.deleteStackWithContextReturnsOnCall == nil {
		fake.deleteStackWithContextReturnsOnCall = make(map[int]struct {
			result1 *cloudformation.DeleteStackOutput
			result2 error
		})
	}
	fake.deleteStackWithContextReturnsOnCall[i] = struct {
		result1 *cloudformation.DeleteStackOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeAWSClient) DescribeStackEventsWithContext(arg1 context.Context, arg2 *cloudformation.DescribeStackEventsInput, arg3 ...request.Option) (*cloudformation.DescribeStackEventsOutput, error) {
	fake.describeStackEventsWithContextMutex.Lock()
	ret, specificReturn := fake.describeStackEventsWithContextReturnsOnCall[len(fake.describeStackEventsWithContextArgsForCall)]
	fake.describeStackEventsWithContextArgsForCall = append(fake.describeStackEventsWithContextArgsForCall, struct {
		arg1 context.Context
		arg2 *cloudformation.DescribeStackEventsInput
		arg3 []request.Option
	}{arg1, arg2, arg3})
	fake.recordInvocation("DescribeStackEventsWithContext", []interface{}{arg1, arg2, arg3})
	fake.describeStackEventsWithContextMutex.Unlock()
	if fake.DescribeStackEventsWithContextStub != nil {
		return fake.DescribeStackEventsWithContextStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.describeStackEventsWithContextReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAWSClient) DescribeStackEventsWithContextCallCount() int {
	fake.describeStackEventsWithContextMutex.RLock()
	defer fake.describeStackEventsWithContextMutex.RUnlock()
	return len(fake.describeStackEventsWithContextArgsForCall)
}

func (fake *FakeAWSClient) DescribeStackEventsWithContextCalls(stub func(context.Context, *cloudformation.DescribeStackEventsInput, ...request.Option) (*cloudformation.DescribeStackEventsOutput, error)) {
	fake.describeStackEventsWithContextMutex.Lock()
	defer fake.describeStackEventsWithContextMutex.Unlock()
	fake.DescribeStackEventsWithContextStub = stub
}

func (fake *FakeAWSClient) DescribeStackEventsWithContextArgsForCall(i int) (context.Context, *cloudformation.DescribeStackEventsInput, []request.Option) {
	fake.describeStackEventsWithContextMutex.RLock()
	defer fake.describeStackEventsWithContextMutex.RUnlock()
	argsForCall := fake.describeStackEventsWithContextArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeAWSClient) DescribeStackEventsWithContextReturns(result1 *cloudformation.DescribeStackEventsOutput, result2 error) {
	fake.describeStackEventsWithContextMutex.Lock()
	defer fake.describeStackEventsWithContextMutex.Unlock()
	fake.DescribeStackEventsWithContextStub = nil
	fake.describeStackEventsWithContextReturns = struct {
		result1 *cloudformation.DescribeStackEventsOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeAWSClient) DescribeStackEventsWithContextReturnsOnCall(i int, result1 *cloudformation.DescribeStackEventsOutput, result2 error) {
	fake.describeStackEventsWithContextMutex.Lock()
	defer fake.describeStackEventsWithContextMutex.Unlock()
	fake.DescribeStackEventsWithContextStub = nil
	if fake.describeStackEventsWithContextReturnsOnCall == nil {
		fake.describeStackEventsWithContextReturnsOnCall = make(map[int]struct {
			result1 *cloudformation.DescribeStackEventsOutput
			result2 error
		})
	}
	fake.describeStackEventsWithContextReturnsOnCall[i] = struct {
		result1 *cloudformation.DescribeStackEventsOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeAWSClient) DescribeStacksWithContext(arg1 context.Context, arg2 *cloudformation.DescribeStacksInput, arg3 ...request.Option) (*cloudformation.DescribeStacksOutput, error) {
	fake.describeStacksWithContextMutex.Lock()
	ret, specificReturn := fake.describeStacksWithContextReturnsOnCall[len(fake.describeStacksWithContextArgsForCall)]
	fake.describeStacksWithContextArgsForCall = append(fake.describeStacksWithContextArgsForCall, struct {
		arg1 context.Context
		arg2 *cloudformation.DescribeStacksInput
		arg3 []request.Option
	}{arg1, arg2, arg3})
	fake.recordInvocation("DescribeStacksWithContext", []interface{}{arg1, arg2, arg3})
	fake.describeStacksWithContextMutex.Unlock()
	if fake.DescribeStacksWithContextStub != nil {
		return fake.DescribeStacksWithContextStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.describeStacksWithContextReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAWSClient) DescribeStacksWithContextCallCount() int {
	fake.describeStacksWithContextMutex.RLock()
	defer fake.describeStacksWithContextMutex.RUnlock()
	return len(fake.describeStacksWithContextArgsForCall)
}

func (fake *FakeAWSClient) DescribeStacksWithContextCalls(stub func(context.Context, *cloudformation.DescribeStacksInput, ...request.Option) (*cloudformation.DescribeStacksOutput, error)) {
	fake.describeStacksWithContextMutex.Lock()
	defer fake.describeStacksWithContextMutex.Unlock()
	fake.DescribeStacksWithContextStub = stub
}

func (fake *FakeAWSClient) DescribeStacksWithContextArgsForCall(i int) (context.Context, *cloudformation.DescribeStacksInput, []request.Option) {
	fake.describeStacksWithContextMutex.RLock()
	defer fake.describeStacksWithContextMutex.RUnlock()
	argsForCall := fake.describeStacksWithContextArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeAWSClient) DescribeStacksWithContextReturns(result1 *cloudformation.DescribeStacksOutput, result2 error) {
	fake.describeStacksWithContextMutex.Lock()
	defer fake.describeStacksWithContextMutex.Unlock()
	fake.DescribeStacksWithContextStub = nil
	fake.describeStacksWithContextReturns = struct {
		result1 *cloudformation.DescribeStacksOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeAWSClient) DescribeStacksWithContextReturnsOnCall(i int, result1 *cloudformation.DescribeStacksOutput, result2 error) {
	fake.describeStacksWithContextMutex.Lock()
	defer fake.describeStacksWithContextMutex.Unlock()
	fake.DescribeStacksWithContextStub = nil
	if fake.describeStacksWithContextReturnsOnCall == nil {
		fake.describeStacksWithContextReturnsOnCall = make(map[int]struct {
			result1 *cloudformation.DescribeStacksOutput
			result2 error
		})
	}
	fake.describeStacksWithContextReturnsOnCall[i] = struct {
		result1 *cloudformation.DescribeStacksOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeAWSClient) UpdateStackWithContext(arg1 context.Context, arg2 *cloudformation.UpdateStackInput, arg3 ...request.Option) (*cloudformation.UpdateStackOutput, error) {
	fake.updateStackWithContextMutex.Lock()
	ret, specificReturn := fake.updateStackWithContextReturnsOnCall[len(fake.updateStackWithContextArgsForCall)]
	fake.updateStackWithContextArgsForCall = append(fake.updateStackWithContextArgsForCall, struct {
		arg1 context.Context
		arg2 *cloudformation.UpdateStackInput
		arg3 []request.Option
	}{arg1, arg2, arg3})
	fake.recordInvocation("UpdateStackWithContext", []interface{}{arg1, arg2, arg3})
	fake.updateStackWithContextMutex.Unlock()
	if fake.UpdateStackWithContextStub != nil {
		return fake.UpdateStackWithContextStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateStackWithContextReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAWSClient) UpdateStackWithContextCallCount() int {
	fake.updateStackWithContextMutex.RLock()
	defer fake.updateStackWithContextMutex.RUnlock()
	return len(fake.updateStackWithContextArgsForCall)
}

func (fake *FakeAWSClient) UpdateStackWithContextCalls(stub func(context.Context, *cloudformation.UpdateStackInput, ...request.Option) (*cloudformation.UpdateStackOutput, error)) {
	fake.updateStackWithContextMutex.Lock()
	defer fake.updateStackWithContextMutex.Unlock()
	fake.UpdateStackWithContextStub = stub
}

func (fake *FakeAWSClient) UpdateStackWithContextArgsForCall(i int) (context.Context, *cloudformation.UpdateStackInput, []request.Option) {
	fake.updateStackWithContextMutex.RLock()
	defer fake.updateStackWithContextMutex.RUnlock()
	argsForCall := fake.updateStackWithContextArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeAWSClient) UpdateStackWithContextReturns(result1 *cloudformation.UpdateStackOutput, result2 error) {
	fake.updateStackWithContextMutex.Lock()
	defer fake.updateStackWithContextMutex.Unlock()
	fake.UpdateStackWithContextStub = nil
	fake.updateStackWithContextReturns = struct {
		result1 *cloudformation.UpdateStackOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeAWSClient) UpdateStackWithContextReturnsOnCall(i int, result1 *cloudformation.UpdateStackOutput, result2 error) {
	fake.updateStackWithContextMutex.Lock()
	defer fake.updateStackWithContextMutex.Unlock()
	fake.UpdateStackWithContextStub = nil
	if fake.updateStackWithContextReturnsOnCall == nil {
		fake.updateStackWithContextReturnsOnCall = make(map[int]struct {
			result1 *cloudformation.UpdateStackOutput
			result2 error
		})
	}
	fake.updateStackWithContextReturnsOnCall[i] = struct {
		result1 *cloudformation.UpdateStackOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeAWSClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createStackWithContextMutex.RLock()
	defer fake.createStackWithContextMutex.RUnlock()
	fake.deleteStackWithContextMutex.RLock()
	defer fake.deleteStackWithContextMutex.RUnlock()
	fake.describeStackEventsWithContextMutex.RLock()
	defer fake.describeStackEventsWithContextMutex.RUnlock()
	fake.describeStacksWithContextMutex.RLock()
	defer fake.describeStacksWithContextMutex.RUnlock()
	fake.updateStackWithContextMutex.RLock()
	defer fake.updateStackWithContextMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAWSClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ aws.AWSClient = new(FakeAWSClient)
