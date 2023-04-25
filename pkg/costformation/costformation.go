package costformation

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	czhttp "github.com/jjttech/cloudzero-client-go/internal/http"
	"github.com/jjttech/cloudzero-client-go/pkg/config"
)

const (
	DefaultBasePath = "/v2/costformation"
	DefinitionPath  = DefaultBasePath + "/definition"
)

// CostFormation API Client
type CostFormation struct {
	client  *czhttp.Client
	baseURL string
}

// New returns a new instance of the CostFormation API client
func New(cfg config.Config) (*CostFormation, error) {
	client, err := czhttp.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return &CostFormation{
		client:  client,
		baseURL: cfg.BaseURL,
	}, nil
}

// ReadFile is a wrapper for reading a yaml definition file
func (c *CostFormation) ReadFile(filename string) (*Definition, error) {
	ret := Definition{}

	if err := ret.ReadFile(filename); err != nil {
		return nil, err
	}

	return &ret, nil
}

// Read is a wrapper for reading a definition from an io.Reader
func (c *CostFormation) Read(input io.Reader) (*Definition, error) {
	ret := Definition{}

	if err := ret.Read(input); err != nil {
		return nil, err
	}

	return &ret, nil
}

// WriteFile is a wrapper for outputing the definition file content to a file (or stdout
// if filename is ""
func (c *CostFormation) WriteFile(d *Definition, filename string) error {
	if nil == d {
		return ErrInvalidDefinition
	}

	return d.WriteFile(filename)
}

// Write is a wrapper for writing out the definition to an io.Writer
func (c *CostFormation) Write(d *Definition, output io.Writer) error {
	if nil == d {
		return ErrInvalidDefinition
	}

	return d.Write(output)
}

// DefinitionVersions returns a list of definition files
func (c *CostFormation) DefinitionVersions(ctx context.Context) ([]DefinitionVersion, error) {
	resp, err := c.client.Get(ctx, c.baseURL+DefinitionPath)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return nil, readErr
	}

	data := defRespListVersions{}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data.Versions, nil
}

// DefintionFetch returns a specific version file
func (c *CostFormation) DefinitionFetch(ctx context.Context, version string) (*Definition, error) {
	if "" == version {
		version = "latest"
	}

	resp, err := c.client.Get(ctx, c.baseURL+DefinitionPath+"/"+version)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return nil, readErr
	}

	data := defRespGetVersion{}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	// We get a URI back, so make another fetch and pass it to the reader
	if "" == data.Version.URI {
		return nil, fmt.Errorf("failed to fetch definition version: %s", version)
	}

	// Use raw HTTP client here, as the URI contains all info needed for the fetch (from S3)
	dResp, dErr := http.Get(data.Version.URI)
	if dErr != nil {
		return nil, err
	}
	defer dResp.Body.Close()

	ret := Definition{}

	if err = ret.Read(dResp.Body); err != nil {
		return nil, err
	}

	ret.LastUpdated = data.Version.LastUpdated
	ret.LastUpdatedBy = data.Version.LastUpdatedBy
	ret.Version = data.Version.Version

	return &ret, nil
}
