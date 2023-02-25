package costformation

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

// ReadFile loads a DefinitionFile from the specified file.
func (d *DefinitionFile) ReadFile(filename string) error {
	if nil == d {
		return ErrInvalidDefinition
	}

	yfile, err := os.Open(filename)
	if err != nil {
		return err
	}

	return d.Read(yfile)
}

// Read a DefintionFile from an io.Reader
func (d *DefinitionFile) Read(input io.Reader) error {
	if nil == d {
		return ErrInvalidDefinition
	}

	if nil == input {
		return ErrInvalidReader
	}

	decoder := yaml.NewDecoder(input)
	decoder.KnownFields(true) // Error on unknown fields, should make this configurable

	var node yaml.Node
	if err := decoder.Decode(&node); err != nil {
		return err
	}

	return node.Decode(&d)
}

// WriteFile outputs a DefintionFile. If filename is "" then output to stdout, otherwise write to the file
func (d *DefinitionFile) WriteFile(filename string) error {
	if nil == d {
		return ErrInvalidDefinition
	}

	var output io.Writer

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

	return d.Write(output)
}

// Write the DefintionFile to the specified output.
func (d *DefinitionFile) Write(output io.Writer) error {
	if nil == d {
		return ErrInvalidDefinition
	}

	if nil == output {
		return ErrInvalidWriter
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
