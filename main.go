package main

import (
	"flag"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

var options struct {
	working string
	pkg     string
	out     string
	gopath  string
}

func main() {
	log.SetFlags(0)

	options.gopath = os.Getenv("GOPATH")
	flag.StringVar(&options.pkg, "package", "",
		"the package to reify (defaults to the current directory)")
	flag.StringVar(&options.out, "out", "reified.gen.go", "the file to write")
	flag.Parse()

	if options.gopath == "" {
		log.Fatalln("GOPATH must be set")
	}

	// use the current working directory to determine the package
	if options.pkg == "" {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatalf("failed to get current working directory: %v", err)
		}
		options.working = dir
		if !strings.HasPrefix(options.working, options.gopath) {
			log.Fatalf("directory must be in GOPATH")
		}
		options.pkg = strings.Replace(options.working[len(options.gopath):],
			string(filepath.Separator), "/", -1)
		if strings.HasPrefix(options.pkg, "/") {
			options.pkg = options.pkg[1:]
		}
	} else {
		options.working = filepath.Join(options.gopath, "src", options.pkg)
	}

	os.Remove(filepath.Join(options.working, options.out))

	err := generate()
	if err != nil {
		log.Fatalf("failed to generate code: %v", err)
	}
}

func generate() error {
	fset := new(token.FileSet)
	pkgs, err := parser.ParseDir(fset, options.working, nil,
		parser.AllErrors|parser.ParseComments)
	if err != nil {
		return errors.Wrapf(err, "failed to parse package")
	}

	g := NewGenerator(path.Base(options.pkg), fset)

	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			err = g.GenerateFromFile(file)
			if err != nil {
				return err
			}
		}
	}

	f, err := os.Create(filepath.Join(options.working, options.out))
	if err != nil {
		return err
	}
	defer f.Close()

	return g.Export(f)
}
