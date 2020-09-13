// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package snackinventory

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// SnackInventoryClient is the client API for SnackInventory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SnackInventoryClient interface {
	CreateSnack(ctx context.Context, in *CreateSnackRequest, opts ...grpc.CallOption) (*CreateSnackResponse, error)
	ListSnacks(ctx context.Context, in *ListSnacksRequest, opts ...grpc.CallOption) (*ListSnacksResponse, error)
	UpdateSnack(ctx context.Context, in *UpdateSnackRequest, opts ...grpc.CallOption) (*UpdateSnackResponse, error)
	DeleteSnack(ctx context.Context, in *DeleteSnackRequest, opts ...grpc.CallOption) (*DeleteSnackResponse, error)
	CreateLocation(ctx context.Context, in *CreateLocationRequest, opts ...grpc.CallOption) (*CreateLocationResponse, error)
	ListLocations(ctx context.Context, in *ListLocationsRequest, opts ...grpc.CallOption) (*ListLocationsResponse, error)
	DeleteLocation(ctx context.Context, in *DeleteLocationRequest, opts ...grpc.CallOption) (*DeleteLocationResponse, error)
	AddSnack(ctx context.Context, in *AddSnackRequest, opts ...grpc.CallOption) (*AddSnackResponse, error)
}

type snackInventoryClient struct {
	cc grpc.ClientConnInterface
}

func NewSnackInventoryClient(cc grpc.ClientConnInterface) SnackInventoryClient {
	return &snackInventoryClient{cc}
}

var snackInventoryCreateSnackStreamDesc = &grpc.StreamDesc{
	StreamName: "CreateSnack",
}

