// Code generated by protoc-gen-stdrpc. DO NOT EDIT.
//
// plugin: https://github.com/chai2010/protorpc/tree/master/protoc-gen-plugin
// plugin: https://github.com/chai2010/protorpc/tree/master/protoc-gen-stdrpc
//
// source: frontgate.proto

package openpitrix_frontgate

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/rpc"
	"time"

	"github.com/golang/protobuf/proto"
)

var (
	_ = fmt.Sprint
	_ = io.Reader(nil)
	_ = log.Print
	_ = net.Addr(nil)
	_ = rpc.Call{}
	_ = time.Second

	_ = proto.String
)

type FrontgateService interface {
	StartConfd(in *Task, out *Empty) error
	RegisterMetadata(in *Task, out *Empty) error
	DeregisterMetadata(in *Task, out *Empty) error
	RegisterCmd(in *Task, out *Empty) error
	DeregisterCmd(in *Task, out *Empty) error
}

// AcceptFrontgateServiceClient accepts connections on the listener and serves requests
// for each incoming connection.  Accept blocks; the caller typically
// invokes it in a go statement.
func AcceptFrontgateServiceClient(lis net.Listener, x FrontgateService) {
	srv := rpc.NewServer()
	if err := srv.RegisterName("FrontgateService", x); err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("lis.Accept(): %v\n", err)
		}
		go srv.ServeConn(conn)
	}
}

// RegisterFrontgateService publish the given FrontgateService implementation on the server.
func RegisterFrontgateService(srv *rpc.Server, x FrontgateService) error {
	if err := srv.RegisterName("FrontgateService", x); err != nil {
		return err
	}
	return nil
}

// NewFrontgateServiceServer returns a new FrontgateService Server.
func NewFrontgateServiceServer(x FrontgateService) *rpc.Server {
	srv := rpc.NewServer()
	if err := srv.RegisterName("FrontgateService", x); err != nil {
		log.Fatal(err)
	}
	return srv
}

// ListenAndServeFrontgateService listen announces on the local network address laddr
// and serves the given FrontgateService implementation.
func ListenAndServeFrontgateService(network, addr string, x FrontgateService) error {
	lis, err := net.Listen(network, addr)
	if err != nil {
		return err
	}
	defer lis.Close()

	srv := rpc.NewServer()
	if err := srv.RegisterName("FrontgateService", x); err != nil {
		return err
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("lis.Accept(): %v\n", err)
		}
		go srv.ServeConn(conn)
	}
}

// ServeFrontgateService serves the given FrontgateService implementation.
func ServeFrontgateService(conn io.ReadWriteCloser, x FrontgateService) {
	srv := rpc.NewServer()
	if err := srv.RegisterName("FrontgateService", x); err != nil {
		log.Fatal(err)
	}
	srv.ServeConn(conn)
}

type FrontgateServiceClient struct {
	*rpc.Client
}

// NewFrontgateServiceClient returns a FrontgateService stub to handle
// requests to the set of FrontgateService at the other end of the connection.
func NewFrontgateServiceClient(conn io.ReadWriteCloser) *FrontgateServiceClient {
	c := rpc.NewClient(conn)
	return &FrontgateServiceClient{c}
}

func (c *FrontgateServiceClient) StartConfd(in *Task) (out *Empty, err error) {
	if in == nil {
		in = new(Task)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(Empty)
	if err = c.Call("FrontgateService.StartConfd", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *FrontgateServiceClient) AsyncStartConfd(in *Task, out *Empty, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(Task)
	}
	return c.Go(
		"FrontgateService.StartConfd",
		in, out,
		done,
	)
}

func (c *FrontgateServiceClient) RegisterMetadata(in *Task) (out *Empty, err error) {
	if in == nil {
		in = new(Task)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(Empty)
	if err = c.Call("FrontgateService.RegisterMetadata", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *FrontgateServiceClient) AsyncRegisterMetadata(in *Task, out *Empty, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(Task)
	}
	return c.Go(
		"FrontgateService.RegisterMetadata",
		in, out,
		done,
	)
}

func (c *FrontgateServiceClient) DeregisterMetadata(in *Task) (out *Empty, err error) {
	if in == nil {
		in = new(Task)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(Empty)
	if err = c.Call("FrontgateService.DeregisterMetadata", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *FrontgateServiceClient) AsyncDeregisterMetadata(in *Task, out *Empty, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(Task)
	}
	return c.Go(
		"FrontgateService.DeregisterMetadata",
		in, out,
		done,
	)
}

func (c *FrontgateServiceClient) RegisterCmd(in *Task) (out *Empty, err error) {
	if in == nil {
		in = new(Task)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(Empty)
	if err = c.Call("FrontgateService.RegisterCmd", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *FrontgateServiceClient) AsyncRegisterCmd(in *Task, out *Empty, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(Task)
	}
	return c.Go(
		"FrontgateService.RegisterCmd",
		in, out,
		done,
	)
}

func (c *FrontgateServiceClient) DeregisterCmd(in *Task) (out *Empty, err error) {
	if in == nil {
		in = new(Task)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(Empty)
	if err = c.Call("FrontgateService.DeregisterCmd", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *FrontgateServiceClient) AsyncDeregisterCmd(in *Task, out *Empty, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(Task)
	}
	return c.Go(
		"FrontgateService.DeregisterCmd",
		in, out,
		done,
	)
}

// DialFrontgateService connects to an FrontgateService at the specified network address.
func DialFrontgateService(network, addr string) (*FrontgateServiceClient, error) {
	c, err := rpc.Dial(network, addr)
	if err != nil {
		return nil, err
	}
	return &FrontgateServiceClient{c}, nil
}

// DialFrontgateServiceTimeout connects to an FrontgateService at the specified network address.
func DialFrontgateServiceTimeout(network, addr string, timeout time.Duration) (*FrontgateServiceClient, error) {
	conn, err := net.DialTimeout(network, addr, timeout)
	if err != nil {
		return nil, err
	}
	return &FrontgateServiceClient{rpc.NewClient(conn)}, nil
}