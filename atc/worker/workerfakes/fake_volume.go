// Code generated by counterfeiter. DO NOT EDIT.
package workerfakes

import (
	"context"
	"io"
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/baggageclaim"
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/worker"
)

type FakeVolume struct {
	COWStrategyStub        func() baggageclaim.COWStrategy
	cOWStrategyMutex       sync.RWMutex
	cOWStrategyArgsForCall []struct {
	}
	cOWStrategyReturns struct {
		result1 baggageclaim.COWStrategy
	}
	cOWStrategyReturnsOnCall map[int]struct {
		result1 baggageclaim.COWStrategy
	}
	CreateChildForContainerStub        func(db.CreatingContainer, string) (db.CreatingVolume, error)
	createChildForContainerMutex       sync.RWMutex
	createChildForContainerArgsForCall []struct {
		arg1 db.CreatingContainer
		arg2 string
	}
	createChildForContainerReturns struct {
		result1 db.CreatingVolume
		result2 error
	}
	createChildForContainerReturnsOnCall map[int]struct {
		result1 db.CreatingVolume
		result2 error
	}
	DestroyStub        func() error
	destroyMutex       sync.RWMutex
	destroyArgsForCall []struct {
	}
	destroyReturns struct {
		result1 error
	}
	destroyReturnsOnCall map[int]struct {
		result1 error
	}
	GetResourceCacheIDStub        func() int
	getResourceCacheIDMutex       sync.RWMutex
	getResourceCacheIDArgsForCall []struct {
	}
	getResourceCacheIDReturns struct {
		result1 int
	}
	getResourceCacheIDReturnsOnCall map[int]struct {
		result1 int
	}
	GetStreamInP2pUrlStub        func(context.Context, string) (string, error)
	getStreamInP2pUrlMutex       sync.RWMutex
	getStreamInP2pUrlArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getStreamInP2pUrlReturns struct {
		result1 string
		result2 error
	}
	getStreamInP2pUrlReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	HandleStub        func() string
	handleMutex       sync.RWMutex
	handleArgsForCall []struct {
	}
	handleReturns struct {
		result1 string
	}
	handleReturnsOnCall map[int]struct {
		result1 string
	}
	InitializeArtifactStub        func(string, int) (db.WorkerArtifact, error)
	initializeArtifactMutex       sync.RWMutex
	initializeArtifactArgsForCall []struct {
		arg1 string
		arg2 int
	}
	initializeArtifactReturns struct {
		result1 db.WorkerArtifact
		result2 error
	}
	initializeArtifactReturnsOnCall map[int]struct {
		result1 db.WorkerArtifact
		result2 error
	}
	InitializeResourceCacheStub        func(db.ResourceCache) error
	initializeResourceCacheMutex       sync.RWMutex
	initializeResourceCacheArgsForCall []struct {
		arg1 db.ResourceCache
	}
	initializeResourceCacheReturns struct {
		result1 error
	}
	initializeResourceCacheReturnsOnCall map[int]struct {
		result1 error
	}
	InitializeTaskCacheStub        func(lager.Logger, int, string, string, bool) error
	initializeTaskCacheMutex       sync.RWMutex
	initializeTaskCacheArgsForCall []struct {
		arg1 lager.Logger
		arg2 int
		arg3 string
		arg4 string
		arg5 bool
	}
	initializeTaskCacheReturns struct {
		result1 error
	}
	initializeTaskCacheReturnsOnCall map[int]struct {
		result1 error
	}
	PathStub        func() string
	pathMutex       sync.RWMutex
	pathArgsForCall []struct {
	}
	pathReturns struct {
		result1 string
	}
	pathReturnsOnCall map[int]struct {
		result1 string
	}
	PropertiesStub        func() (baggageclaim.VolumeProperties, error)
	propertiesMutex       sync.RWMutex
	propertiesArgsForCall []struct {
	}
	propertiesReturns struct {
		result1 baggageclaim.VolumeProperties
		result2 error
	}
	propertiesReturnsOnCall map[int]struct {
		result1 baggageclaim.VolumeProperties
		result2 error
	}
	SetPrivilegedStub        func(bool) error
	setPrivilegedMutex       sync.RWMutex
	setPrivilegedArgsForCall []struct {
		arg1 bool
	}
	setPrivilegedReturns struct {
		result1 error
	}
	setPrivilegedReturnsOnCall map[int]struct {
		result1 error
	}
	SetPropertyStub        func(string, string) error
	setPropertyMutex       sync.RWMutex
	setPropertyArgsForCall []struct {
		arg1 string
		arg2 string
	}
	setPropertyReturns struct {
		result1 error
	}
	setPropertyReturnsOnCall map[int]struct {
		result1 error
	}
	StreamInStub        func(context.Context, string, baggageclaim.Encoding, io.Reader) error
	streamInMutex       sync.RWMutex
	streamInArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 baggageclaim.Encoding
		arg4 io.Reader
	}
	streamInReturns struct {
		result1 error
	}
	streamInReturnsOnCall map[int]struct {
		result1 error
	}
	StreamOutStub        func(context.Context, string, baggageclaim.Encoding) (io.ReadCloser, error)
	streamOutMutex       sync.RWMutex
	streamOutArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 baggageclaim.Encoding
	}
	streamOutReturns struct {
		result1 io.ReadCloser
		result2 error
	}
	streamOutReturnsOnCall map[int]struct {
		result1 io.ReadCloser
		result2 error
	}
	StreamP2pOutStub        func(context.Context, string, string, baggageclaim.Encoding) error
	streamP2pOutMutex       sync.RWMutex
	streamP2pOutArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 baggageclaim.Encoding
	}
	streamP2pOutReturns struct {
		result1 error
	}
	streamP2pOutReturnsOnCall map[int]struct {
		result1 error
	}
	WorkerNameStub        func() string
	workerNameMutex       sync.RWMutex
	workerNameArgsForCall []struct {
	}
	workerNameReturns struct {
		result1 string
	}
	workerNameReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVolume) COWStrategy() baggageclaim.COWStrategy {
	fake.cOWStrategyMutex.Lock()
	ret, specificReturn := fake.cOWStrategyReturnsOnCall[len(fake.cOWStrategyArgsForCall)]
	fake.cOWStrategyArgsForCall = append(fake.cOWStrategyArgsForCall, struct {
	}{})
	stub := fake.COWStrategyStub
	fakeReturns := fake.cOWStrategyReturns
	fake.recordInvocation("COWStrategy", []interface{}{})
	fake.cOWStrategyMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeVolume) COWStrategyCallCount() int {
	fake.cOWStrategyMutex.RLock()
	defer fake.cOWStrategyMutex.RUnlock()
	return len(fake.cOWStrategyArgsForCall)
}

