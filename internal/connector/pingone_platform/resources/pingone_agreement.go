package resources

import (
	"context"
	"fmt"

	sdk "github.com/patrickcping/pingone-go-sdk-v2/pingone"
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneAgreementResource{}
)

type PingoneAgreementResource struct {
	context       context.Context
	apiClient     *sdk.Client
	environmentID string
}

// Utility method for creating a PingoneAgreementResource
func AgreementResource(ctx context.Context, apiClient *sdk.Client, environmentID string) *PingoneAgreementResource {
	return &PingoneAgreementResource{
		context:       ctx,
		apiClient:     apiClient,
		environmentID: environmentID,
	}
}

func (r *PingoneAgreementResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all pingone_agreement resources...")

	entityArray, response, err := r.apiClient.ManagementAPIClient.AgreementsResourcesApi.ReadAllAgreements(r.context, r.environmentID).Execute()
	defer response.Body.Close()
	if err != nil {
		l.Error().Err(err).Msgf("ReadAllAgreements Response Code: %s\nResponse Body: %s", response.Status, response.Body)
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	if entityArray != nil && entityArray.Embedded != nil && entityArray.Embedded.Agreements != nil {
		l.Debug().Msgf("Generating Import Blocks for all pingone_agreement resources...")
		for _, agreement := range entityArray.Embedded.Agreements {
			if agreement.Id != nil && agreement.Name != "" && agreement.Environment != nil && agreement.Environment.Id != nil {
				importBlocks = append(importBlocks, connector.ImportBlock{
					ResourceType: r.ResourceType(),
					ResourceName: agreement.Name,
					ResourceID:   fmt.Sprintf("%s/%s", *agreement.Environment.Id, *agreement.Id),
				})
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingoneAgreementResource) ResourceType() string {
	return "pingone_agreement"
}
