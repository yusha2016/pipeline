// +build !ignore_autogenerated

// Code generated by mga tool. DO NOT EDIT.

package pke

import (
	"context"
	"github.com/banzaicloud/pipeline/internal/cluster"
	"github.com/banzaicloud/pipeline/internal/cluster/distribution/awscommon"
	"github.com/stretchr/testify/mock"
)

// MockNodePoolManager is an autogenerated mock for the NodePoolManager type.
type MockNodePoolManager struct {
	mock.Mock
}

// ListNodePools provides a mock function.
func (_m *MockNodePoolManager) ListNodePools(ctx context.Context, c cluster.Cluster, existingNodePools map[string]awscommon.ExistingNodePool) (_result_0 []NodePool, _result_1 error) {
	ret := _m.Called(ctx, c, existingNodePools)

	var r0 []NodePool
	if rf, ok := ret.Get(0).(func(context.Context, cluster.Cluster, map[string]awscommon.ExistingNodePool) []NodePool); ok {
		r0 = rf(ctx, c, existingNodePools)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]NodePool)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, cluster.Cluster, map[string]awscommon.ExistingNodePool) error); ok {
		r1 = rf(ctx, c, existingNodePools)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateNodePool provides a mock function.
func (_m *MockNodePoolManager) UpdateNodePool(ctx context.Context, c cluster.Cluster, nodePoolName string, nodePoolUpdate NodePoolUpdate) (_result_0 string, _result_1 error) {
	ret := _m.Called(ctx, c, nodePoolName, nodePoolUpdate)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, cluster.Cluster, string, NodePoolUpdate) string); ok {
		r0 = rf(ctx, c, nodePoolName, nodePoolUpdate)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, cluster.Cluster, string, NodePoolUpdate) error); ok {
		r1 = rf(ctx, c, nodePoolName, nodePoolUpdate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStore is an autogenerated mock for the Store type.
type MockStore struct {
	mock.Mock
}

// GetCluster provides a mock function.
func (_m *MockStore) GetCluster(ctx context.Context, id uint) (_result_0 cluster.Cluster, _result_1 error) {
	ret := _m.Called(ctx, id)

	var r0 cluster.Cluster
	if rf, ok := ret.Get(0).(func(context.Context, uint) cluster.Cluster); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(cluster.Cluster)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetStatus provides a mock function.
func (_m *MockStore) SetStatus(ctx context.Context, id uint, status string, statusMessage string) (_result_0 error) {
	ret := _m.Called(ctx, id, status, statusMessage)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint, string, string) error); ok {
		r0 = rf(ctx, id, status, statusMessage)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
