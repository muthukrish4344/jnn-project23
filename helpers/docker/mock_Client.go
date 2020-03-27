// Code generated by mockery v1.0.0. DO NOT EDIT.

package docker_helpers

import (
	context "context"

	container "github.com/docker/docker/api/types/container"

	io "io"

	mock "github.com/stretchr/testify/mock"

	network "github.com/docker/docker/api/types/network"

	types "github.com/docker/docker/api/types"
)

// MockClient is an autogenerated mock type for the Client type
type MockClient struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *MockClient) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ContainerAttach provides a mock function with given fields: ctx, _a1, options
func (_m *MockClient) ContainerAttach(ctx context.Context, _a1 string, options types.ContainerAttachOptions) (types.HijackedResponse, error) {
	ret := _m.Called(ctx, _a1, options)

	var r0 types.HijackedResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ContainerAttachOptions) types.HijackedResponse); ok {
		r0 = rf(ctx, _a1, options)
	} else {
		r0 = ret.Get(0).(types.HijackedResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.ContainerAttachOptions) error); ok {
		r1 = rf(ctx, _a1, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainerCreate provides a mock function with given fields: ctx, config, hostConfig, networkingConfig, containerName
func (_m *MockClient) ContainerCreate(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, containerName string) (container.ContainerCreateCreatedBody, error) {
	ret := _m.Called(ctx, config, hostConfig, networkingConfig, containerName)

	var r0 container.ContainerCreateCreatedBody
	if rf, ok := ret.Get(0).(func(context.Context, *container.Config, *container.HostConfig, *network.NetworkingConfig, string) container.ContainerCreateCreatedBody); ok {
		r0 = rf(ctx, config, hostConfig, networkingConfig, containerName)
	} else {
		r0 = ret.Get(0).(container.ContainerCreateCreatedBody)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *container.Config, *container.HostConfig, *network.NetworkingConfig, string) error); ok {
		r1 = rf(ctx, config, hostConfig, networkingConfig, containerName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainerExecAttach provides a mock function with given fields: ctx, execID, config
func (_m *MockClient) ContainerExecAttach(ctx context.Context, execID string, config types.ExecStartCheck) (types.HijackedResponse, error) {
	ret := _m.Called(ctx, execID, config)

	var r0 types.HijackedResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ExecStartCheck) types.HijackedResponse); ok {
		r0 = rf(ctx, execID, config)
	} else {
		r0 = ret.Get(0).(types.HijackedResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.ExecStartCheck) error); ok {
		r1 = rf(ctx, execID, config)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainerExecCreate provides a mock function with given fields: ctx, _a1, config
func (_m *MockClient) ContainerExecCreate(ctx context.Context, _a1 string, config types.ExecConfig) (types.IDResponse, error) {
	ret := _m.Called(ctx, _a1, config)

	var r0 types.IDResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ExecConfig) types.IDResponse); ok {
		r0 = rf(ctx, _a1, config)
	} else {
		r0 = ret.Get(0).(types.IDResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.ExecConfig) error); ok {
		r1 = rf(ctx, _a1, config)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainerInspect provides a mock function with given fields: ctx, containerID
func (_m *MockClient) ContainerInspect(ctx context.Context, containerID string) (types.ContainerJSON, error) {
	ret := _m.Called(ctx, containerID)

	var r0 types.ContainerJSON
	if rf, ok := ret.Get(0).(func(context.Context, string) types.ContainerJSON); ok {
		r0 = rf(ctx, containerID)
	} else {
		r0 = ret.Get(0).(types.ContainerJSON)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, containerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainerKill provides a mock function with given fields: ctx, containerID, signal
func (_m *MockClient) ContainerKill(ctx context.Context, containerID string, signal string) error {
	ret := _m.Called(ctx, containerID, signal)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, containerID, signal)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ContainerLogs provides a mock function with given fields: ctx, _a1, options
func (_m *MockClient) ContainerLogs(ctx context.Context, _a1 string, options types.ContainerLogsOptions) (io.ReadCloser, error) {
	ret := _m.Called(ctx, _a1, options)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ContainerLogsOptions) io.ReadCloser); ok {
		r0 = rf(ctx, _a1, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.ContainerLogsOptions) error); ok {
		r1 = rf(ctx, _a1, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainerRemove provides a mock function with given fields: ctx, containerID, options
func (_m *MockClient) ContainerRemove(ctx context.Context, containerID string, options types.ContainerRemoveOptions) error {
	ret := _m.Called(ctx, containerID, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ContainerRemoveOptions) error); ok {
		r0 = rf(ctx, containerID, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ContainerStart provides a mock function with given fields: ctx, containerID, options
func (_m *MockClient) ContainerStart(ctx context.Context, containerID string, options types.ContainerStartOptions) error {
	ret := _m.Called(ctx, containerID, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ContainerStartOptions) error); ok {
		r0 = rf(ctx, containerID, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ImageImportBlocking provides a mock function with given fields: ctx, source, ref, options
func (_m *MockClient) ImageImportBlocking(ctx context.Context, source types.ImageImportSource, ref string, options types.ImageImportOptions) error {
	ret := _m.Called(ctx, source, ref, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.ImageImportSource, string, types.ImageImportOptions) error); ok {
		r0 = rf(ctx, source, ref, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ImageInspectWithRaw provides a mock function with given fields: ctx, imageID
func (_m *MockClient) ImageInspectWithRaw(ctx context.Context, imageID string) (types.ImageInspect, []byte, error) {
	ret := _m.Called(ctx, imageID)

	var r0 types.ImageInspect
	if rf, ok := ret.Get(0).(func(context.Context, string) types.ImageInspect); ok {
		r0 = rf(ctx, imageID)
	} else {
		r0 = ret.Get(0).(types.ImageInspect)
	}

	var r1 []byte
	if rf, ok := ret.Get(1).(func(context.Context, string) []byte); ok {
		r1 = rf(ctx, imageID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, imageID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ImagePullBlocking provides a mock function with given fields: ctx, ref, options
func (_m *MockClient) ImagePullBlocking(ctx context.Context, ref string, options types.ImagePullOptions) error {
	ret := _m.Called(ctx, ref, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ImagePullOptions) error); ok {
		r0 = rf(ctx, ref, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Info provides a mock function with given fields: ctx
func (_m *MockClient) Info(ctx context.Context) (types.Info, error) {
	ret := _m.Called(ctx)

	var r0 types.Info
	if rf, ok := ret.Get(0).(func(context.Context) types.Info); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(types.Info)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NetworkCreate provides a mock function with given fields: ctx, networkName, options
func (_m *MockClient) NetworkCreate(ctx context.Context, networkName string, options types.NetworkCreate) (types.NetworkCreateResponse, error) {
	ret := _m.Called(ctx, networkName, options)

	var r0 types.NetworkCreateResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, types.NetworkCreate) types.NetworkCreateResponse); ok {
		r0 = rf(ctx, networkName, options)
	} else {
		r0 = ret.Get(0).(types.NetworkCreateResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.NetworkCreate) error); ok {
		r1 = rf(ctx, networkName, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NetworkDisconnect provides a mock function with given fields: ctx, networkID, containerID, force
func (_m *MockClient) NetworkDisconnect(ctx context.Context, networkID string, containerID string, force bool) error {
	ret := _m.Called(ctx, networkID, containerID, force)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, bool) error); ok {
		r0 = rf(ctx, networkID, containerID, force)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NetworkInspect provides a mock function with given fields: ctx, networkID
func (_m *MockClient) NetworkInspect(ctx context.Context, networkID string) (types.NetworkResource, error) {
	ret := _m.Called(ctx, networkID)

	var r0 types.NetworkResource
	if rf, ok := ret.Get(0).(func(context.Context, string) types.NetworkResource); ok {
		r0 = rf(ctx, networkID)
	} else {
		r0 = ret.Get(0).(types.NetworkResource)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, networkID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NetworkList provides a mock function with given fields: ctx, options
func (_m *MockClient) NetworkList(ctx context.Context, options types.NetworkListOptions) ([]types.NetworkResource, error) {
	ret := _m.Called(ctx, options)

	var r0 []types.NetworkResource
	if rf, ok := ret.Get(0).(func(context.Context, types.NetworkListOptions) []types.NetworkResource); ok {
		r0 = rf(ctx, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.NetworkResource)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.NetworkListOptions) error); ok {
		r1 = rf(ctx, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NetworkRemove provides a mock function with given fields: ctx, networkID
func (_m *MockClient) NetworkRemove(ctx context.Context, networkID string) error {
	ret := _m.Called(ctx, networkID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, networkID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
