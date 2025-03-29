//go:generate go run . -output=../provider_gen.go

package main

import (
	"bytes"
	"flag"
	"go/format"
	"log"
	"os"
	"text/template"
)

var filename = flag.String("output", "dein_gen.go", "output file name")

const num = 10

func main() {
	flag.Parse()

	var buf bytes.Buffer

	err := template.Must(template.New("").Parse(tmpl)).Execute(&buf, num)
	if err != nil {
		log.Fatal(err)
	}

	data, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(*filename, data, 0o644)
	if err != nil {
		log.Fatal(err)
	}
}

const tmpl = `// Code generated by dein/gen.go. DO NOT EDIT.
package dein

import (
	"github.com/suger-131997/dein/internal/provider"
)

{{range $i := .}}
// P{{$i}} creates a new provider that initializes with a constructor that takes {{if eq $i 0}}no{{else}}{{$i}}{{end}}  arguments and returns a value of type R.
func P{{$i}}[{{range $ii := $i}}T{{$ii}} ,{{end}}R any](f func({{range $ii := $i}}T{{$ii}} ,{{end}}) R) *provider.Provider  {
	return provider.NewConstructorProvider(f, false)
}

// PE{{$i}} creates a new provider that initializes with a constructor that takes {{if eq $i 0}}no{{else}}{{$i}}{{end}}  arguments and returns a value of type R or an error.
func PE{{$i}}[{{range $ii := $i}}T{{$ii}} ,{{end}}R any](f func({{range $ii := $i}}T{{$ii}} ,{{end}}) (R, error)) *provider.Provider  {
	return provider.NewConstructorProvider(f, true)
}

// PF{{$i}} creates a new provider that initializes with a function provided at build time, which takes {{if eq $i 0}}no{{else}}{{$i}}{{end}} arguments and returns a value of type R.
func PF{{$i}}[{{range $ii := $i}}T{{$ii}} ,{{end}}R any]() *provider.Provider  {
	return provider.NewFunctionProvider(rt[R](), false, {{range $ii := $i}}rt[T{{$ii}}](),{{end}})
}

// PFE{{$i}} creates a new provider that initializes with a function provided at build time, which takes {{if eq $i 0}}no{{else}}{{$i}}{{end}} arguments and returns a value of type R or an error.
func PFE{{$i}}[{{range $ii := $i}}T{{$ii}} ,{{end}}R any]() *provider.Provider  {
	return provider.NewFunctionProvider(rt[R](), true, {{range $ii := $i}}rt[T{{$ii}}](),{{end}})
}
{{end}}
`
