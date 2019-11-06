package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

//WriteFile saves the tifle
func WriteFile(filename, format string, root *Root) error {
	extension := filepath.Ext(filename)
	if len(extension) > 0 {
		extension = extension[1:]
		filename = strings.TrimSuffix(filename, extension)
	}

	filename = fmt.Sprintf("%sconsul", filename)

	//  if format is not json/hcl use the same of the extension
	if format != "json" && format != "hcl" {
		format = extension
	}

	filename = fmt.Sprintf("%s.%s", filename, format)

	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer f.Close()

	return Write(f, format, root)
}

//Write info to do
func Write(w io.Writer, extension string, root *Root) error {
	switch extension {
	case "hcl":
		return WriteHCL(w, root)
	case "json":
		return WriteJSON(w, root)
	default:
		return ErrorExtension
	}
}

//WriteHCL the config in that language
func WriteHCL(w io.Writer, root *Root) error {
	f := hclwrite.NewEmptyFile()
	gohcl.EncodeIntoBody(root, f.Body())
	_, err := f.WriteTo(w)
	return err
}

//WriteJSON config in json
func WriteJSON(w io.Writer, root *Root) error {
	payload, err := json.MarshalIndent(root, "", "	")
	if err != nil {
		return err
	}
	_, err = w.Write(payload)
	return err
}
