package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/concourse/concourse/vars/interp/interpgen/generate"
	"github.com/hashicorp/go-multierror"
	"golang.org/x/tools/go/packages"
)

func main() {
	if len(os.Args) != 3 {
		fail("expected usage: %s ./path_to_directory ./path_to_destination.go", os.Args[0])
	}

	dir, output := os.Args[1], os.Args[2]
	cfg := &packages.Config{Mode: packages.NeedSyntax | packages.NeedName}
	pkgs, err := packages.Load(cfg, dir)
	if err != nil {
		fail("failed to load directory: %v", err)
	}
	if len(pkgs) == 0 {
		fail("no packages found")
	}
	pkg := pkgs[0]
	var merr *multierror.Error
	for _, err := range pkg.Errors {
		merr = multierror.Append(merr, err)
	}
	if merr.ErrorOrNil() != nil {
		fail("%v", merr)
	}

	types, err := generate.Scan(pkg.Syntax)
	if err != nil {
		fail("failed to scan files: %v", err)
	}

	content, err := generate.Render(pkg.Name, types)
	if err != nil {
		fail("failed to generate types: %v", err)
	}

	if err := ioutil.WriteFile(output, content, 0755); err != nil {
		fail("failed to save file: %v", err)
	}
}

func fail(format string, args ...interface{}) {
	panic(fmt.Sprintf(format, args...))
}