func (fake *FakeVolume) COWStrategyCalls(stub func() baggageclaim.COWStrategy) {
	fake.cOWStrategyMutex.Lock()
	defer fake.cOWStrategyMutex.Unlock()
	fake.COWStrategyStub = stub
}

func (fake *FakeVolume) COWStrategyReturns(result1 baggageclaim.COWStrategy) {
	fake.cOWStrategyMutex.Lock()
	defer fake.cOWStrategyMutex.Unlock()
	fake.COWStrategyStub = nil
	fake.cOWStrategyReturns = struct {
		result1 baggageclaim.COWStrategy
	}{result1}
}

func (fake *FakeVolume) COWStrategyReturnsOnCall(i int, result1 baggageclaim.COWStrategy) {
	fake.cOWStrategyMutex.Lock()
	defer fake.cOWStrategyMutex.Unlock()
	fake.COWStrategyStub = nil
	if fake.cOWStrategyReturnsOnCall == nil {
		fake.cOWStrategyReturnsOnCall = make(map[int]struct {
			result1 baggageclaim.COWStrategy
		})
	}
	fake.cOWStrategyReturnsOnCall[i] = struct {
		result1 baggageclaim.COWStrategy
	}{result1}
}

func (fake *FakeVolume) CreateChildForContainer(arg1 db.CreatingContainer, arg2 string) (db.CreatingVolume, error) {
	fake.createChildForContainerMutex.Lock()
	ret, specificReturn := fake.createChildForContainerReturnsOnCall[len(fake.createChildForContainerArgsForCall)]
	fake.createChildForContainerArgsForCall = append(fake.createChildForContainerArgsForCall, struct {
		arg1 db.CreatingContainer
		arg2 string
	}{arg1, arg2})
	stub := fake.CreateChildForContainerStub
	fakeReturns := fake.createChildForContainerReturns
	fake.recordInvocation("CreateChildForContainer", []interface{}{arg1, arg2})
	fake.createChildForContainerMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVolume) CreateChildForContainerCallCount() int {
	fake.createChildForContainerMutex.RLock()
	defer fake.createChildForContainerMutex.RUnlock()
	return len(fake.createChildForContainerArgsForCall)
}

func (fake *FakeVolume) CreateChildForContainerCalls(stub func(db.CreatingContainer, string) (db.CreatingVolume, error)) {
	fake.createChildForContainerMutex.Lock()
	defer fake.createChildForContainerMutex.Unlock()
	fake.CreateChildForContainerStub = stub
}

