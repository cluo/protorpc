// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package service

import (
	"log"
	"net"
	"net/rpc"
	"testing"

	"github.com/chai2010/protorpc"
	"github.com/chai2010/protorpc/proto"
)

func init() {
	err := listenAndServeArithAndEchoService("tcp", ":1984")
	if err != nil {
		log.Fatalf("listenAndServeArithAndEchoService: %v", err)
	}
}

func TestAll(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:1984")
	if err != nil {
		t.Fatalf(`net.Dial("tcp", "127.0.0.1:1984"): %v`, err)
	}
	client := rpc.NewClientWithCodec(protorpc.NewClientCodec(conn))
	defer client.Close()

	testArithClient(t, client)
	testEchoClient(t, client)

	arithStub := &ArithServiceClient{client}
	echoStub := &EchoServiceClient{client}

	testArithStub(t, arithStub)
	testEchoStub(t, echoStub)
}

func listenAndServeArithAndEchoService(network, addr string) error {
	clients, err := net.Listen(network, addr)
	if err != nil {
		return err
	}
	srv := rpc.NewServer()
	if err := RegisterArithService(srv, new(Arith)); err != nil {
		return err
	}
	if err := RegisterEchoService(srv, new(Echo)); err != nil {
		return err
	}
	go func() {
		for {
			conn, err := clients.Accept()
			if err != nil {
				log.Printf("clients.Accept(): %v\n", err)
				continue
			}
			go srv.ServeCodec(protorpc.NewServerCodec(conn))
		}
	}()
	return nil
}

func testArithClient(t *testing.T, client *rpc.Client) {
	var args ArithRequest
	var reply ArithResponse
	var err error

	// Add
	args.A = proto.Int32(1)
	args.B = proto.Int32(2)
	if err = client.Call("ArithService.Add", &args, &reply); err != nil {
		t.Fatalf(`arith.Add: %v`, err)
	}
	if reply.GetC() != 3 {
		t.Fatalf(`arith.Add: expected = %d, got = %d`, 3, reply.GetC())
	}

	// Mul
	args.A = proto.Int32(2)
	args.B = proto.Int32(3)
	if err = client.Call("ArithService.Mul", &args, &reply); err != nil {
		t.Fatalf(`arith.Mul: %v`, err)
	}
	if reply.GetC() != 6 {
		t.Fatalf(`arith.Mul: expected = %d, got = %d`, 6, reply.GetC())
	}

	// Div
	args.A = proto.Int32(13)
	args.B = proto.Int32(5)
	if err = client.Call("ArithService.Div", &args, &reply); err != nil {
		t.Fatalf(`arith.Div: %v`, err)
	}
	if reply.GetC() != 2 {
		t.Fatalf(`arith.Div: expected = %d, got = %d`, 2, reply.GetC())
	}

	// Div zero
	args.A = proto.Int32(1)
	args.B = proto.Int32(0)
	if err = client.Call("ArithService.Div", &args, &reply); err.Error() != "divide by zero" {
		t.Fatalf(`arith.Div: expected = "%s", got = "%s"`, "divide by zero", err.Error())
	}

	// Error
	args.A = proto.Int32(1)
	args.B = proto.Int32(2)
	if err = client.Call("ArithService.Error", &args, &reply); err.Error() != "ArithError" {
		t.Fatalf(`arith.Error: expected = "%s", got = "%s"`, "ArithError", err.Error())
	}
}

func testEchoClient(t *testing.T, client *rpc.Client) {
	var args EchoRequest
	var reply EchoResponse
	var err error

	// EchoService.Echo
	args.Msg = proto.String("Hello, Protobuf-RPC")
	if err = client.Call("EchoService.Echo", &args, &reply); err != nil {
		t.Fatalf(`echo.Echo: %v`, err)
	}
	if reply.GetMsg() != args.GetMsg() {
		t.Fatalf(`echo.Echo: expected = "%s", got = "%s"`, args.GetMsg(), reply.GetMsg())
	}
}

func testArithStub(t *testing.T, stub ArithService) {
	var args ArithRequest
	var reply ArithResponse
	var err error

	// Add
	args.A = proto.Int32(1)
	args.B = proto.Int32(2)
	if err = stub.Add(&args, &reply); err != nil {
		t.Fatalf(`stub.Add: %v`, err)
	}
	if reply.GetC() != 3 {
		t.Fatalf(`stub.Add: expected = %d, got = %d`, 3, reply.GetC())
	}

	// Mul
	args.A = proto.Int32(2)
	args.B = proto.Int32(3)
	if err = stub.Mul(&args, &reply); err != nil {
		t.Fatalf(`stub.Mul: %v`, err)
	}
	if reply.GetC() != 6 {
		t.Fatalf(`stub.Mul: expected = %d, got = %d`, 6, reply.GetC())
	}

	// Div
	args.A = proto.Int32(13)
	args.B = proto.Int32(5)
	if err = stub.Div(&args, &reply); err != nil {
		t.Fatalf(`stub.Div: %v`, err)
	}
	if reply.GetC() != 2 {
		t.Fatalf(`stub.Div: expected = %d, got = %d`, 2, reply.GetC())
	}

	// Div zero
	args.A = proto.Int32(1)
	args.B = proto.Int32(0)
	if err = stub.Div(&args, &reply); err.Error() != "divide by zero" {
		t.Fatalf(`stub.Div: expected = "%s", got = "%s"`, "divide by zero", err.Error())
	}

	// Error
	args.A = proto.Int32(1)
	args.B = proto.Int32(2)
	if err = stub.Error(&args, &reply); err.Error() != "ArithError" {
		t.Fatalf(`stub.Error: expected = "%s", got = "%s"`, "ArithError", err.Error())
	}
}
func testEchoStub(t *testing.T, stub EchoService) {
	var args EchoRequest
	var reply EchoResponse
	var err error

	// EchoService.Echo
	args.Msg = proto.String("Hello, Protobuf-RPC")
	if err = stub.Echo(&args, &reply); err != nil {
		t.Fatalf(`stub.Echo: %v`, err)
	}
	if reply.GetMsg() != args.GetMsg() {
		t.Fatalf(`stub.Echo: expected = "%s", got = "%s"`, args.GetMsg(), reply.GetMsg())
	}
}
