package component

import (
	"errors"
	"path"
	"reflect"
	"strings"

	"github.com/suger-131997/dein/internal/utils"
)

type Component struct {
	name      string
	pkgPath   string
	isPointer bool
}

func NewComponent(t reflect.Type) (Component, error) {
	isPointer := t.Kind() == reflect.Ptr
	if isPointer {
		t = t.Elem()
	}

	if t.PkgPath() == "" {
		if t.Name() == "" {
			return Component{}, errors.New("anonymous struct is not supported")
		}

		return Component{}, errors.New("builtin type is not supported")
	}

	name := t.Name()
	if name == "" {
		return Component{}, errors.New("unnamed type is not supported")
	}

	if i := strings.Index(name, "["); i >= 0 {
		params := strings.Split(name[i+1:len(name)-1], ",")
		for _, p := range params {
			if strings.Contains(p, " ") {
				return Component{}, errors.New("anonymous struct for type param is not supported")
			}
		}
	}

	return Component{
		name:      name,
		pkgPath:   t.PkgPath(),
		isPointer: isPointer,
	}, nil
}

func (c Component) Less(other Component) bool {
	if c.pkgPath != other.pkgPath {
		return c.pkgPath < other.pkgPath
	}

	if c.name != other.name {
		return c.name < other.name
	}

	if c.isPointer {
		return false
	}

	return true
}

func (c Component) Name() string {
	if i := strings.Index(c.name, "["); i >= 0 {
		return c.name[:i]
	}

	return c.name
}

func (c Component) PkgPath() string {
	return c.pkgPath
}

func (c Component) IsPointer() bool {
	return c.isPointer
}

func (c Component) TypeParams() []TypeParam {
	typeParams := make([]TypeParam, 0)

	if i := strings.Index(c.name, "["); i >= 0 {
		params := strings.Split(c.name[i+1:len(c.name)-1], ",")
		for _, p := range params {
			typeParams = append(typeParams, newTypeParams(p))
		}
	}

	return typeParams
}

func (c Component) PkgPaths() []string {
	pkgPaths := make([]string, 0)

	for _, tp := range c.TypeParams() {
		pkgPaths = append(pkgPaths, tp.pkgPaths()...)
	}

	return utils.Uniq(append(pkgPaths, c.pkgPath))
}

type TypeParam struct {
	name    string
	pkgPath string
}

func newTypeParams(typeStr string) TypeParam {
	if strings.Count(path.Base(typeStr), ".") == 0 {
		return TypeParam{
			name:    typeStr,
			pkgPath: "",
		}
	}

	var i int
	if ii := strings.Index(typeStr, "["); ii >= 0 {
		i = strings.LastIndex(typeStr[:ii], ".")
	} else {
		i = strings.LastIndex(typeStr, ".")
	}

	return TypeParam{
		name:    typeStr[i+1:],
		pkgPath: typeStr[:i],
	}
}

func (t TypeParam) Name() string {
	if i := strings.Index(t.name, "["); i >= 0 {
		return t.name[:i]
	}

	return t.name
}

func (t TypeParam) PkgPath() string {
	return t.pkgPath
}

func (t TypeParam) TypeParams() []TypeParam {
	typeParams := make([]TypeParam, 0)

	if i := strings.Index(t.name, "["); i >= 0 {
		params := strings.Split(t.name[i+1:len(t.name)-1], ",")
		for _, p := range params {
			typeParams = append(typeParams, newTypeParams(p))
		}
	}

	return typeParams
}

func (t TypeParam) pkgPaths() []string {
	pkgPaths := make([]string, 0)

	if t.pkgPath != "" {
		pkgPaths = append(pkgPaths, t.pkgPath)
	}

	for _, tp := range t.TypeParams() {
		pkgPaths = append(pkgPaths, tp.pkgPaths()...)
	}

	return utils.Uniq(pkgPaths)
}
