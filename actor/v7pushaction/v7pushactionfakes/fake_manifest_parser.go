// Code generated by counterfeiter. DO NOT EDIT.
package v7pushactionfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7pushaction"
	"code.cloudfoundry.org/cli/util/manifestparser"
)

type FakeManifestParser struct {
	AppsStub        func(string) ([]manifestparser.Application, error)
	appsMutex       sync.RWMutex
	appsArgsForCall []struct {
		arg1 string
	}
	appsReturns struct {
		result1 []manifestparser.Application
		result2 error
	}
	appsReturnsOnCall map[int]struct {
		result1 []manifestparser.Application
		result2 error
	}
	FullRawManifestStub        func() []byte
	fullRawManifestMutex       sync.RWMutex
	fullRawManifestArgsForCall []struct {
	}
	fullRawManifestReturns struct {
		result1 []byte
	}
	fullRawManifestReturnsOnCall map[int]struct {
		result1 []byte
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeManifestParser) Apps(arg1 string) ([]manifestparser.Application, error) {
	fake.appsMutex.Lock()
	ret, specificReturn := fake.appsReturnsOnCall[len(fake.appsArgsForCall)]
	fake.appsArgsForCall = append(fake.appsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Apps", []interface{}{arg1})
	fake.appsMutex.Unlock()
	if fake.AppsStub != nil {
		return fake.AppsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.appsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeManifestParser) AppsCallCount() int {
	fake.appsMutex.RLock()
	defer fake.appsMutex.RUnlock()
	return len(fake.appsArgsForCall)
}

func (fake *FakeManifestParser) AppsCalls(stub func(string) ([]manifestparser.Application, error)) {
	fake.appsMutex.Lock()
	defer fake.appsMutex.Unlock()
	fake.AppsStub = stub
}

func (fake *FakeManifestParser) AppsArgsForCall(i int) string {
	fake.appsMutex.RLock()
	defer fake.appsMutex.RUnlock()
	argsForCall := fake.appsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeManifestParser) AppsReturns(result1 []manifestparser.Application, result2 error) {
	fake.appsMutex.Lock()
	defer fake.appsMutex.Unlock()
	fake.AppsStub = nil
	fake.appsReturns = struct {
		result1 []manifestparser.Application
		result2 error
	}{result1, result2}
}

func (fake *FakeManifestParser) AppsReturnsOnCall(i int, result1 []manifestparser.Application, result2 error) {
	fake.appsMutex.Lock()
	defer fake.appsMutex.Unlock()
	fake.AppsStub = nil
	if fake.appsReturnsOnCall == nil {
		fake.appsReturnsOnCall = make(map[int]struct {
			result1 []manifestparser.Application
			result2 error
		})
	}
	fake.appsReturnsOnCall[i] = struct {
		result1 []manifestparser.Application
		result2 error
	}{result1, result2}
}

func (fake *FakeManifestParser) FullRawManifest() []byte {
	fake.fullRawManifestMutex.Lock()
	ret, specificReturn := fake.fullRawManifestReturnsOnCall[len(fake.fullRawManifestArgsForCall)]
	fake.fullRawManifestArgsForCall = append(fake.fullRawManifestArgsForCall, struct {
	}{})
	fake.recordInvocation("FullRawManifest", []interface{}{})
	fake.fullRawManifestMutex.Unlock()
	if fake.FullRawManifestStub != nil {
		return fake.FullRawManifestStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.fullRawManifestReturns
	return fakeReturns.result1
}

func (fake *FakeManifestParser) FullRawManifestCallCount() int {
	fake.fullRawManifestMutex.RLock()
	defer fake.fullRawManifestMutex.RUnlock()
	return len(fake.fullRawManifestArgsForCall)
}

func (fake *FakeManifestParser) FullRawManifestCalls(stub func() []byte) {
	fake.fullRawManifestMutex.Lock()
	defer fake.fullRawManifestMutex.Unlock()
	fake.FullRawManifestStub = stub
}

func (fake *FakeManifestParser) FullRawManifestReturns(result1 []byte) {
	fake.fullRawManifestMutex.Lock()
	defer fake.fullRawManifestMutex.Unlock()
	fake.FullRawManifestStub = nil
	fake.fullRawManifestReturns = struct {
		result1 []byte
	}{result1}
}

func (fake *FakeManifestParser) FullRawManifestReturnsOnCall(i int, result1 []byte) {
	fake.fullRawManifestMutex.Lock()
	defer fake.fullRawManifestMutex.Unlock()
	fake.FullRawManifestStub = nil
	if fake.fullRawManifestReturnsOnCall == nil {
		fake.fullRawManifestReturnsOnCall = make(map[int]struct {
			result1 []byte
		})
	}
	fake.fullRawManifestReturnsOnCall[i] = struct {
		result1 []byte
	}{result1}
}

func (fake *FakeManifestParser) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.appsMutex.RLock()
	defer fake.appsMutex.RUnlock()
	fake.fullRawManifestMutex.RLock()
	defer fake.fullRawManifestMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeManifestParser) recordInvocation(key string, args []interface{}) {
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

var _ v7pushaction.ManifestParser = new(FakeManifestParser)