package config

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	hcl2 "github.com/hashicorp/hcl/v2"
	hcl2syntax "github.com/hashicorp/hcl/v2/hclsyntax"

)

//ParseFile gets the io.Reader and the extension of the file
func ParseFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
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
func Parse(r io.Reader, filename, extension string) error {
	switch extension {
	case "hcl":
		return parseHCL(r, filename)

	case "json":
		return parseJSON(r, filename)

	default:
		return fmt.Errorf("extension must be either 'hcl' or 'json'")

	}
}

func parseHCL(r io.Reader, filename string) error {
	src, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	_, diag := hcl2syntax.ParseConfig(src, filename, hcl2.Pos{})

	if diag.HasErrors() {
		return diag
	}

	return nil
}

func parseJSON(r io.Reader, filename string) error {
	return nil
}
