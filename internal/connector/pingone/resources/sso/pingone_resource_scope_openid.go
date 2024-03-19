package sso

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/resources/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneResourceScopeOpenIdResource{}
)

type PingoneResourceScopeOpenIdResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneResourceScopeOpenIdResource
func ResourceScopeOpenId(clientInfo *connector.SDKClientInfo) *PingoneResourceScopeOpenIdResource {
	return &PingoneResourceScopeOpenIdResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneResourceScopeOpenIdResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.ResourcesApi.ReadAllResources(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllResources"

	embedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, resource := range embedded.GetResources() {
		resourceId, resourceIdOk := resource.GetIdOk()
		resourceName, resourceNameOk := resource.GetNameOk()

		if resourceIdOk && resourceNameOk {
			apiResourceScopeOpenIdsExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.ResourceScopesApi.ReadAllResourceScopes(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *resourceId).Execute
			apiResourceScopeOpenIdsFunctionName := "ReadAllResourceScopes"

			embeddedResourceScopeOpenIds, err := common.GetManagementEmbedded(apiResourceScopeOpenIdsExecuteFunc, apiResourceScopeOpenIdsFunctionName, r.ResourceType())
			if err != nil {
				return nil, err
			}

			for _, scopeOpenId := range embeddedResourceScopeOpenIds.GetScopes() {
				scopeOpenIdId, scopeOpenIdIdOk := scopeOpenId.GetIdOk()
				scopeOpenIdName, scopeOpenIdNameOk := scopeOpenId.GetNameOk()
				_, mappedClaimsOk := scopeOpenId.GetMappedClaimsOk()
				if scopeOpenIdIdOk && scopeOpenIdNameOk && mappedClaimsOk {
					importBlocks = append(importBlocks, connector.ImportBlock{
						ResourceType: r.ResourceType(),
						ResourceName: fmt.Sprintf("%s_%s", *resourceName, *scopeOpenIdName),
						ResourceID:   fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *scopeOpenIdId),
					})
				}
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingoneResourceScopeOpenIdResource) ResourceType() string {
	return "pingone_resource_scope_openid"
}
