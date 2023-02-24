package costformation

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

// Read a DefintionFile from disk
func (d *DefinitionFile) Read(filename string) error {
	yfile, err := os.Open(filename)
	if err != nil {
		return err
	}

	decoder := yaml.NewDecoder(yfile)
	decoder.KnownFields(true) // Error on unknown fields, should make this configurable

	var node yaml.Node
	if err = decoder.Decode(&node); err != nil {
		return err
	}

	return node.Decode(&d)
}

// Write a DefintionFile out. If filename is "" then output to stdout, otherwise write to the file
func (d *DefinitionFile) Write(filename string) error {
	var output io.Writer

	if nil == d {
		return ErrInvalidDefinition
	}

	if "" != filename {
		fh, err := os.Create(filename)
		if err != nil {
			return ErrUnableToWrite
		}
		defer fh.Close()

		w := bufio.NewWriter(fh)
		defer w.Flush()

		output = w
	} else {
		output = os.Stdout
	}

	if d.HeadComment != "" {
		if _, err := fmt.Fprint(output, d.HeadComment, "\n\n"); err != nil {
			return err
		}
	}

	enc := yaml.NewEncoder(output)
	enc.SetIndent(2) // TODO: Should be configurable
	if err := enc.Encode(d); err != nil {
		return err
	}

	return enc.Close()
}
