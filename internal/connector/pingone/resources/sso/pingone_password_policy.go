package sso

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/resources/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingonePasswordPolicyResource{}
)

type PingonePasswordPolicyResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingonePasswordPolicyResource
func PasswordPolicy(clientInfo *connector.SDKClientInfo) *PingonePasswordPolicyResource {
	return &PingonePasswordPolicyResource{
		clientInfo: clientInfo,
	}
}

func (r *PingonePasswordPolicyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.PasswordPoliciesApi.ReadAllPasswordPolicies(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllPasswordPolicies"

	embedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, passwordPolicy := range embedded.GetPasswordPolicies() {
		passwordPolicyId, passwordPolicyIdOk := passwordPolicy.GetIdOk()
		passwordPolicyName, passwordPolicyNameOk := passwordPolicy.GetNameOk()

		if passwordPolicyIdOk && passwordPolicyNameOk {
			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType: r.ResourceType(),
				ResourceName: *passwordPolicyName,
				ResourceID:   fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *passwordPolicyId),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingonePasswordPolicyResource) ResourceType() string {
	return "pingone_password_policy"
}
