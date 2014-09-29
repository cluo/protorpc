// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"strings"
	"text/template"

	descriptor "github.com/chai2010/protorpc/protoc-gen-go/descriptor"
	generator "github.com/chai2010/protorpc/protoc-gen-go/generator"
)

// optimizePlugin produce the proto.Marshaler and proto.Unmarshaler interface.
type optimizePlugin struct {
	*generator.Generator
}

// Name returns the name of the plugin.
func (p *optimizePlugin) Name() string { return "optimize" }

// Init is called once after data structures are built but before
// code generation begins.
func (p *optimizePlugin) Init(g *generator.Generator) {
	p.Generator = g
}

// Generate produces the code generated by the plugin for this file.
func (p *optimizePlugin) GenerateImports(file *generator.FileDescriptor) {
	if !p.isOptimizeForSpeed(file) {
		return
	}
}

// Generate generates the Service interface.
// rpc service can't handle other proto message!!!
func (p *optimizePlugin) Generate(file *generator.FileDescriptor) {
	if !p.isOptimizeForSpeed(file) {
		return
	}
	for _, msg := range file.MessageType {
		p.genSizerInterface(file, msg)
		p.genMarshalerInterface(file, msg)
		p.genUnmarshalerInterface(file, msg)
	}
}

func (p *optimizePlugin) isOptimizeForSpeed(file *generator.FileDescriptor) bool {
	// try command line first
	// protoc --go_out=go_generic_services=true:. xxx.proto
	if value, ok := p.Generator.Param["optimize_for"]; ok {
		if strings.ToLower(value) == "speed" {
			return true
		}
	}

	// try optimize_for second
	// if the optimize_for is nil, don't use the default(SPEED) value.
	if opt := file.GetOptions(); opt != nil {
		if opt.OptimizeFor != nil && *opt.OptimizeFor == descriptor.FileOptions_SPEED {
			return true
		}
	}

	return false
}

func (p *optimizePlugin) genSizerInterface(
	file *generator.FileDescriptor,
	msg *descriptor.DescriptorProto,
) {
	const interfaceTmpl = `
func (msg *{{.MessageName}}) _Size() int {
	panic("TODO")
}
`
	out := bytes.NewBuffer([]byte{})
	t := template.Must(template.New("").Parse(interfaceTmpl))
	t.Execute(out, &struct{ MessageName string }{
		MessageName: generator.CamelCase(msg.GetName()),
	})
	p.P(out.String())
}

func (p *optimizePlugin) genMarshalerInterface(
	file *generator.FileDescriptor,
	msg *descriptor.DescriptorProto,
) {
	const interfaceTmpl = `
func (msg *{{.MessageName}}) _Marshal() ([]byte, error) {
	panic("TODO")
}
`
	out := bytes.NewBuffer([]byte{})
	t := template.Must(template.New("").Parse(interfaceTmpl))
	t.Execute(out, &struct{ MessageName string }{
		MessageName: generator.CamelCase(msg.GetName()),
	})
	p.P(out.String())
}

func (p *optimizePlugin) genUnmarshalerInterface(
	file *generator.FileDescriptor,
	msg *descriptor.DescriptorProto,
) {
	const interfaceTmpl = `
func (msg *{{.MessageName}}) _Unmarshal([]byte) error {
	panic("TODO")
}
`
	out := bytes.NewBuffer([]byte{})
	t := template.Must(template.New("").Parse(interfaceTmpl))
	t.Execute(out, &struct{ MessageName string }{
		MessageName: generator.CamelCase(msg.GetName()),
	})
	p.P(out.String())
}

func init() {
	generator.RegisterPlugin(new(optimizePlugin))
}