func (c *snackInventoryClient) CreateSnack(ctx context.Context, in *CreateSnackRequest, opts ...grpc.CallOption) (*CreateSnackResponse, error) {
	out := new(CreateSnackResponse)
	err := c.cc.Invoke(ctx, "/snackinventory.SnackInventory/CreateSnack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var snackInventoryListSnacksStreamDesc = &grpc.StreamDesc{
	StreamName: "ListSnacks",
}

func (c *snackInventoryClient) ListSnacks(ctx context.Context, in *ListSnacksRequest, opts ...grpc.CallOption) (*ListSnacksResponse, error) {
	out := new(ListSnacksResponse)
	err := c.cc.Invoke(ctx, "/snackinventory.SnackInventory/ListSnacks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var snackInventoryUpdateSnackStreamDesc = &grpc.StreamDesc{
	StreamName: "updateSnack",
}

func (c *snackInventoryClient) UpdateSnack(ctx context.Context, in *UpdateSnackRequest, opts ...grpc.CallOption) (*UpdateSnackResponse, error) {
	out := new(UpdateSnackResponse)
	err := c.cc.Invoke(ctx, "/snackinventory.SnackInventory/updateSnack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var snackInventoryDeleteSnackStreamDesc = &grpc.StreamDesc{
	StreamName: "DeleteSnack",
}

func (c *snackInventoryClient) DeleteSnack(ctx context.Context, in *DeleteSnackRequest, opts ...grpc.CallOption) (*DeleteSnackResponse, error) {
	out := new(DeleteSnackResponse)
	err := c.cc.Invoke(ctx, "/snackinventory.SnackInventory/DeleteSnack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var snackInventoryCreateLocationStreamDesc = &grpc.StreamDesc{
	StreamName: "CreateLocation",
}

func (c *snackInventoryClient) CreateLocation(ctx context.Context, in *CreateLocationRequest, opts ...grpc.CallOption) (*CreateLocationResponse, error) {
	out := new(CreateLocationResponse)
	err := c.cc.Invoke(ctx, "/snackinventory.SnackInventory/CreateLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var snackInventoryListLocationsStreamDesc = &grpc.StreamDesc{
	StreamName: "ListLocations",
}

func (c *snackInventoryClient) ListLocations(ctx context.Context, in *ListLocationsRequest, opts ...grpc.CallOption) (*ListLocationsResponse, error) {
	out := new(ListLocationsResponse)
	err := c.cc.Invoke(ctx, "/snackinventory.SnackInventory/ListLocations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var snackInventoryDeleteLocationStreamDesc = &grpc.StreamDesc{
	StreamName: "DeleteLocation",
}

func (c *snackInventoryClient) DeleteLocation(ctx context.Context, in *DeleteLocationRequest, opts ...grpc.CallOption) (*DeleteLocationResponse, error) {
	out := new(DeleteLocationResponse)
	err := c.cc.Invoke(ctx, "/snackinventory.SnackInventory/DeleteLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var snackInventoryAddSnackStreamDesc = &grpc.StreamDesc{
	StreamName: "AddSnack",
}

func (c *snackInventoryClient) AddSnack(ctx context.Context, in *AddSnackRequest, opts ...grpc.CallOption) (*AddSnackResponse, error) {
	out := new(AddSnackResponse)
	err := c.cc.Invoke(ctx, "/snackinventory.SnackInventory/AddSnack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SnackInventoryService is the service API for SnackInventory service.
// Fields should be assigned to their respective handler implementations only before
// RegisterSnackInventoryService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type SnackInventoryService struct {
	CreateSnack    func(context.Context, *CreateSnackRequest) (*CreateSnackResponse, error)
	ListSnacks     func(context.Context, *ListSnacksRequest) (*ListSnacksResponse, error)
	UpdateSnack    func(context.Context, *UpdateSnackRequest) (*UpdateSnackResponse, error)
	DeleteSnack    func(context.Context, *DeleteSnackRequest) (*DeleteSnackResponse, error)
	CreateLocation func(context.Context, *CreateLocationRequest) (*CreateLocationResponse, error)
	ListLocations  func(context.Context, *ListLocationsRequest) (*ListLocationsResponse, error)
	DeleteLocation func(context.Context, *DeleteLocationRequest) (*DeleteLocationResponse, error)
	AddSnack       func(context.Context, *AddSnackRequest) (*AddSnackResponse, error)
}

func (s *SnackInventoryService) createSnack(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSnackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.CreateSnack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/snackinventory.SnackInventory/CreateSnack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.CreateSnack(ctx, req.(*CreateSnackRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SnackInventoryService) listSnacks(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSnacksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.ListSnacks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/snackinventory.SnackInventory/ListSnacks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.ListSnacks(ctx, req.(*ListSnacksRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SnackInventoryService) updateSnack(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSnackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.UpdateSnack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/snackinventory.SnackInventory/UpdateSnack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.UpdateSnack(ctx, req.(*UpdateSnackRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SnackInventoryService) deleteSnack(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSnackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.DeleteSnack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/snackinventory.SnackInventory/DeleteSnack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.DeleteSnack(ctx, req.(*DeleteSnackRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SnackInventoryService) createLocation(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.CreateLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/snackinventory.SnackInventory/CreateLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.CreateLocation(ctx, req.(*CreateLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SnackInventoryService) listLocations(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLocationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.ListLocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/snackinventory.SnackInventory/ListLocations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.ListLocations(ctx, req.(*ListLocationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SnackInventoryService) deleteLocation(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.DeleteLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/snackinventory.SnackInventory/DeleteLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.DeleteLocation(ctx, req.(*DeleteLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SnackInventoryService) addSnack(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSnackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.AddSnack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/snackinventory.SnackInventory/AddSnack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.AddSnack(ctx, req.(*AddSnackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterSnackInventoryService registers a service implementation with a gRPC server.
func RegisterSnackInventoryService(s grpc.ServiceRegistrar, srv *SnackInventoryService) {
	srvCopy := *srv
	if srvCopy.CreateSnack == nil {
		srvCopy.CreateSnack = func(context.Context, *CreateSnackRequest) (*CreateSnackResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method CreateSnack not implemented")
		}
	}
	if srvCopy.ListSnacks == nil {
		srvCopy.ListSnacks = func(context.Context, *ListSnacksRequest) (*ListSnacksResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method ListSnacks not implemented")
		}
	}
	if srvCopy.UpdateSnack == nil {
		srvCopy.UpdateSnack = func(context.Context, *UpdateSnackRequest) (*UpdateSnackResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method UpdateSnack not implemented")
		}
	}
	if srvCopy.DeleteSnack == nil {
		srvCopy.DeleteSnack = func(context.Context, *DeleteSnackRequest) (*DeleteSnackResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method DeleteSnack not implemented")
		}
	}
	if srvCopy.CreateLocation == nil {
		srvCopy.CreateLocation = func(context.Context, *CreateLocationRequest) (*CreateLocationResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method CreateLocation not implemented")
		}
	}
	if srvCopy.ListLocations == nil {
		srvCopy.ListLocations = func(context.Context, *ListLocationsRequest) (*ListLocationsResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method ListLocations not implemented")
		}
	}
	if srvCopy.DeleteLocation == nil {
		srvCopy.DeleteLocation = func(context.Context, *DeleteLocationRequest) (*DeleteLocationResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method DeleteLocation not implemented")
		}
	}
	if srvCopy.AddSnack == nil {
		srvCopy.AddSnack = func(context.Context, *AddSnackRequest) (*AddSnackResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method AddSnack not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "snackinventory.SnackInventory",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "CreateSnack",
				Handler:    srvCopy.createSnack,
			},
			{
				MethodName: "ListSnacks",
				Handler:    srvCopy.listSnacks,
			},
			{
				MethodName: "updateSnack",
				Handler:    srvCopy.updateSnack,
			},
			{
				MethodName: "DeleteSnack",
				Handler:    srvCopy.deleteSnack,
			},
			{
				MethodName: "CreateLocation",
				Handler:    srvCopy.createLocation,
			},
			{
				MethodName: "ListLocations",
				Handler:    srvCopy.listLocations,
			},
			{
				MethodName: "DeleteLocation",
				Handler:    srvCopy.deleteLocation,
			},
			{
				MethodName: "AddSnack",
				Handler:    srvCopy.addSnack,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "snackinventory.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewSnackInventoryService creates a new SnackInventoryService containing the
// implemented methods of the SnackInventory service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewSnackInventoryService(s interface{}) *SnackInventoryService {
	ns := &SnackInventoryService{}
	if h, ok := s.(interface {
		CreateSnack(context.Context, *CreateSnackRequest) (*CreateSnackResponse, error)
	}); ok {
		ns.CreateSnack = h.CreateSnack
	}
	if h, ok := s.(interface {
		ListSnacks(context.Context, *ListSnacksRequest) (*ListSnacksResponse, error)
	}); ok {
		ns.ListSnacks = h.ListSnacks
	}
	if h, ok := s.(interface {
		UpdateSnack(context.Context, *UpdateSnackRequest) (*UpdateSnackResponse, error)
	}); ok {
		ns.UpdateSnack = h.UpdateSnack
	}
	if h, ok := s.(interface {
		DeleteSnack(context.Context, *DeleteSnackRequest) (*DeleteSnackResponse, error)
	}); ok {
		ns.DeleteSnack = h.DeleteSnack
	}
	if h, ok := s.(interface {
		CreateLocation(context.Context, *CreateLocationRequest) (*CreateLocationResponse, error)
	}); ok {
		ns.CreateLocation = h.CreateLocation
	}
	if h, ok := s.(interface {
		ListLocations(context.Context, *ListLocationsRequest) (*ListLocationsResponse, error)
	}); ok {
		ns.ListLocations = h.ListLocations
	}
	if h, ok := s.(interface {
		DeleteLocation(context.Context, *DeleteLocationRequest) (*DeleteLocationResponse, error)
	}); ok {
		ns.DeleteLocation = h.DeleteLocation
	}
	if h, ok := s.(interface {
		AddSnack(context.Context, *AddSnackRequest) (*AddSnackResponse, error)
	}); ok {
		ns.AddSnack = h.AddSnack
	}
	return ns
}

// UnstableSnackInventoryService is the service API for SnackInventory service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstableSnackInventoryService interface {
	CreateSnack(context.Context, *CreateSnackRequest) (*CreateSnackResponse, error)
	ListSnacks(context.Context, *ListSnacksRequest) (*ListSnacksResponse, error)
	UpdateSnack(context.Context, *UpdateSnackRequest) (*UpdateSnackResponse, error)
	DeleteSnack(context.Context, *DeleteSnackRequest) (*DeleteSnackResponse, error)
	CreateLocation(context.Context, *CreateLocationRequest) (*CreateLocationResponse, error)
	ListLocations(context.Context, *ListLocationsRequest) (*ListLocationsResponse, error)
	DeleteLocation(context.Context, *DeleteLocationRequest) (*DeleteLocationResponse, error)
	AddSnack(context.Context, *AddSnackRequest) (*AddSnackResponse, error)
}
