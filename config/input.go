package config

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	hcl2 "github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	hcl2syntax "github.com/hashicorp/hcl/v2/hclsyntax"
	hcl2json "github.com/hashicorp/hcl/v2/json"
)

var (
	ErrorExtension = errors.New("extension must be either 'hcl' or 'json'")
)

//ParseFile gets the io.Reader and the extension of the file
func ParseFile(filename string) (*Root, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	// check extension hcl or json
	extension := filepath.Ext(filename)
	if len(extension) > 0 {
		extension = extension[1:]
	}

	return Parse(f, filename, extension)
}

//Parse parses the configuration from hcl or json
func Parse(r io.Reader, filename, extension string) (*Root, error) {
	switch extension {
	case "hcl":
		return parseHCL(r, filename)

	case "json":
		return parseJSON(r, filename)

	default:
		return nil, ErrorExtension

	}
}

func parseHCL(r io.Reader, filename string) (*Root, error) {
	src, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	f, diag := hcl2syntax.ParseConfig(src, filename, hcl2.Pos{})

	if diag.HasErrors() {
		return nil, diag
	}

	var config Root
	diag = gohcl.DecodeBody(f.Body, nil, &config)
	if diag.HasErrors() {
		return nil, diag
	}

	return &config, nil
}

func parseJSON(r io.Reader, filename string) (*Root, error) {
	src, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	f, diag := hcl2json.Parse(src, filename)

	if diag.HasErrors() {
		return nil, diag
	}

	var config Root
	diag = gohcl.DecodeBody(f.Body, nil, &config)

	if diag.HasErrors() {
		return nil, diag
	}

	return &config, nil
}
