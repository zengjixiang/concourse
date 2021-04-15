// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"

	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/db"
)

type FakeCreatedVolume struct {
	BaseResourceTypeStub        func() (*db.UsedWorkerBaseResourceType, error)
	baseResourceTypeMutex       sync.RWMutex
	baseResourceTypeArgsForCall []struct {
	}
	baseResourceTypeReturns struct {
		result1 *db.UsedWorkerBaseResourceType
		result2 error
	}
	baseResourceTypeReturnsOnCall map[int]struct {
		result1 *db.UsedWorkerBaseResourceType
		result2 error
	}
	ContainerHandleStub        func() string
	containerHandleMutex       sync.RWMutex
	containerHandleArgsForCall []struct {
	}
	containerHandleReturns struct {
		result1 string
	}
	containerHandleReturnsOnCall map[int]struct {
		result1 string
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
	DestroyingStub        func() (db.DestroyingVolume, error)
	destroyingMutex       sync.RWMutex
	destroyingArgsForCall []struct {
	}
	destroyingReturns struct {
		result1 db.DestroyingVolume
		result2 error
	}
	destroyingReturnsOnCall map[int]struct {
		result1 db.DestroyingVolume
		result2 error
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
	InitializeTaskCacheStub        func(int, string, string) error
	initializeTaskCacheMutex       sync.RWMutex
	initializeTaskCacheArgsForCall []struct {
		arg1 int
		arg2 string
		arg3 string
	}
	initializeTaskCacheReturns struct {
		result1 error
	}
	initializeTaskCacheReturnsOnCall map[int]struct {
		result1 error
	}
	ParentHandleStub        func() string
	parentHandleMutex       sync.RWMutex
	parentHandleArgsForCall []struct {
	}
	parentHandleReturns struct {
		result1 string
	}
	parentHandleReturnsOnCall map[int]struct {
		result1 string
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
	ResourceTypeStub        func() (*db.VolumeResourceType, error)
	resourceTypeMutex       sync.RWMutex
	resourceTypeArgsForCall []struct {
	}
	resourceTypeReturns struct {
		result1 *db.VolumeResourceType
		result2 error
	}
	resourceTypeReturnsOnCall map[int]struct {
		result1 *db.VolumeResourceType
		result2 error
	}
	TaskIdentifierStub        func() (int, atc.PipelineRef, string, string, error)
	taskIdentifierMutex       sync.RWMutex
	taskIdentifierArgsForCall []struct {
	}
	taskIdentifierReturns struct {
		result1 int
		result2 atc.PipelineRef
		result3 string
		result4 string
		result5 error
	}
	taskIdentifierReturnsOnCall map[int]struct {
		result1 int
		result2 atc.PipelineRef
		result3 string
		result4 string
		result5 error
	}
	TeamIDStub        func() int
	teamIDMutex       sync.RWMutex
	teamIDArgsForCall []struct {
	}
	teamIDReturns struct {
		result1 int
	}
	teamIDReturnsOnCall map[int]struct {
		result1 int
	}
	TypeStub        func() db.VolumeType
	typeMutex       sync.RWMutex
	typeArgsForCall []struct {
	}
	typeReturns struct {
		result1 db.VolumeType
	}
	typeReturnsOnCall map[int]struct {
		result1 db.VolumeType
	}
	WorkerArtifactIDStub        func() int
	workerArtifactIDMutex       sync.RWMutex
	workerArtifactIDArgsForCall []struct {
	}
	workerArtifactIDReturns struct {
		result1 int
	}
	workerArtifactIDReturnsOnCall map[int]struct {
		result1 int
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

func (fake *FakeCreatedVolume) BaseResourceType() (*db.UsedWorkerBaseResourceType, error) {
	fake.baseResourceTypeMutex.Lock()
	ret, specificReturn := fake.baseResourceTypeReturnsOnCall[len(fake.baseResourceTypeArgsForCall)]
	fake.baseResourceTypeArgsForCall = append(fake.baseResourceTypeArgsForCall, struct {
	}{})
	stub := fake.BaseResourceTypeStub
	fakeReturns := fake.baseResourceTypeReturns
	fake.recordInvocation("BaseResourceType", []interface{}{})
	fake.baseResourceTypeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCreatedVolume) BaseResourceTypeCallCount() int {
	fake.baseResourceTypeMutex.RLock()
	defer fake.baseResourceTypeMutex.RUnlock()
	return len(fake.baseResourceTypeArgsForCall)
}

func (fake *FakeCreatedVolume) BaseResourceTypeCalls(stub func() (*db.UsedWorkerBaseResourceType, error)) {
	fake.baseResourceTypeMutex.Lock()
	defer fake.baseResourceTypeMutex.Unlock()
	fake.BaseResourceTypeStub = stub
}

func (fake *FakeCreatedVolume) BaseResourceTypeReturns(result1 *db.UsedWorkerBaseResourceType, result2 error) {
	fake.baseResourceTypeMutex.Lock()
	defer fake.baseResourceTypeMutex.Unlock()
	fake.BaseResourceTypeStub = nil
	fake.baseResourceTypeReturns = struct {
		result1 *db.UsedWorkerBaseResourceType
		result2 error
	}{result1, result2}
}

func (fake *FakeCreatedVolume) BaseResourceTypeReturnsOnCall(i int, result1 *db.UsedWorkerBaseResourceType, result2 error) {
	fake.baseResourceTypeMutex.Lock()
	defer fake.baseResourceTypeMutex.Unlock()
	fake.BaseResourceTypeStub = nil
	if fake.baseResourceTypeReturnsOnCall == nil {
		fake.baseResourceTypeReturnsOnCall = make(map[int]struct {
			result1 *db.UsedWorkerBaseResourceType
			result2 error
		})
	}
	fake.baseResourceTypeReturnsOnCall[i] = struct {
		result1 *db.UsedWorkerBaseResourceType
		result2 error
	}{result1, result2}
}

func (fake *FakeCreatedVolume) ContainerHandle() string {
	fake.containerHandleMutex.Lock()
	ret, specificReturn := fake.containerHandleReturnsOnCall[len(fake.containerHandleArgsForCall)]
	fake.containerHandleArgsForCall = append(fake.containerHandleArgsForCall, struct {
	}{})
	stub := fake.ContainerHandleStub
	fakeReturns := fake.containerHandleReturns
	fake.recordInvocation("ContainerHandle", []interface{}{})
	fake.containerHandleMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCreatedVolume) ContainerHandleCallCount() int {
	fake.containerHandleMutex.RLock()
	defer fake.containerHandleMutex.RUnlock()
	return len(fake.containerHandleArgsForCall)
}

func (fake *FakeCreatedVolume) ContainerHandleCalls(stub func() string) {
	fake.containerHandleMutex.Lock()
	defer fake.containerHandleMutex.Unlock()
	fake.ContainerHandleStub = stub
}

func (fake *FakeCreatedVolume) ContainerHandleReturns(result1 string) {
	fake.containerHandleMutex.Lock()
	defer fake.containerHandleMutex.Unlock()
	fake.ContainerHandleStub = nil
	fake.containerHandleReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCreatedVolume) ContainerHandleReturnsOnCall(i int, result1 string) {
	fake.containerHandleMutex.Lock()
	defer fake.containerHandleMutex.Unlock()
	fake.ContainerHandleStub = nil
	if fake.containerHandleReturnsOnCall == nil {
		fake.containerHandleReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.containerHandleReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCreatedVolume) CreateChildForContainer(arg1 db.CreatingContainer, arg2 string) (db.CreatingVolume, error) {
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

func (fake *FakeCreatedVolume) CreateChildForContainerCallCount() int {
	fake.createChildForContainerMutex.RLock()
	defer fake.createChildForContainerMutex.RUnlock()
	return len(fake.createChildForContainerArgsForCall)
}

func (fake *FakeCreatedVolume) CreateChildForContainerCalls(stub func(db.CreatingContainer, string) (db.CreatingVolume, error)) {
	fake.createChildForContainerMutex.Lock()
	defer fake.createChildForContainerMutex.Unlock()
	fake.CreateChildForContainerStub = stub
}

func (fake *FakeCreatedVolume) CreateChildForContainerArgsForCall(i int) (db.CreatingContainer, string) {
	fake.createChildForContainerMutex.RLock()
	defer fake.createChildForContainerMutex.RUnlock()
	argsForCall := fake.createChildForContainerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCreatedVolume) CreateChildForContainerReturns(result1 db.CreatingVolume, result2 error) {
	fake.createChildForContainerMutex.Lock()
	defer fake.createChildForContainerMutex.Unlock()
	fake.CreateChildForContainerStub = nil
	fake.createChildForContainerReturns = struct {
		result1 db.CreatingVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeCreatedVolume) CreateChildForContainerReturnsOnCall(i int, result1 db.CreatingVolume, result2 error) {
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

func (fake *FakeCreatedVolume) Destroying() (db.DestroyingVolume, error) {
	fake.destroyingMutex.Lock()
	ret, specificReturn := fake.destroyingReturnsOnCall[len(fake.destroyingArgsForCall)]
	fake.destroyingArgsForCall = append(fake.destroyingArgsForCall, struct {
	}{})
	stub := fake.DestroyingStub
	fakeReturns := fake.destroyingReturns
	fake.recordInvocation("Destroying", []interface{}{})
	fake.destroyingMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCreatedVolume) DestroyingCallCount() int {
	fake.destroyingMutex.RLock()
	defer fake.destroyingMutex.RUnlock()
	return len(fake.destroyingArgsForCall)
}

func (fake *FakeCreatedVolume) DestroyingCalls(stub func() (db.DestroyingVolume, error)) {
	fake.destroyingMutex.Lock()
	defer fake.destroyingMutex.Unlock()
	fake.DestroyingStub = stub
}

func (fake *FakeCreatedVolume) DestroyingReturns(result1 db.DestroyingVolume, result2 error) {
	fake.destroyingMutex.Lock()
	defer fake.destroyingMutex.Unlock()
	fake.DestroyingStub = nil
	fake.destroyingReturns = struct {
		result1 db.DestroyingVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeCreatedVolume) DestroyingReturnsOnCall(i int, result1 db.DestroyingVolume, result2 error) {
	fake.destroyingMutex.Lock()
	defer fake.destroyingMutex.Unlock()
	fake.DestroyingStub = nil
	if fake.destroyingReturnsOnCall == nil {
		fake.destroyingReturnsOnCall = make(map[int]struct {
			result1 db.DestroyingVolume
			result2 error
		})
	}
	fake.destroyingReturnsOnCall[i] = struct {
		result1 db.DestroyingVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeCreatedVolume) GetResourceCacheID() int {
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

func (fake *FakeCreatedVolume) GetResourceCacheIDCallCount() int {
	fake.getResourceCacheIDMutex.RLock()
	defer fake.getResourceCacheIDMutex.RUnlock()
	return len(fake.getResourceCacheIDArgsForCall)
}

func (fake *FakeCreatedVolume) GetResourceCacheIDCalls(stub func() int) {
	fake.getResourceCacheIDMutex.Lock()
	defer fake.getResourceCacheIDMutex.Unlock()
	fake.GetResourceCacheIDStub = stub
}

func (fake *FakeCreatedVolume) GetResourceCacheIDReturns(result1 int) {
	fake.getResourceCacheIDMutex.Lock()
	defer fake.getResourceCacheIDMutex.Unlock()
	fake.GetResourceCacheIDStub = nil
	fake.getResourceCacheIDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeCreatedVolume) GetResourceCacheIDReturnsOnCall(i int, result1 int) {
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

func (fake *FakeCreatedVolume) Handle() string {
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

func (fake *FakeCreatedVolume) HandleCallCount() int {
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	return len(fake.handleArgsForCall)
}

func (fake *FakeCreatedVolume) HandleCalls(stub func() string) {
	fake.handleMutex.Lock()
	defer fake.handleMutex.Unlock()
	fake.HandleStub = stub
}

func (fake *FakeCreatedVolume) HandleReturns(result1 string) {
	fake.handleMutex.Lock()
	defer fake.handleMutex.Unlock()
	fake.HandleStub = nil
	fake.handleReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCreatedVolume) HandleReturnsOnCall(i int, result1 string) {
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

func (fake *FakeCreatedVolume) InitializeArtifact(arg1 string, arg2 int) (db.WorkerArtifact, error) {
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

func (fake *FakeCreatedVolume) InitializeArtifactCallCount() int {
	fake.initializeArtifactMutex.RLock()
	defer fake.initializeArtifactMutex.RUnlock()
	return len(fake.initializeArtifactArgsForCall)
}

func (fake *FakeCreatedVolume) InitializeArtifactCalls(stub func(string, int) (db.WorkerArtifact, error)) {
	fake.initializeArtifactMutex.Lock()
	defer fake.initializeArtifactMutex.Unlock()
	fake.InitializeArtifactStub = stub
}

func (fake *FakeCreatedVolume) InitializeArtifactArgsForCall(i int) (string, int) {
	fake.initializeArtifactMutex.RLock()
	defer fake.initializeArtifactMutex.RUnlock()
	argsForCall := fake.initializeArtifactArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCreatedVolume) InitializeArtifactReturns(result1 db.WorkerArtifact, result2 error) {
	fake.initializeArtifactMutex.Lock()
	defer fake.initializeArtifactMutex.Unlock()
	fake.InitializeArtifactStub = nil
	fake.initializeArtifactReturns = struct {
		result1 db.WorkerArtifact
		result2 error
	}{result1, result2}
}

func (fake *FakeCreatedVolume) InitializeArtifactReturnsOnCall(i int, result1 db.WorkerArtifact, result2 error) {
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

func (fake *FakeCreatedVolume) InitializeResourceCache(arg1 db.ResourceCache) error {
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

func (fake *FakeCreatedVolume) InitializeResourceCacheCallCount() int {
	fake.initializeResourceCacheMutex.RLock()
	defer fake.initializeResourceCacheMutex.RUnlock()
	return len(fake.initializeResourceCacheArgsForCall)
}

func (fake *FakeCreatedVolume) InitializeResourceCacheCalls(stub func(db.ResourceCache) error) {
	fake.initializeResourceCacheMutex.Lock()
	defer fake.initializeResourceCacheMutex.Unlock()
	fake.InitializeResourceCacheStub = stub
}

func (fake *FakeCreatedVolume) InitializeResourceCacheArgsForCall(i int) db.ResourceCache {
	fake.initializeResourceCacheMutex.RLock()
	defer fake.initializeResourceCacheMutex.RUnlock()
	argsForCall := fake.initializeResourceCacheArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCreatedVolume) InitializeResourceCacheReturns(result1 error) {
	fake.initializeResourceCacheMutex.Lock()
	defer fake.initializeResourceCacheMutex.Unlock()
	fake.InitializeResourceCacheStub = nil
	fake.initializeResourceCacheReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCreatedVolume) InitializeResourceCacheReturnsOnCall(i int, result1 error) {
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

func (fake *FakeCreatedVolume) InitializeTaskCache(arg1 int, arg2 string, arg3 string) error {
	fake.initializeTaskCacheMutex.Lock()
	ret, specificReturn := fake.initializeTaskCacheReturnsOnCall[len(fake.initializeTaskCacheArgsForCall)]
	fake.initializeTaskCacheArgsForCall = append(fake.initializeTaskCacheArgsForCall, struct {
		arg1 int
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.InitializeTaskCacheStub
	fakeReturns := fake.initializeTaskCacheReturns
	fake.recordInvocation("InitializeTaskCache", []interface{}{arg1, arg2, arg3})
	fake.initializeTaskCacheMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCreatedVolume) InitializeTaskCacheCallCount() int {
	fake.initializeTaskCacheMutex.RLock()
	defer fake.initializeTaskCacheMutex.RUnlock()
	return len(fake.initializeTaskCacheArgsForCall)
}

func (fake *FakeCreatedVolume) InitializeTaskCacheCalls(stub func(int, string, string) error) {
	fake.initializeTaskCacheMutex.Lock()
	defer fake.initializeTaskCacheMutex.Unlock()
	fake.InitializeTaskCacheStub = stub
}

func (fake *FakeCreatedVolume) InitializeTaskCacheArgsForCall(i int) (int, string, string) {
	fake.initializeTaskCacheMutex.RLock()
	defer fake.initializeTaskCacheMutex.RUnlock()
	argsForCall := fake.initializeTaskCacheArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCreatedVolume) InitializeTaskCacheReturns(result1 error) {
	fake.initializeTaskCacheMutex.Lock()
	defer fake.initializeTaskCacheMutex.Unlock()
	fake.InitializeTaskCacheStub = nil
	fake.initializeTaskCacheReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCreatedVolume) InitializeTaskCacheReturnsOnCall(i int, result1 error) {
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

func (fake *FakeCreatedVolume) ParentHandle() string {
	fake.parentHandleMutex.Lock()
	ret, specificReturn := fake.parentHandleReturnsOnCall[len(fake.parentHandleArgsForCall)]
	fake.parentHandleArgsForCall = append(fake.parentHandleArgsForCall, struct {
	}{})
	stub := fake.ParentHandleStub
	fakeReturns := fake.parentHandleReturns
	fake.recordInvocation("ParentHandle", []interface{}{})
	fake.parentHandleMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCreatedVolume) ParentHandleCallCount() int {
	fake.parentHandleMutex.RLock()
	defer fake.parentHandleMutex.RUnlock()
	return len(fake.parentHandleArgsForCall)
}

func (fake *FakeCreatedVolume) ParentHandleCalls(stub func() string) {
	fake.parentHandleMutex.Lock()
	defer fake.parentHandleMutex.Unlock()
	fake.ParentHandleStub = stub
}

func (fake *FakeCreatedVolume) ParentHandleReturns(result1 string) {
	fake.parentHandleMutex.Lock()
	defer fake.parentHandleMutex.Unlock()
	fake.ParentHandleStub = nil
	fake.parentHandleReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCreatedVolume) ParentHandleReturnsOnCall(i int, result1 string) {
	fake.parentHandleMutex.Lock()
	defer fake.parentHandleMutex.Unlock()
	fake.ParentHandleStub = nil
	if fake.parentHandleReturnsOnCall == nil {
		fake.parentHandleReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.parentHandleReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCreatedVolume) Path() string {
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

func (fake *FakeCreatedVolume) PathCallCount() int {
	fake.pathMutex.RLock()
	defer fake.pathMutex.RUnlock()
	return len(fake.pathArgsForCall)
}

func (fake *FakeCreatedVolume) PathCalls(stub func() string) {
	fake.pathMutex.Lock()
	defer fake.pathMutex.Unlock()
	fake.PathStub = stub
}

func (fake *FakeCreatedVolume) PathReturns(result1 string) {
	fake.pathMutex.Lock()
	defer fake.pathMutex.Unlock()
	fake.PathStub = nil
	fake.pathReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCreatedVolume) PathReturnsOnCall(i int, result1 string) {
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

func (fake *FakeCreatedVolume) ResourceType() (*db.VolumeResourceType, error) {
	fake.resourceTypeMutex.Lock()
	ret, specificReturn := fake.resourceTypeReturnsOnCall[len(fake.resourceTypeArgsForCall)]
	fake.resourceTypeArgsForCall = append(fake.resourceTypeArgsForCall, struct {
	}{})
	stub := fake.ResourceTypeStub
	fakeReturns := fake.resourceTypeReturns
	fake.recordInvocation("ResourceType", []interface{}{})
	fake.resourceTypeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCreatedVolume) ResourceTypeCallCount() int {
	fake.resourceTypeMutex.RLock()
	defer fake.resourceTypeMutex.RUnlock()
	return len(fake.resourceTypeArgsForCall)
}

func (fake *FakeCreatedVolume) ResourceTypeCalls(stub func() (*db.VolumeResourceType, error)) {
	fake.resourceTypeMutex.Lock()
	defer fake.resourceTypeMutex.Unlock()
	fake.ResourceTypeStub = stub
}

func (fake *FakeCreatedVolume) ResourceTypeReturns(result1 *db.VolumeResourceType, result2 error) {
	fake.resourceTypeMutex.Lock()
	defer fake.resourceTypeMutex.Unlock()
	fake.ResourceTypeStub = nil
	fake.resourceTypeReturns = struct {
		result1 *db.VolumeResourceType
		result2 error
	}{result1, result2}
}

func (fake *FakeCreatedVolume) ResourceTypeReturnsOnCall(i int, result1 *db.VolumeResourceType, result2 error) {
	fake.resourceTypeMutex.Lock()
	defer fake.resourceTypeMutex.Unlock()
	fake.ResourceTypeStub = nil
	if fake.resourceTypeReturnsOnCall == nil {
		fake.resourceTypeReturnsOnCall = make(map[int]struct {
			result1 *db.VolumeResourceType
			result2 error
		})
	}
	fake.resourceTypeReturnsOnCall[i] = struct {
		result1 *db.VolumeResourceType
		result2 error
	}{result1, result2}
}

func (fake *FakeCreatedVolume) TaskIdentifier() (int, atc.PipelineRef, string, string, error) {
	fake.taskIdentifierMutex.Lock()
	ret, specificReturn := fake.taskIdentifierReturnsOnCall[len(fake.taskIdentifierArgsForCall)]
	fake.taskIdentifierArgsForCall = append(fake.taskIdentifierArgsForCall, struct {
	}{})
	stub := fake.TaskIdentifierStub
	fakeReturns := fake.taskIdentifierReturns
	fake.recordInvocation("TaskIdentifier", []interface{}{})
	fake.taskIdentifierMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4, ret.result5
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3, fakeReturns.result4, fakeReturns.result5
}

func (fake *FakeCreatedVolume) TaskIdentifierCallCount() int {
	fake.taskIdentifierMutex.RLock()
	defer fake.taskIdentifierMutex.RUnlock()
	return len(fake.taskIdentifierArgsForCall)
}

func (fake *FakeCreatedVolume) TaskIdentifierCalls(stub func() (int, atc.PipelineRef, string, string, error)) {
	fake.taskIdentifierMutex.Lock()
	defer fake.taskIdentifierMutex.Unlock()
	fake.TaskIdentifierStub = stub
}

func (fake *FakeCreatedVolume) TaskIdentifierReturns(result1 int, result2 atc.PipelineRef, result3 string, result4 string, result5 error) {
	fake.taskIdentifierMutex.Lock()
	defer fake.taskIdentifierMutex.Unlock()
	fake.TaskIdentifierStub = nil
	fake.taskIdentifierReturns = struct {
		result1 int
		result2 atc.PipelineRef
		result3 string
		result4 string
		result5 error
	}{result1, result2, result3, result4, result5}
}

func (fake *FakeCreatedVolume) TaskIdentifierReturnsOnCall(i int, result1 int, result2 atc.PipelineRef, result3 string, result4 string, result5 error) {
	fake.taskIdentifierMutex.Lock()
	defer fake.taskIdentifierMutex.Unlock()
	fake.TaskIdentifierStub = nil
	if fake.taskIdentifierReturnsOnCall == nil {
		fake.taskIdentifierReturnsOnCall = make(map[int]struct {
			result1 int
			result2 atc.PipelineRef
			result3 string
			result4 string
			result5 error
		})
	}
	fake.taskIdentifierReturnsOnCall[i] = struct {
		result1 int
		result2 atc.PipelineRef
		result3 string
		result4 string
		result5 error
	}{result1, result2, result3, result4, result5}
}

func (fake *FakeCreatedVolume) TeamID() int {
	fake.teamIDMutex.Lock()
	ret, specificReturn := fake.teamIDReturnsOnCall[len(fake.teamIDArgsForCall)]
	fake.teamIDArgsForCall = append(fake.teamIDArgsForCall, struct {
	}{})
	stub := fake.TeamIDStub
	fakeReturns := fake.teamIDReturns
	fake.recordInvocation("TeamID", []interface{}{})
	fake.teamIDMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCreatedVolume) TeamIDCallCount() int {
	fake.teamIDMutex.RLock()
	defer fake.teamIDMutex.RUnlock()
	return len(fake.teamIDArgsForCall)
}

func (fake *FakeCreatedVolume) TeamIDCalls(stub func() int) {
	fake.teamIDMutex.Lock()
	defer fake.teamIDMutex.Unlock()
	fake.TeamIDStub = stub
}

func (fake *FakeCreatedVolume) TeamIDReturns(result1 int) {
	fake.teamIDMutex.Lock()
	defer fake.teamIDMutex.Unlock()
	fake.TeamIDStub = nil
	fake.teamIDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeCreatedVolume) TeamIDReturnsOnCall(i int, result1 int) {
	fake.teamIDMutex.Lock()
	defer fake.teamIDMutex.Unlock()
	fake.TeamIDStub = nil
	if fake.teamIDReturnsOnCall == nil {
		fake.teamIDReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.teamIDReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeCreatedVolume) Type() db.VolumeType {
	fake.typeMutex.Lock()
	ret, specificReturn := fake.typeReturnsOnCall[len(fake.typeArgsForCall)]
	fake.typeArgsForCall = append(fake.typeArgsForCall, struct {
	}{})
	stub := fake.TypeStub
	fakeReturns := fake.typeReturns
	fake.recordInvocation("Type", []interface{}{})
	fake.typeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCreatedVolume) TypeCallCount() int {
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	return len(fake.typeArgsForCall)
}

func (fake *FakeCreatedVolume) TypeCalls(stub func() db.VolumeType) {
	fake.typeMutex.Lock()
	defer fake.typeMutex.Unlock()
	fake.TypeStub = stub
}

func (fake *FakeCreatedVolume) TypeReturns(result1 db.VolumeType) {
	fake.typeMutex.Lock()
	defer fake.typeMutex.Unlock()
	fake.TypeStub = nil
	fake.typeReturns = struct {
		result1 db.VolumeType
	}{result1}
}

func (fake *FakeCreatedVolume) TypeReturnsOnCall(i int, result1 db.VolumeType) {
	fake.typeMutex.Lock()
	defer fake.typeMutex.Unlock()
	fake.TypeStub = nil
	if fake.typeReturnsOnCall == nil {
		fake.typeReturnsOnCall = make(map[int]struct {
			result1 db.VolumeType
		})
	}
	fake.typeReturnsOnCall[i] = struct {
		result1 db.VolumeType
	}{result1}
}

func (fake *FakeCreatedVolume) WorkerArtifactID() int {
	fake.workerArtifactIDMutex.Lock()
	ret, specificReturn := fake.workerArtifactIDReturnsOnCall[len(fake.workerArtifactIDArgsForCall)]
	fake.workerArtifactIDArgsForCall = append(fake.workerArtifactIDArgsForCall, struct {
	}{})
	stub := fake.WorkerArtifactIDStub
	fakeReturns := fake.workerArtifactIDReturns
	fake.recordInvocation("WorkerArtifactID", []interface{}{})
	fake.workerArtifactIDMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCreatedVolume) WorkerArtifactIDCallCount() int {
	fake.workerArtifactIDMutex.RLock()
	defer fake.workerArtifactIDMutex.RUnlock()
	return len(fake.workerArtifactIDArgsForCall)
}

func (fake *FakeCreatedVolume) WorkerArtifactIDCalls(stub func() int) {
	fake.workerArtifactIDMutex.Lock()
	defer fake.workerArtifactIDMutex.Unlock()
	fake.WorkerArtifactIDStub = stub
}

func (fake *FakeCreatedVolume) WorkerArtifactIDReturns(result1 int) {
	fake.workerArtifactIDMutex.Lock()
	defer fake.workerArtifactIDMutex.Unlock()
	fake.WorkerArtifactIDStub = nil
	fake.workerArtifactIDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeCreatedVolume) WorkerArtifactIDReturnsOnCall(i int, result1 int) {
	fake.workerArtifactIDMutex.Lock()
	defer fake.workerArtifactIDMutex.Unlock()
	fake.WorkerArtifactIDStub = nil
	if fake.workerArtifactIDReturnsOnCall == nil {
		fake.workerArtifactIDReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.workerArtifactIDReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeCreatedVolume) WorkerName() string {
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

func (fake *FakeCreatedVolume) WorkerNameCallCount() int {
	fake.workerNameMutex.RLock()
	defer fake.workerNameMutex.RUnlock()
	return len(fake.workerNameArgsForCall)
}

func (fake *FakeCreatedVolume) WorkerNameCalls(stub func() string) {
	fake.workerNameMutex.Lock()
	defer fake.workerNameMutex.Unlock()
	fake.WorkerNameStub = stub
}

func (fake *FakeCreatedVolume) WorkerNameReturns(result1 string) {
	fake.workerNameMutex.Lock()
	defer fake.workerNameMutex.Unlock()
	fake.WorkerNameStub = nil
	fake.workerNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCreatedVolume) WorkerNameReturnsOnCall(i int, result1 string) {
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

func (fake *FakeCreatedVolume) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.baseResourceTypeMutex.RLock()
	defer fake.baseResourceTypeMutex.RUnlock()
	fake.containerHandleMutex.RLock()
	defer fake.containerHandleMutex.RUnlock()
	fake.createChildForContainerMutex.RLock()
	defer fake.createChildForContainerMutex.RUnlock()
	fake.destroyingMutex.RLock()
	defer fake.destroyingMutex.RUnlock()
	fake.getResourceCacheIDMutex.RLock()
	defer fake.getResourceCacheIDMutex.RUnlock()
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	fake.initializeArtifactMutex.RLock()
	defer fake.initializeArtifactMutex.RUnlock()
	fake.initializeResourceCacheMutex.RLock()
	defer fake.initializeResourceCacheMutex.RUnlock()
	fake.initializeTaskCacheMutex.RLock()
	defer fake.initializeTaskCacheMutex.RUnlock()
	fake.parentHandleMutex.RLock()
	defer fake.parentHandleMutex.RUnlock()
	fake.pathMutex.RLock()
	defer fake.pathMutex.RUnlock()
	fake.resourceTypeMutex.RLock()
	defer fake.resourceTypeMutex.RUnlock()
	fake.taskIdentifierMutex.RLock()
	defer fake.taskIdentifierMutex.RUnlock()
	fake.teamIDMutex.RLock()
	defer fake.teamIDMutex.RUnlock()
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	fake.workerArtifactIDMutex.RLock()
	defer fake.workerArtifactIDMutex.RUnlock()
	fake.workerNameMutex.RLock()
	defer fake.workerNameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCreatedVolume) recordInvocation(key string, args []interface{}) {
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

var _ db.CreatedVolume = new(FakeCreatedVolume)