func (fake *FakeVolume) CreateChildForContainerArgsForCall(i int) (db.CreatingContainer, string) {
	fake.createChildForContainerMutex.RLock()
	defer fake.createChildForContainerMutex.RUnlock()
	argsForCall := fake.createChildForContainerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeVolume) CreateChildForContainerReturns(result1 db.CreatingVolume, result2 error) {
	fake.createChildForContainerMutex.Lock()
	defer fake.createChildForContainerMutex.Unlock()
	fake.CreateChildForContainerStub = nil
	fake.createChildForContainerReturns = struct {
		result1 db.CreatingVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolume) CreateChildForContainerReturnsOnCall(i int, result1 db.CreatingVolume, result2 error) {
	fake.createChildForContainerMutex.Lock()
	defer fake.createChildForContainerMutex.Unlock()
	fake.CreateChildForContainerStub = nil
	if fake.createChildForContainerReturnsOnCall == nil {
		fake.createChildForContainerReturnsOnCall = make(map[int]struct {
			result1 db.CreatingVolume
			result2 error
		})
	}
	fake.createChildForContainerReturnsOnCall[i] = struct {
		result1 db.CreatingVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolume) Destroy() error {
	fake.destroyMutex.Lock()
	ret, specificReturn := fake.destroyReturnsOnCall[len(fake.destroyArgsForCall)]
	fake.destroyArgsForCall = append(fake.destroyArgsForCall, struct {
	}{})
	stub := fake.DestroyStub
	fakeReturns := fake.destroyReturns
	fake.recordInvocation("Destroy", []interface{}{})
	fake.destroyMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeVolume) DestroyCallCount() int {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return len(fake.destroyArgsForCall)
}

func (fake *FakeVolume) DestroyCalls(stub func() error) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = stub
}

func (fake *FakeVolume) DestroyReturns(result1 error) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = nil
	fake.destroyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) DestroyReturnsOnCall(i int, result1 error) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = nil
	if fake.destroyReturnsOnCall == nil {
		fake.destroyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.destroyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) GetResourceCacheID() int {
	fake.getResourceCacheIDMutex.Lock()
	ret, specificReturn := fake.getResourceCacheIDReturnsOnCall[len(fake.getResourceCacheIDArgsForCall)]
	fake.getResourceCacheIDArgsForCall = append(fake.getResourceCacheIDArgsForCall, struct {
	}{})
	stub := fake.GetResourceCacheIDStub
	fakeReturns := fake.getResourceCacheIDReturns
	fake.recordInvocation("GetResourceCacheID", []interface{}{})
	fake.getResourceCacheIDMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeVolume) GetResourceCacheIDCallCount() int {
	fake.getResourceCacheIDMutex.RLock()
	defer fake.getResourceCacheIDMutex.RUnlock()
	return len(fake.getResourceCacheIDArgsForCall)
}

func (fake *FakeVolume) GetResourceCacheIDCalls(stub func() int) {
	fake.getResourceCacheIDMutex.Lock()
	defer fake.getResourceCacheIDMutex.Unlock()
	fake.GetResourceCacheIDStub = stub
}

