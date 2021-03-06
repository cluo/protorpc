// Code generated by protoc-gen-go.
// source: arith.proto
// DO NOT EDIT!

/*
Package service is a generated protocol buffer package.

It is generated from these files:
	arith.proto
	echo.proto

It has these top-level messages:
	ArithRequest
	ArithResponse
*/
package service

import proto "github.com/chai2010/protorpc/proto"
import math "math"

import "io"
import "log"
import "net"
import "net/rpc"
import "time"
import protorpc "github.com/chai2010/protorpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type ArithRequest struct {
	A                *int32 `protobuf:"varint,1,opt,name=a" json:"a,omitempty"`
	B                *int32 `protobuf:"varint,2,opt,name=b" json:"b,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ArithRequest) Reset()         { *m = ArithRequest{} }
func (m *ArithRequest) String() string { return proto.CompactTextString(m) }
func (*ArithRequest) ProtoMessage()    {}

func (m *ArithRequest) GetA() int32 {
	if m != nil && m.A != nil {
		return *m.A
	}
	return 0
}

func (m *ArithRequest) GetB() int32 {
	if m != nil && m.B != nil {
		return *m.B
	}
	return 0
}

type ArithResponse struct {
	C                *int32 `protobuf:"varint,1,opt,name=c" json:"c,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ArithResponse) Reset()         { *m = ArithResponse{} }
func (m *ArithResponse) String() string { return proto.CompactTextString(m) }
func (*ArithResponse) ProtoMessage()    {}

func (m *ArithResponse) GetC() int32 {
	if m != nil && m.C != nil {
		return *m.C
	}
	return 0
}

func init() {
}

type ArithService interface {
	Add(in *ArithRequest, out *ArithResponse) error
	Mul(in *ArithRequest, out *ArithResponse) error
	Div(in *ArithRequest, out *ArithResponse) error
	Error(in *ArithRequest, out *ArithResponse) error
}

// AcceptArithServiceClient accepts connections on the listener and serves requests
// for each incoming connection.  Accept blocks; the caller typically
// invokes it in a go statement.
func AcceptArithServiceClient(lis net.Listener, x ArithService) {
	srv := rpc.NewServer()
	if err := srv.RegisterName("ArithService", x); err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("lis.Accept(): %v\n", err)
		}
		go srv.ServeCodec(protorpc.NewServerCodec(conn))
	}
}

// RegisterArithService publish the given ArithService implementation on the server.
func RegisterArithService(srv *rpc.Server, x ArithService) error {
	if err := srv.RegisterName("ArithService", x); err != nil {
		return err
	}
	return nil
}

// NewArithServiceServer returns a new ArithService Server.
func NewArithServiceServer(x ArithService) *rpc.Server {
	srv := rpc.NewServer()
	if err := srv.RegisterName("ArithService", x); err != nil {
		log.Fatal(err)
	}
	return srv
}

// ListenAndServeArithService listen announces on the local network address laddr
// and serves the given ArithService implementation.
func ListenAndServeArithService(network, addr string, x ArithService) error {
	lis, err := net.Listen(network, addr)
	if err != nil {
		return err
	}
	defer lis.Close()

	srv := rpc.NewServer()
	if err := srv.RegisterName("ArithService", x); err != nil {
		return err
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("lis.Accept(): %v\n", err)
		}
		go srv.ServeCodec(protorpc.NewServerCodec(conn))
	}
}

type ArithServiceClient struct {
	*rpc.Client
}

// NewArithServiceClient returns a ArithService rpc.Client and stub to handle
// requests to the set of ArithService at the other end of the connection.
func NewArithServiceClient(conn io.ReadWriteCloser) (*ArithServiceClient, *rpc.Client) {
	c := rpc.NewClientWithCodec(protorpc.NewClientCodec(conn))
	return &ArithServiceClient{c}, c
}

func (c *ArithServiceClient) Add(in *ArithRequest, out *ArithResponse) error {
	return c.Call("ArithService.Add", in, out)
}
func (c *ArithServiceClient) Mul(in *ArithRequest, out *ArithResponse) error {
	return c.Call("ArithService.Mul", in, out)
}
func (c *ArithServiceClient) Div(in *ArithRequest, out *ArithResponse) error {
	return c.Call("ArithService.Div", in, out)
}
func (c *ArithServiceClient) Error(in *ArithRequest, out *ArithResponse) error {
	return c.Call("ArithService.Error", in, out)
}

// DialArithService connects to an ArithService at the specified network address.
func DialArithService(network, addr string) (*ArithServiceClient, *rpc.Client, error) {
	c, err := protorpc.Dial(network, addr)
	if err != nil {
		return nil, nil, err
	}
	return &ArithServiceClient{c}, c, nil
}

// DialArithServiceTimeout connects to an ArithService at the specified network address.
func DialArithServiceTimeout(network, addr string,
	timeout time.Duration) (*ArithServiceClient, *rpc.Client, error) {
	c, err := protorpc.DialTimeout(network, addr, timeout)
	if err != nil {
		return nil, nil, err
	}
	return &ArithServiceClient{c}, c, nil
}
