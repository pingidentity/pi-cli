package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingctl/internal/testutils/testutils_helpers"
)

func TestPasswordPolicyExport(t *testing.T) {
	// Get initialized apiClient and resource
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	resource := resources.PasswordPolicy(sdkClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_password_policy",
			ResourceName: "Standard",
			ResourceID:   fmt.Sprintf("%s/10c1f1bc-3dff-49ca-9abb-cf034b728793", testutils_helpers.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_password_policy",
			ResourceName: "Basic",
			ResourceID:   fmt.Sprintf("%s/48641620-f51d-4675-86e1-e45d378ac0b2", testutils_helpers.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_password_policy",
			ResourceName: "Passphrase",
			ResourceID:   fmt.Sprintf("%s/686e2710-d59f-484a-8ba5-47959753012c", testutils_helpers.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_password_policy",
			ResourceName: "LDAP Gateway Policy",
			ResourceID:   fmt.Sprintf("%s/c79032d2-b156-46a5-a9c9-7d18e93095b7", testutils_helpers.GetEnvironmentID()),
		},
	}

	testutils_helpers.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