func (fake *FakeVolume) GetResourceCacheIDReturns(result1 int) {
	fake.getResourceCacheIDMutex.Lock()
	defer fake.getResourceCacheIDMutex.Unlock()
	fake.GetResourceCacheIDStub = nil
	fake.getResourceCacheIDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeVolume) GetResourceCacheIDReturnsOnCall(i int, result1 int) {
	fake.getResourceCacheIDMutex.Lock()
	defer fake.getResourceCacheIDMutex.Unlock()
	fake.GetResourceCacheIDStub = nil
	if fake.getResourceCacheIDReturnsOnCall == nil {
		fake.getResourceCacheIDReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.getResourceCacheIDReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeVolume) GetStreamInP2pUrl(arg1 context.Context, arg2 string) (string, error) {
	fake.getStreamInP2pUrlMutex.Lock()
	ret, specificReturn := fake.getStreamInP2pUrlReturnsOnCall[len(fake.getStreamInP2pUrlArgsForCall)]
	fake.getStreamInP2pUrlArgsForCall = append(fake.getStreamInP2pUrlArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.GetStreamInP2pUrlStub
	fakeReturns := fake.getStreamInP2pUrlReturns
	fake.recordInvocation("GetStreamInP2pUrl", []interface{}{arg1, arg2})
	fake.getStreamInP2pUrlMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVolume) GetStreamInP2pUrlCallCount() int {
	fake.getStreamInP2pUrlMutex.RLock()
	defer fake.getStreamInP2pUrlMutex.RUnlock()
	return len(fake.getStreamInP2pUrlArgsForCall)
}

func (fake *FakeVolume) GetStreamInP2pUrlCalls(stub func(context.Context, string) (string, error)) {
	fake.getStreamInP2pUrlMutex.Lock()
	defer fake.getStreamInP2pUrlMutex.Unlock()
	fake.GetStreamInP2pUrlStub = stub
}

func (fake *FakeVolume) GetStreamInP2pUrlArgsForCall(i int) (context.Context, string) {
	fake.getStreamInP2pUrlMutex.RLock()
	defer fake.getStreamInP2pUrlMutex.RUnlock()
	argsForCall := fake.getStreamInP2pUrlArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeVolume) GetStreamInP2pUrlReturns(result1 string, result2 error) {
	fake.getStreamInP2pUrlMutex.Lock()
	defer fake.getStreamInP2pUrlMutex.Unlock()
	fake.GetStreamInP2pUrlStub = nil
	fake.getStreamInP2pUrlReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeVolume) GetStreamInP2pUrlReturnsOnCall(i int, result1 string, result2 error) {
	fake.getStreamInP2pUrlMutex.Lock()
	defer fake.getStreamInP2pUrlMutex.Unlock()
	fake.GetStreamInP2pUrlStub = nil
	if fake.getStreamInP2pUrlReturnsOnCall == nil {
		fake.getStreamInP2pUrlReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getStreamInP2pUrlReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeVolume) Handle() string {
	fake.handleMutex.Lock()
	ret, specificReturn := fake.handleReturnsOnCall[len(fake.handleArgsForCall)]
	fake.handleArgsForCall = append(fake.handleArgsForCall, struct {
	}{})
	stub := fake.HandleStub
	fakeReturns := fake.handleReturns
	fake.recordInvocation("Handle", []interface{}{})
	fake.handleMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeVolume) HandleCallCount() int {
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	return len(fake.handleArgsForCall)
}

func (fake *FakeVolume) HandleCalls(stub func() string) {
	fake.handleMutex.Lock()
	defer fake.handleMutex.Unlock()
	fake.HandleStub = stub
}

func (fake *FakeVolume) HandleReturns(result1 string) {
	fake.handleMutex.Lock()
	defer fake.handleMutex.Unlock()
	fake.HandleStub = nil
	fake.handleReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeVolume) HandleReturnsOnCall(i int, result1 string) {
	fake.handleMutex.Lock()
	defer fake.handleMutex.Unlock()
	fake.HandleStub = nil
	if fake.handleReturnsOnCall == nil {
		fake.handleReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.handleReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeVolume) InitializeArtifact(arg1 string, arg2 int) (db.WorkerArtifact, error) {
	fake.initializeArtifactMutex.Lock()
	ret, specificReturn := fake.initializeArtifactReturnsOnCall[len(fake.initializeArtifactArgsForCall)]
	fake.initializeArtifactArgsForCall = append(fake.initializeArtifactArgsForCall, struct {
		arg1 string
		arg2 int
	}{arg1, arg2})
	stub := fake.InitializeArtifactStub
	fakeReturns := fake.initializeArtifactReturns
	fake.recordInvocation("InitializeArtifact", []interface{}{arg1, arg2})
	fake.initializeArtifactMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVolume) InitializeArtifactCallCount() int {
	fake.initializeArtifactMutex.RLock()
	defer fake.initializeArtifactMutex.RUnlock()
	return len(fake.initializeArtifactArgsForCall)
}

func (fake *FakeVolume) InitializeArtifactCalls(stub func(string, int) (db.WorkerArtifact, error)) {
	fake.initializeArtifactMutex.Lock()
	defer fake.initializeArtifactMutex.Unlock()
	fake.InitializeArtifactStub = stub
}

func (fake *FakeVolume) InitializeArtifactArgsForCall(i int) (string, int) {
	fake.initializeArtifactMutex.RLock()
	defer fake.initializeArtifactMutex.RUnlock()
	argsForCall := fake.initializeArtifactArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeVolume) InitializeArtifactReturns(result1 db.WorkerArtifact, result2 error) {
	fake.initializeArtifactMutex.Lock()
	defer fake.initializeArtifactMutex.Unlock()
	fake.InitializeArtifactStub = nil
	fake.initializeArtifactReturns = struct {
		result1 db.WorkerArtifact
		result2 error
	}{result1, result2}
}

func (fake *FakeVolume) InitializeArtifactReturnsOnCall(i int, result1 db.WorkerArtifact, result2 error) {
	fake.initializeArtifactMutex.Lock()
	defer fake.initializeArtifactMutex.Unlock()
	fake.InitializeArtifactStub = nil
	if fake.initializeArtifactReturnsOnCall == nil {
		fake.initializeArtifactReturnsOnCall = make(map[int]struct {
			result1 db.WorkerArtifact
			result2 error
		})
	}
	fake.initializeArtifactReturnsOnCall[i] = struct {
		result1 db.WorkerArtifact
		result2 error
	}{result1, result2}
}

func (fake *FakeVolume) InitializeResourceCache(arg1 db.ResourceCache) error {
	fake.initializeResourceCacheMutex.Lock()
	ret, specificReturn := fake.initializeResourceCacheReturnsOnCall[len(fake.initializeResourceCacheArgsForCall)]
	fake.initializeResourceCacheArgsForCall = append(fake.initializeResourceCacheArgsForCall, struct {
		arg1 db.ResourceCache
	}{arg1})
	stub := fake.InitializeResourceCacheStub
	fakeReturns := fake.initializeResourceCacheReturns
	fake.recordInvocation("InitializeResourceCache", []interface{}{arg1})
	fake.initializeResourceCacheMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeVolume) InitializeResourceCacheCallCount() int {
	fake.initializeResourceCacheMutex.RLock()
	defer fake.initializeResourceCacheMutex.RUnlock()
	return len(fake.initializeResourceCacheArgsForCall)
}

func (fake *FakeVolume) InitializeResourceCacheCalls(stub func(db.ResourceCache) error) {
	fake.initializeResourceCacheMutex.Lock()
	defer fake.initializeResourceCacheMutex.Unlock()
	fake.InitializeResourceCacheStub = stub
}

func (fake *FakeVolume) InitializeResourceCacheArgsForCall(i int) db.ResourceCache {
	fake.initializeResourceCacheMutex.RLock()
	defer fake.initializeResourceCacheMutex.RUnlock()
	argsForCall := fake.initializeResourceCacheArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeVolume) InitializeResourceCacheReturns(result1 error) {
	fake.initializeResourceCacheMutex.Lock()
	defer fake.initializeResourceCacheMutex.Unlock()
	fake.InitializeResourceCacheStub = nil
	fake.initializeResourceCacheReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) InitializeResourceCacheReturnsOnCall(i int, result1 error) {
	fake.initializeResourceCacheMutex.Lock()
	defer fake.initializeResourceCacheMutex.Unlock()
	fake.InitializeResourceCacheStub = nil
	if fake.initializeResourceCacheReturnsOnCall == nil {
		fake.initializeResourceCacheReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.initializeResourceCacheReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) InitializeTaskCache(arg1 lager.Logger, arg2 int, arg3 string, arg4 string, arg5 bool) error {
	fake.initializeTaskCacheMutex.Lock()
	ret, specificReturn := fake.initializeTaskCacheReturnsOnCall[len(fake.initializeTaskCacheArgsForCall)]
	fake.initializeTaskCacheArgsForCall = append(fake.initializeTaskCacheArgsForCall, struct {
		arg1 lager.Logger
		arg2 int
		arg3 string
		arg4 string
		arg5 bool
	}{arg1, arg2, arg3, arg4, arg5})
	stub := fake.InitializeTaskCacheStub
	fakeReturns := fake.initializeTaskCacheReturns
	fake.recordInvocation("InitializeTaskCache", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.initializeTaskCacheMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeVolume) InitializeTaskCacheCallCount() int {
	fake.initializeTaskCacheMutex.RLock()
	defer fake.initializeTaskCacheMutex.RUnlock()
	return len(fake.initializeTaskCacheArgsForCall)
}

func (fake *FakeVolume) InitializeTaskCacheCalls(stub func(lager.Logger, int, string, string, bool) error) {
	fake.initializeTaskCacheMutex.Lock()
	defer fake.initializeTaskCacheMutex.Unlock()
	fake.InitializeTaskCacheStub = stub
}

func (fake *FakeVolume) InitializeTaskCacheArgsForCall(i int) (lager.Logger, int, string, string, bool) {
	fake.initializeTaskCacheMutex.RLock()
	defer fake.initializeTaskCacheMutex.RUnlock()
	argsForCall := fake.initializeTaskCacheArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeVolume) InitializeTaskCacheReturns(result1 error) {
	fake.initializeTaskCacheMutex.Lock()
	defer fake.initializeTaskCacheMutex.Unlock()
	fake.InitializeTaskCacheStub = nil
	fake.initializeTaskCacheReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) InitializeTaskCacheReturnsOnCall(i int, result1 error) {
	fake.initializeTaskCacheMutex.Lock()
	defer fake.initializeTaskCacheMutex.Unlock()
	fake.InitializeTaskCacheStub = nil
	if fake.initializeTaskCacheReturnsOnCall == nil {
		fake.initializeTaskCacheReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.initializeTaskCacheReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) Path() string {
	fake.pathMutex.Lock()
	ret, specificReturn := fake.pathReturnsOnCall[len(fake.pathArgsForCall)]
	fake.pathArgsForCall = append(fake.pathArgsForCall, struct {
	}{})
	stub := fake.PathStub
	fakeReturns := fake.pathReturns
	fake.recordInvocation("Path", []interface{}{})
	fake.pathMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeVolume) PathCallCount() int {
	fake.pathMutex.RLock()
	defer fake.pathMutex.RUnlock()
	return len(fake.pathArgsForCall)
}

func (fake *FakeVolume) PathCalls(stub func() string) {
	fake.pathMutex.Lock()
	defer fake.pathMutex.Unlock()
	fake.PathStub = stub
}

func (fake *FakeVolume) PathReturns(result1 string) {
	fake.pathMutex.Lock()
	defer fake.pathMutex.Unlock()
	fake.PathStub = nil
	fake.pathReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeVolume) PathReturnsOnCall(i int, result1 string) {
	fake.pathMutex.Lock()
	defer fake.pathMutex.Unlock()
	fake.PathStub = nil
	if fake.pathReturnsOnCall == nil {
		fake.pathReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.pathReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeVolume) Properties() (baggageclaim.VolumeProperties, error) {
	fake.propertiesMutex.Lock()
	ret, specificReturn := fake.propertiesReturnsOnCall[len(fake.propertiesArgsForCall)]
	fake.propertiesArgsForCall = append(fake.propertiesArgsForCall, struct {
	}{})
	stub := fake.PropertiesStub
	fakeReturns := fake.propertiesReturns
	fake.recordInvocation("Properties", []interface{}{})
	fake.propertiesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVolume) PropertiesCallCount() int {
	fake.propertiesMutex.RLock()
	defer fake.propertiesMutex.RUnlock()
	return len(fake.propertiesArgsForCall)
}

func (fake *FakeVolume) PropertiesCalls(stub func() (baggageclaim.VolumeProperties, error)) {
	fake.propertiesMutex.Lock()
	defer fake.propertiesMutex.Unlock()
	fake.PropertiesStub = stub
}

func (fake *FakeVolume) PropertiesReturns(result1 baggageclaim.VolumeProperties, result2 error) {
	fake.propertiesMutex.Lock()
	defer fake.propertiesMutex.Unlock()
	fake.PropertiesStub = nil
	fake.propertiesReturns = struct {
		result1 baggageclaim.VolumeProperties
		result2 error
	}{result1, result2}
}

func (fake *FakeVolume) PropertiesReturnsOnCall(i int, result1 baggageclaim.VolumeProperties, result2 error) {
	fake.propertiesMutex.Lock()
	defer fake.propertiesMutex.Unlock()
	fake.PropertiesStub = nil
	if fake.propertiesReturnsOnCall == nil {
		fake.propertiesReturnsOnCall = make(map[int]struct {
			result1 baggageclaim.VolumeProperties
			result2 error
		})
	}
	fake.propertiesReturnsOnCall[i] = struct {
		result1 baggageclaim.VolumeProperties
		result2 error
	}{result1, result2}
}

func (fake *FakeVolume) SetPrivileged(arg1 bool) error {
	fake.setPrivilegedMutex.Lock()
	ret, specificReturn := fake.setPrivilegedReturnsOnCall[len(fake.setPrivilegedArgsForCall)]
	fake.setPrivilegedArgsForCall = append(fake.setPrivilegedArgsForCall, struct {
		arg1 bool
	}{arg1})
	stub := fake.SetPrivilegedStub
	fakeReturns := fake.setPrivilegedReturns
	fake.recordInvocation("SetPrivileged", []interface{}{arg1})
	fake.setPrivilegedMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeVolume) SetPrivilegedCallCount() int {
	fake.setPrivilegedMutex.RLock()
	defer fake.setPrivilegedMutex.RUnlock()
	return len(fake.setPrivilegedArgsForCall)
}

func (fake *FakeVolume) SetPrivilegedCalls(stub func(bool) error) {
	fake.setPrivilegedMutex.Lock()
	defer fake.setPrivilegedMutex.Unlock()
	fake.SetPrivilegedStub = stub
}

func (fake *FakeVolume) SetPrivilegedArgsForCall(i int) bool {
	fake.setPrivilegedMutex.RLock()
	defer fake.setPrivilegedMutex.RUnlock()
	argsForCall := fake.setPrivilegedArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeVolume) SetPrivilegedReturns(result1 error) {
	fake.setPrivilegedMutex.Lock()
	defer fake.setPrivilegedMutex.Unlock()
	fake.SetPrivilegedStub = nil
	fake.setPrivilegedReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) SetPrivilegedReturnsOnCall(i int, result1 error) {
	fake.setPrivilegedMutex.Lock()
	defer fake.setPrivilegedMutex.Unlock()
	fake.SetPrivilegedStub = nil
	if fake.setPrivilegedReturnsOnCall == nil {
		fake.setPrivilegedReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setPrivilegedReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) SetProperty(arg1 string, arg2 string) error {
	fake.setPropertyMutex.Lock()
	ret, specificReturn := fake.setPropertyReturnsOnCall[len(fake.setPropertyArgsForCall)]
	fake.setPropertyArgsForCall = append(fake.setPropertyArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.SetPropertyStub
	fakeReturns := fake.setPropertyReturns
	fake.recordInvocation("SetProperty", []interface{}{arg1, arg2})
	fake.setPropertyMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeVolume) SetPropertyCallCount() int {
	fake.setPropertyMutex.RLock()
	defer fake.setPropertyMutex.RUnlock()
	return len(fake.setPropertyArgsForCall)
}

func (fake *FakeVolume) SetPropertyCalls(stub func(string, string) error) {
	fake.setPropertyMutex.Lock()
	defer fake.setPropertyMutex.Unlock()
	fake.SetPropertyStub = stub
}

func (fake *FakeVolume) SetPropertyArgsForCall(i int) (string, string) {
	fake.setPropertyMutex.RLock()
	defer fake.setPropertyMutex.RUnlock()
	argsForCall := fake.setPropertyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeVolume) SetPropertyReturns(result1 error) {
	fake.setPropertyMutex.Lock()
	defer fake.setPropertyMutex.Unlock()
	fake.SetPropertyStub = nil
	fake.setPropertyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) SetPropertyReturnsOnCall(i int, result1 error) {
	fake.setPropertyMutex.Lock()
	defer fake.setPropertyMutex.Unlock()
	fake.SetPropertyStub = nil
	if fake.setPropertyReturnsOnCall == nil {
		fake.setPropertyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setPropertyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) StreamIn(arg1 context.Context, arg2 string, arg3 baggageclaim.Encoding, arg4 io.Reader) error {
	fake.streamInMutex.Lock()
	ret, specificReturn := fake.streamInReturnsOnCall[len(fake.streamInArgsForCall)]
	fake.streamInArgsForCall = append(fake.streamInArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 baggageclaim.Encoding
		arg4 io.Reader
	}{arg1, arg2, arg3, arg4})
	stub := fake.StreamInStub
	fakeReturns := fake.streamInReturns
	fake.recordInvocation("StreamIn", []interface{}{arg1, arg2, arg3, arg4})
	fake.streamInMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeVolume) StreamInCallCount() int {
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	return len(fake.streamInArgsForCall)
}

func (fake *FakeVolume) StreamInCalls(stub func(context.Context, string, baggageclaim.Encoding, io.Reader) error) {
	fake.streamInMutex.Lock()
	defer fake.streamInMutex.Unlock()
	fake.StreamInStub = stub
}

func (fake *FakeVolume) StreamInArgsForCall(i int) (context.Context, string, baggageclaim.Encoding, io.Reader) {
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	argsForCall := fake.streamInArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeVolume) StreamInReturns(result1 error) {
	fake.streamInMutex.Lock()
	defer fake.streamInMutex.Unlock()
	fake.StreamInStub = nil
	fake.streamInReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) StreamInReturnsOnCall(i int, result1 error) {
	fake.streamInMutex.Lock()
	defer fake.streamInMutex.Unlock()
	fake.StreamInStub = nil
	if fake.streamInReturnsOnCall == nil {
		fake.streamInReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.streamInReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) StreamOut(arg1 context.Context, arg2 string, arg3 baggageclaim.Encoding) (io.ReadCloser, error) {
	fake.streamOutMutex.Lock()
	ret, specificReturn := fake.streamOutReturnsOnCall[len(fake.streamOutArgsForCall)]
	fake.streamOutArgsForCall = append(fake.streamOutArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 baggageclaim.Encoding
	}{arg1, arg2, arg3})
	stub := fake.StreamOutStub
	fakeReturns := fake.streamOutReturns
	fake.recordInvocation("StreamOut", []interface{}{arg1, arg2, arg3})
	fake.streamOutMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVolume) StreamOutCallCount() int {
	fake.streamOutMutex.RLock()
	defer fake.streamOutMutex.RUnlock()
	return len(fake.streamOutArgsForCall)
}

func (fake *FakeVolume) StreamOutCalls(stub func(context.Context, string, baggageclaim.Encoding) (io.ReadCloser, error)) {
	fake.streamOutMutex.Lock()
	defer fake.streamOutMutex.Unlock()
	fake.StreamOutStub = stub
}

func (fake *FakeVolume) StreamOutArgsForCall(i int) (context.Context, string, baggageclaim.Encoding) {
	fake.streamOutMutex.RLock()
	defer fake.streamOutMutex.RUnlock()
	argsForCall := fake.streamOutArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeVolume) StreamOutReturns(result1 io.ReadCloser, result2 error) {
	fake.streamOutMutex.Lock()
	defer fake.streamOutMutex.Unlock()
	fake.StreamOutStub = nil
	fake.streamOutReturns = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeVolume) StreamOutReturnsOnCall(i int, result1 io.ReadCloser, result2 error) {
	fake.streamOutMutex.Lock()
	defer fake.streamOutMutex.Unlock()
	fake.StreamOutStub = nil
	if fake.streamOutReturnsOnCall == nil {
		fake.streamOutReturnsOnCall = make(map[int]struct {
			result1 io.ReadCloser
			result2 error
		})
	}
	fake.streamOutReturnsOnCall[i] = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeVolume) StreamP2pOut(arg1 context.Context, arg2 string, arg3 string, arg4 baggageclaim.Encoding) error {
	fake.streamP2pOutMutex.Lock()
	ret, specificReturn := fake.streamP2pOutReturnsOnCall[len(fake.streamP2pOutArgsForCall)]
	fake.streamP2pOutArgsForCall = append(fake.streamP2pOutArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 baggageclaim.Encoding
	}{arg1, arg2, arg3, arg4})
	stub := fake.StreamP2pOutStub
	fakeReturns := fake.streamP2pOutReturns
	fake.recordInvocation("StreamP2pOut", []interface{}{arg1, arg2, arg3, arg4})
	fake.streamP2pOutMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeVolume) StreamP2pOutCallCount() int {
	fake.streamP2pOutMutex.RLock()
	defer fake.streamP2pOutMutex.RUnlock()
	return len(fake.streamP2pOutArgsForCall)
}

func (fake *FakeVolume) StreamP2pOutCalls(stub func(context.Context, string, string, baggageclaim.Encoding) error) {
	fake.streamP2pOutMutex.Lock()
	defer fake.streamP2pOutMutex.Unlock()
	fake.StreamP2pOutStub = stub
}

func (fake *FakeVolume) StreamP2pOutArgsForCall(i int) (context.Context, string, string, baggageclaim.Encoding) {
	fake.streamP2pOutMutex.RLock()
	defer fake.streamP2pOutMutex.RUnlock()
	argsForCall := fake.streamP2pOutArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeVolume) StreamP2pOutReturns(result1 error) {
	fake.streamP2pOutMutex.Lock()
	defer fake.streamP2pOutMutex.Unlock()
	fake.StreamP2pOutStub = nil
	fake.streamP2pOutReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) StreamP2pOutReturnsOnCall(i int, result1 error) {
	fake.streamP2pOutMutex.Lock()
	defer fake.streamP2pOutMutex.Unlock()
	fake.StreamP2pOutStub = nil
	if fake.streamP2pOutReturnsOnCall == nil {
		fake.streamP2pOutReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.streamP2pOutReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) WorkerName() string {
	fake.workerNameMutex.Lock()
	ret, specificReturn := fake.workerNameReturnsOnCall[len(fake.workerNameArgsForCall)]
	fake.workerNameArgsForCall = append(fake.workerNameArgsForCall, struct {
	}{})
	stub := fake.WorkerNameStub
	fakeReturns := fake.workerNameReturns
	fake.recordInvocation("WorkerName", []interface{}{})
	fake.workerNameMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeVolume) WorkerNameCallCount() int {
	fake.workerNameMutex.RLock()
	defer fake.workerNameMutex.RUnlock()
	return len(fake.workerNameArgsForCall)
}

func (fake *FakeVolume) WorkerNameCalls(stub func() string) {
	fake.workerNameMutex.Lock()
	defer fake.workerNameMutex.Unlock()
	fake.WorkerNameStub = stub
}

func (fake *FakeVolume) WorkerNameReturns(result1 string) {
	fake.workerNameMutex.Lock()
	defer fake.workerNameMutex.Unlock()
	fake.WorkerNameStub = nil
	fake.workerNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeVolume) WorkerNameReturnsOnCall(i int, result1 string) {
	fake.workerNameMutex.Lock()
	defer fake.workerNameMutex.Unlock()
	fake.WorkerNameStub = nil
	if fake.workerNameReturnsOnCall == nil {
		fake.workerNameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.workerNameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeVolume) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cOWStrategyMutex.RLock()
	defer fake.cOWStrategyMutex.RUnlock()
	fake.createChildForContainerMutex.RLock()
	defer fake.createChildForContainerMutex.RUnlock()
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	fake.getResourceCacheIDMutex.RLock()
	defer fake.getResourceCacheIDMutex.RUnlock()
	fake.getStreamInP2pUrlMutex.RLock()
	defer fake.getStreamInP2pUrlMutex.RUnlock()
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	fake.initializeArtifactMutex.RLock()
	defer fake.initializeArtifactMutex.RUnlock()
	fake.initializeResourceCacheMutex.RLock()
	defer fake.initializeResourceCacheMutex.RUnlock()
	fake.initializeTaskCacheMutex.RLock()
	defer fake.initializeTaskCacheMutex.RUnlock()
	fake.pathMutex.RLock()
	defer fake.pathMutex.RUnlock()
	fake.propertiesMutex.RLock()
	defer fake.propertiesMutex.RUnlock()
	fake.setPrivilegedMutex.RLock()
	defer fake.setPrivilegedMutex.RUnlock()
	fake.setPropertyMutex.RLock()
	defer fake.setPropertyMutex.RUnlock()
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	fake.streamOutMutex.RLock()
	defer fake.streamOutMutex.RUnlock()
	fake.streamP2pOutMutex.RLock()
	defer fake.streamP2pOutMutex.RUnlock()
	fake.workerNameMutex.RLock()
	defer fake.workerNameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeVolume) recordInvocation(key string, args []interface{}) {
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

var _ worker.Volume = new(FakeVolume)
