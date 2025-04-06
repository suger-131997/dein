package component

import (
	"errors"
	"path"
	"reflect"
	"regexp"
	"strings"

	"github.com/suger-131997/dein/internal/utils"
)

type Component struct {
	name    string
	pkgPath string
	prefix  string
}

var mapRegex = regexp.MustCompile(`[]*,]*map\[`)

func NewComponent(t reflect.Type) (Component, error) {
	prefix := ""

	for t.Kind() == reflect.Slice || t.Kind() == reflect.Array || t.Kind() == reflect.Ptr {
		if t.Kind() == reflect.Ptr {
			prefix += "*"
		}

		if t.Kind() == reflect.Slice || t.Kind() == reflect.Array {
			prefix += "[]"
		}

		t = t.Elem()
	}

	if t.Kind() == reflect.Map {
		return Component{}, errors.New("map type is not supported")
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
		typeParams := name[i+1 : len(name)-1]
		if strings.Contains(typeParams, " ") {
			return Component{}, errors.New("anonymous struct for type param is not supported")
		}

		if mapRegex.MatchString(typeParams) {
			return Component{}, errors.New("map type is not supported in type param")
		}
	}

	c := Component{
		name:    name,
		pkgPath: t.PkgPath(),
		prefix:  prefix,
	}

	return c, nil
}

func (c Component) Less(other Component) bool {
	if c.pkgPath != other.pkgPath {
		return c.pkgPath < other.pkgPath
	}

	if c.name != other.name {
		return c.name < other.name
	}

	if c.prefix != other.prefix {
		return c.prefix < other.prefix
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

func (c Component) Prefix() string {
	return c.prefix
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
	prefix  string
}

func newTypeParams(rawTypeStr string) TypeParam {
	pi := 0

	for rawTypeStr[pi] == '[' || rawTypeStr[pi] == ']' || rawTypeStr[pi] == '*' {
		pi++
	}

	prefix := rawTypeStr[:pi]
	typeStr := rawTypeStr[pi:]

	if strings.Count(path.Base(typeStr), ".") == 0 {
		return TypeParam{
			name:    typeStr,
			pkgPath: "",
			prefix:  prefix,
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
		prefix:  prefix,
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

func (t TypeParam) Prefix() string {
	return t.prefix
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
