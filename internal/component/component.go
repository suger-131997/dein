package component

import (
	"errors"
	"github.com/suger-131997/dein/internal/utils"
	"reflect"
	"strings"
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
		return Component{}, errors.New("builtin type is not supported")
	}

	return Component{
		name:      t.Name(),
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

func (c Component) PkgPaths() []string {
	pkgPaths := make([]string, 0)

	if i := strings.Index(c.name, "["); i >= 0 {
		params := strings.Split(c.name[i+1:len(c.name)-1], ",")
		for _, p := range params {
			pkgPaths = append(pkgPaths, p[:strings.LastIndex(p, ".")])
		}
	}

	return utils.Uniq(append(pkgPaths, c.pkgPath))
}
