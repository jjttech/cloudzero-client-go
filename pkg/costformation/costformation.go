package costformation

import (
	"context"
	"encoding/json"
	"io"

	"github.com/jjttech/cloudzero-client-go/internal/http"
	"github.com/jjttech/cloudzero-client-go/pkg/config"
)

const (
	DefaultBasePath = "/v2/costformation"
	DefinitionPath  = DefaultBasePath + "/definition"
)

// CostFormation API Client
type CostFormation struct {
	client  *http.Client
	baseURL string
}

// New returns a new instance of the CostFormation API client
func New(cfg config.Config) (*CostFormation, error) {
	client, err := http.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return &CostFormation{
		client:  client,
		baseURL: cfg.BaseURL,
	}, nil
}

// DefinitionList returns a list of definition files
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
func (c *CostFormation) DefinitionFetch(ctx context.Context, version string) (*DefinitionFile, error) {
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

	// TODO: We have to fetch the actual file from S3 as well...

	ret := DefinitionFile{
		LastUpdated:   data.Version.LastUpdated,
		LastUpdatedBy: data.Version.LastUpdatedBy,
		Version:       data.Version.Version,
	}

	return &ret, nil
}
