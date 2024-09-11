package testutils_viper

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/pingidentity/pingctl/internal/configuration/options"
	"github.com/pingidentity/pingctl/internal/profiles"
)

const (
	outputDirectoryReplacement = "[REPLACE_WITH_OUTPUT_DIRECTORY]"
)

var (
	configFileContents        string
	defaultConfigFileContents string = fmt.Sprintf(`activeProfile: default
default:
    description: "default description"
    pingctl:
        color: true
        outputFormat: text
    export:
        outputDirectory: %s
        pingone:
            region: %s
            worker:
                clientid: %s
                clientsecret: %s
                environmentid: %s
        pingfederate:
            adminapipath: "%s"
            clientcredentialsauth:
                clientid: "%s"
                clientsecret: "%s"
                scopes: "%s"
                tokenurl: "%s"
            httpshost: "%s"
            insecureTrustAllTLS: true
            xBypassExternalValidationHeader: true
production:
    description: "test profile description"
    pingctl:
        color: true
        outputFormat: text
    export:
        pingfederate:
            insecureTrustAllTLS: false
            xBypassExternalValidationHeader: false`,
		outputDirectoryReplacement,
		os.Getenv(options.PlatformExportPingoneRegionOption.EnvVar),
		os.Getenv(options.PlatformExportPingoneWorkerClientIDOption.EnvVar),
		os.Getenv(options.PlatformExportPingoneWorkerClientSecretOption.EnvVar),
		os.Getenv(options.PlatformExportPingoneWorkerEnvironmentIDOption.EnvVar),
		os.Getenv(options.PlatformExportPingfederateAdminAPIPathOption.EnvVar),
		os.Getenv(options.PlatformExportPingfederateClientIDOption.EnvVar),
		os.Getenv(options.PlatformExportPingfederateClientSecretOption.EnvVar),
		os.Getenv(options.PlatformExportPingfederateScopesOption.EnvVar),
		os.Getenv(options.PlatformExportPingfederateTokenURLOption.EnvVar),
		os.Getenv(options.PlatformExportPingfederateHTTPSHostOption.EnvVar))
)

func CreateConfigFile(t *testing.T) string {
	t.Helper()

	if configFileContents == "" {
		configFileContents = strings.Replace(defaultConfigFileContents, outputDirectoryReplacement, t.TempDir(), 1)
	}

	configFilepath := t.TempDir() + "/config.yaml"
	if err := os.WriteFile(configFilepath, []byte(configFileContents), 0600); err != nil {
		t.Fatalf("Failed to create config file: %s", err)
	}

	return configFilepath
}

func configureMainViper(t *testing.T) {
	t.Helper()

	// Create and write to a temporary config file
	configFilepath := CreateConfigFile(t)
	// Give main viper instance a file location to write to
	mainViper := profiles.GetMainConfig().ViperInstance()
	mainViper.SetConfigFile(configFilepath)
	if err := mainViper.ReadInConfig(); err != nil {
		t.Fatal(err)
	}

	activePName := profiles.GetMainConfig().ViperInstance().GetString(options.RootActiveProfileOption.ViperKey)

	if err := profiles.GetMainConfig().ChangeActiveProfile(activePName); err != nil {
		t.Fatal(err)
	}
}

func InitVipers(t *testing.T) {
	t.Helper()

	configFileContents = strings.Replace(defaultConfigFileContents, outputDirectoryReplacement, t.TempDir(), 1)

	configureMainViper(t)
}

func InitVipersCustomFile(t *testing.T, fileContents string) {
	t.Helper()

	configFileContents = fileContents
	configureMainViper(t)
}
