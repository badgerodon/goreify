package main

import (
	"bytes"
	"flag"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"golang.org/x/tools/imports"

	"github.com/pkg/errors"
)

type stringList struct {
	strings *[]string
}

var _ flag.Value = (*stringList)(nil)

func (l *stringList) Set(str string) error {
	if str == "" {
		*l.strings = nil
		return nil
	}
	*l.strings = strings.Split(str, ",")
	for i := 0; i < len(*l.strings); i++ {
		(*l.strings)[i] = strings.TrimSpace((*l.strings)[i])
	}
	return nil
}
func (l *stringList) String() string {
	return strings.Join(*l.strings, ",")
}

var options = struct {
	gopath string
	out    string
	in     string
	cfg    *ReifyConfig
}{
	gopath: os.Getenv("GOPATH"),
	cfg:    &ReifyConfig{TypeSpecs: make(map[string][]string)},
}

func main() {
	log.SetFlags(0)

	flag.StringVar(&options.out, "out", options.out, "the file to write")
	// these 2 can also be sent in positionally
	flag.StringVar(&options.in, "in", options.in, "the type or function to reify")
	flag.Var(options.cfg, "types", "the types to generate")
	flag.Parse()

	// positional args
	offset := 0
	if options.in == "" {
		options.in = flag.Arg(offset)
		offset++
	}
	for flag.Arg(offset) != "" {
		options.cfg.Set(flag.Arg(offset))
		offset++
	}

	if options.gopath == "" {
		log.Fatalln("GOPATH must be set")
	}
	if options.in == "" || !strings.ContainsRune(options.in, '.') {
		log.Fatalln("an input type or function is required")
	}
	if len(options.cfg.TypeSpecs) == 0 {
		log.Fatalln("one or more output types are required")
	}

	if options.out == "" {
		_, entity := getIn()
		options.out = "reified_" + strings.ToLower(entity) + ".go"
	}

	os.Remove(options.out)

	err := generate()
	if err != nil {
		log.Fatalf("failed to generate code: %v", err)
	}
}

func getIn() (pkg, name string) {
	pos := strings.LastIndexByte(options.in, '.')
	return options.in[:pos], options.in[pos+1:]
}

func generate() error {
	pkg, entity := getIn()
	working := filepath.Join(options.gopath, "src", pkg)

	fset := new(token.FileSet)
	pkgs, err := parser.ParseDir(fset, working, nil,
		parser.AllErrors|parser.ParseComments)
	if err != nil {
		return errors.Wrapf(err, "failed to parse package")
	}

	g := NewGenerator(path.Base(pkg), fset)

	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			err = g.GenerateFromFile(file, entity, options.cfg)
			if err != nil {
				return err
			}
		}
	}

	var buf bytes.Buffer
	g.Export(&buf)
	raw := buf.Bytes()

	bs, err := imports.Process(options.out, raw, nil)
	if err != nil {
		bs = raw
	}

	return ioutil.WriteFile(options.out, bs, 0755)
}
