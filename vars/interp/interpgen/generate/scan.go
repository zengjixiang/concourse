package generate

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/hashicorp/go-multierror"
)

const generateCommentPrefix = "//interpgen:generate"

type TypeInfo struct {
	Name   string
	Export bool
}

type namedType struct {
}

func Scan(files []*ast.File) ([]TypeInfo, error) {
	namedTypes := detectNamedTypes(files)
	typeInfos, err := detectGenerateTypes(files, namedTypes)
	if err != nil {
		return nil, err
	}
	return typeInfos, nil
}

func forEachComment(files []*ast.File, cb func(comment *ast.Comment) error) error {
	var merr error
	for _, f := range files {
		for _, commentGroup := range f.Comments {
			for _, comment := range commentGroup.List {
				if err := cb(comment); err != nil {
					merr = multierror.Append(merr, err)
				}
			}
		}
	}
	return merr
}

func detectNamedTypes(files []*ast.File) map[string]namedType {
	namedTypes := map[string]namedType{}
	for _, f := range files {
		ast.Inspect(f, func(n ast.Node) bool {
			switch v := n.(type) {
			case *ast.TypeSpec:
				if _, exists := namedTypes[v.Name.Name]; !exists {
					namedTypes[v.Name.Name] = namedType{}
				}
			}
			return true
		})
	}
	return namedTypes
}

func detectGenerateTypes(files []*ast.File, namedTypes map[string]namedType) ([]TypeInfo, error) {
	typeInfos := map[string]TypeInfo{}
	err := forEachComment(files, func(comment *ast.Comment) error {
		if !strings.HasPrefix(comment.Text, generateCommentPrefix) {
			return nil
		}
		// e.g. `//interpgen:generate MyType (export)`
		fields := strings.Fields(comment.Text[len(generateCommentPrefix):])
		var exists bool
		var typeInfo TypeInfo
		switch len(fields) {
		case 1:
			name := fields[0]
			_, exists = typeInfos[name]
			typeInfo = TypeInfo{Name: name}
		case 2:
			name, option := fields[0], fields[1]
			if option != "export" {
				return fmt.Errorf("invalid option %q - expected \"export\" (or nothing)", fields[1])
			}
			typeInfo = TypeInfo{Name: name, Export: true}
		default:
			return fmt.Errorf("expected usage: %s MyType (export)", generateCommentPrefix)
		}
		if exists {
			return fmt.Errorf("duplicate type %q", typeInfo.Name)
		}
		if _, ok := namedTypes[typeInfo.Name]; !ok {
			return fmt.Errorf("interpgen can only generated named types, not %q", typeInfo.Name)
		}
		typeInfos[typeInfo.Name] = typeInfo
		return nil
	})
	if err != nil {
		return nil, err
	}

	typeInfosL := make([]TypeInfo, 0, len(typeInfos))
	for _, ti := range typeInfos {
		typeInfosL = append(typeInfosL, ti)
	}
	return typeInfosL, nil
}
