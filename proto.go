package gendoc

import (
	"fmt"

	"github.com/Kunde21/protoc-gen-doc/parser"
)

type ProtoPackages struct {
	Packages map[string]Package
	Enums    map[string]*Enum
	Messages map[string]*Message
	Services map[string]*Service
}

type Package struct {
	Name     string
	Enums    map[string]*Enum
	Messages map[string]*Message
	Services map[string]*Service
}

func newProtoPackages() ProtoPackages {
	return ProtoPackages{
		Packages: make(map[string]Package),
		Enums:    make(map[string]*Enum),
		Messages: make(map[string]*Message),
		Services: make(map[string]*Service),
	}
}

func newPackage(pkgName string) Package {
	return Package{
		Name:     pkgName,
		Enums:    make(map[string]*Enum),
		Messages: make(map[string]*Message),
		Services: make(map[string]*Service),
	}
}

func (pp ProtoPackages) parseEnum(pe *parser.Enum) ProtoPackages {
	pkg, ok := pp.Packages[pe.Package]
	if !ok {
		pkg = newPackage(pe.Package)
		pp.Packages[pe.Package] = pkg
	}
	if _, ok := pp.Enums[pe.FullName()]; ok {
		panic(fmt.Sprintf("%s in package %s defined twice.", pe.FullName(), pe.Package))
	}
	enum := parseEnum(pe)
	pkg.Enums[pe.FullName()] = enum
	pp.Enums[pe.FullName()] = enum
	return pp
}

func (pp ProtoPackages) parseMessage(pe *parser.Message) ProtoPackages {
	pkg, ok := pp.Packages[pe.Package]
	if !ok {
		pkg = newPackage(pe.Package)
		pp.Packages[pe.Package] = pkg
	}
	if _, ok := pp.Messages[pe.FullName()]; ok {
		panic(fmt.Sprintf("%s in package %s defined twice.", pe.FullName(), pe.Package))
	}
	msg := parseMessage(pe)
	pkg.Messages[pe.FullName()] = msg
	pp.Messages[pe.FullName()] = msg
	return pp
}

func (pp ProtoPackages) parseService(pe *parser.Service) ProtoPackages {
	pkg, ok := pp.Packages[pe.Package]
	if !ok {
		pkg = newPackage(pe.Package)
		pp.Packages[pe.Package] = pkg
	}
	if _, ok := pp.Services[pe.FullName()]; ok {
		panic(fmt.Sprintf("%s in package %s defined twice.", pe.FullName(), pe.Package))
	}
	svc := parseService(pe)
	pkg.Services[pe.FullName()] = svc
	pp.Services[pe.FullName()] = svc
	return pp
}
