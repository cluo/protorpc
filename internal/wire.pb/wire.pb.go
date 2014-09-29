// Code generated by protoc-gen-go.
// source: wire.proto
// DO NOT EDIT!

/*
Package google_protobuf_rpc_wire is a generated protocol buffer package.

	protorpc wire format wrapper

	0. Frame Format
	len : uvarint64
	data: byte[len]

	1. Client Send Request
	Send RequestHeader: sendFrame(conn, hdr, len(hdr))
	Send Request: sendFrame(conn, body, hdr.snappy_compressed_request_len)

	2. Server Recv Request
	Recv RequestHeader: recvFrame(conn, hdr, max_hdr_len, 0)
	Recv Request: recvFrame(conn, body, hdr.snappy_compressed_request_len, 0)

	3. Server Send Response
	Send ResponseHeader: sendFrame(conn, hdr, len(hdr))
	Send Response: sendFrame(conn, body, hdr.snappy_compressed_response_len)

	4. Client Recv Response
	Recv ResponseHeader: recvFrame(conn, hdr, max_hdr_len, 0)
	Recv Response: recvFrame(conn, body, hdr.snappy_compressed_response_len, 0)

	5. Header Size
	len(RequestHeader)  < Const.max_header_len.default
	len(ResponseHeader) < Const.max_header_len.default

It is generated from these files:
	wire.proto

It has these top-level messages:
	Const
	RequestHeader
	ResponseHeader
*/
package google_protobuf_rpc_wire

import proto "github.com/chai2010/protorpc/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Const struct {
	MaxHeaderLen     *uint32 `protobuf:"varint,1,opt,name=max_header_len,def=1024" json:"max_header_len,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Const) Reset()         { *m = Const{} }
func (m *Const) String() string { return proto.CompactTextString(m) }
func (*Const) ProtoMessage()    {}

const Default_Const_MaxHeaderLen uint32 = 1024

func (m *Const) GetMaxHeaderLen() uint32 {
	if m != nil && m.MaxHeaderLen != nil {
		return *m.MaxHeaderLen
	}
	return Default_Const_MaxHeaderLen
}

type RequestHeader struct {
	Id                         *uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Method                     *string `protobuf:"bytes,2,opt,name=method" json:"method,omitempty"`
	RawRequestLen              *uint32 `protobuf:"varint,3,opt,name=raw_request_len" json:"raw_request_len,omitempty"`
	SnappyCompressedRequestLen *uint32 `protobuf:"varint,4,opt,name=snappy_compressed_request_len" json:"snappy_compressed_request_len,omitempty"`
	Checksum                   *uint32 `protobuf:"varint,5,opt,name=checksum" json:"checksum,omitempty"`
	XXX_unrecognized           []byte  `json:"-"`
}

func (m *RequestHeader) Reset()         { *m = RequestHeader{} }
func (m *RequestHeader) String() string { return proto.CompactTextString(m) }
func (*RequestHeader) ProtoMessage()    {}

func (m *RequestHeader) GetId() uint64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *RequestHeader) GetMethod() string {
	if m != nil && m.Method != nil {
		return *m.Method
	}
	return ""
}

func (m *RequestHeader) GetRawRequestLen() uint32 {
	if m != nil && m.RawRequestLen != nil {
		return *m.RawRequestLen
	}
	return 0
}

func (m *RequestHeader) GetSnappyCompressedRequestLen() uint32 {
	if m != nil && m.SnappyCompressedRequestLen != nil {
		return *m.SnappyCompressedRequestLen
	}
	return 0
}

func (m *RequestHeader) GetChecksum() uint32 {
	if m != nil && m.Checksum != nil {
		return *m.Checksum
	}
	return 0
}

type ResponseHeader struct {
	Id                          *uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Error                       *string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	RawResponseLen              *uint32 `protobuf:"varint,3,opt,name=raw_response_len" json:"raw_response_len,omitempty"`
	SnappyCompressedResponseLen *uint32 `protobuf:"varint,4,opt,name=snappy_compressed_response_len" json:"snappy_compressed_response_len,omitempty"`
	Checksum                    *uint32 `protobuf:"varint,5,opt,name=checksum" json:"checksum,omitempty"`
	XXX_unrecognized            []byte  `json:"-"`
}

func (m *ResponseHeader) Reset()         { *m = ResponseHeader{} }
func (m *ResponseHeader) String() string { return proto.CompactTextString(m) }
func (*ResponseHeader) ProtoMessage()    {}

func (m *ResponseHeader) GetId() uint64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *ResponseHeader) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

func (m *ResponseHeader) GetRawResponseLen() uint32 {
	if m != nil && m.RawResponseLen != nil {
		return *m.RawResponseLen
	}
	return 0
}

func (m *ResponseHeader) GetSnappyCompressedResponseLen() uint32 {
	if m != nil && m.SnappyCompressedResponseLen != nil {
		return *m.SnappyCompressedResponseLen
	}
	return 0
}

func (m *ResponseHeader) GetChecksum() uint32 {
	if m != nil && m.Checksum != nil {
		return *m.Checksum
	}
	return 0
}

func init() {
}
