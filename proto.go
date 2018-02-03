package gendoc

import (
	"fmt"

	"github.com/pseudomuto/protoc-gen-doc/parser"
)

type ProtoPackages struct {
	pkgs map[string]Package
}

type Package struct {
	Name           string
	Enums          map[string]*Enum
	FileExtensions map[string]*FileExtension
	Messages       map[string]*Message
	Services       map[string]*Service
}

func newProtoPackages() *ProtoPackages {
	return &ProtoPackages{
		pkgs: make(map[string]Package),
	}
}

func newPackage(pkgName string) Package {
	return Package{
		Name:           pkgName,
		Enums:          make(map[string]*Enum),
		FileExtensions: make(map[string]*FileExtension),
		Messages:       make(map[string]*Message),
		Services:       make(map[string]*Service),
	}
}

func (pp *ProtoPackages) parseEnum(pe *parser.Enum) *ProtoPackages {
	pkg, ok := pp.pkgs[pe.Package]
	if !ok {
		pkg = newPackage(pe.Package)
		pp.pkgs[pe.Package] = pkg
	}
	if _, ok := pkg.Enums[pe.FullName()]; ok {
		panic(fmt.Sprintf("%s in package %s defined twice.", pe.FullName(), pe.Package))
	}
	pkg.Enums[pe.FullName()] = parseEnum(pe)
	return pp
}

func (pp *ProtoPackages) parseExtension(pe *parser.Extension) *ProtoPackages {
	pkg, ok := pp.pkgs[pe.Package]
	if !ok {
		pkg = newPackage(pe.Package)
		pp.pkgs[pe.Package] = pkg
	}
	if _, ok := pkg.FileExtensions[pe.FullName()]; ok {
		panic(fmt.Sprintf("%s in package %s defined twice.", pe.FullName(), pe.Package))
	}
	pkg.FileExtensions[pe.FullName()] = parseFileExtension(pe)
	return pp
}

func (pp *ProtoPackages) parseMessage(pe *parser.Message) *ProtoPackages {
	pkg, ok := pp.pkgs[pe.Package]
	if !ok {
		pkg = newPackage(pe.Package)
		pp.pkgs[pe.Package] = pkg
	}
	if _, ok := pkg.Messages[pe.FullName()]; ok {
		panic(fmt.Sprintf("%s in package %s defined twice.", pe.FullName(), pe.Package))
	}
	pkg.Messages[pe.FullName()] = parseMessage(pe)
	return pp
}

func (pp *ProtoPackages) parseService(pe *parser.Service) *ProtoPackages {
	pkg, ok := pp.pkgs[pe.Package]
	if !ok {
		pkg = newPackage(pe.Package)
		pp.pkgs[pe.Package] = pkg
	}
	if _, ok := pkg.Services[pe.FullName()]; ok {
		panic(fmt.Sprintf("%s in package %s defined twice.", pe.FullName(), pe.Package))
	}
	pkg.Services[pe.FullName()] = parseService(pe)
	return pp
}
