package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/mfa/resources"
	"github.com/pingidentity/pingctl/internal/testutils/testutils_helpers"
)

func TestMFAPolicyExport(t *testing.T) {
	// Get initialized apiClient and resource
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	resource := resources.MFAPolicy(sdkClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_mfa_policy",
			ResourceName: "Default MFA Policy",
			ResourceID:   fmt.Sprintf("%s/6adc6dfa-d883-08ed-37c5-ea8f61029ad9", testutils_helpers.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_mfa_policy",
			ResourceName: "Test MFA Policy",
			ResourceID:   fmt.Sprintf("%s/5ae2227f-cb5b-47c3-bb40-440db09a98e6", testutils_helpers.GetEnvironmentID()),
		},
	}

	testutils_helpers.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
