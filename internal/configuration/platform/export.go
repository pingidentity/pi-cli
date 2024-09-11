package configuration_platform

import (
	"fmt"
	"os"
	"strings"

	"github.com/pingidentity/pingctl/internal/configuration/options"
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/customtypes"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/pflag"
)

func InitPlatformExportOptions() {
	initExportFormatOption()
	initServicesOption()
	initOutputDirectoryOption()
	initOverwriteOption()

	initPingOneWorkerEnvironmentIDOption()
	initPingOneExportEnvironmentIDOption()
	initPingOneWorkerClientIDOption()
	initPingOneWorkerClientSecretOption()
	initPingOneRegionOption()

	initPingFederateHTTPSHostOption()
	initPingFederateAdminAPIPathOption()
	initPingFederateXBypassExternalValidationHeaderOption()
	initPingFederateCACertificatePemFilesOption()
	initPingFederateInsecureTrustAllTLSOption()
	initPingFederateUsernameOption()
	initPingFederatePasswordOption()
	initPingFederateAccessTokenOption()
	initPingFederateClientIDOption()
	initPingFederateClientSecretOption()
	initPingFederateTokenURLOption()
	initPingFederateScopesOption()
}

func initExportFormatOption() {
	cobraParamName := "export-format"
	cobraValue := new(customtypes.ExportFormat)
	defaultValue := customtypes.ExportFormat(connector.ENUMEXPORTFORMAT_HCL)
	envVar := "PINGCTL_EXPORT_FORMAT"

	options.PlatformExportExportFormatOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "e",
			Usage:     fmt.Sprintf("Specifies export format\nAllowed: [%s]. Also configurable via environment variable %s", strings.Join(customtypes.ExportFormatValidValues(), ", "), envVar),
			Value:     cobraValue,
			DefValue:  connector.ENUMEXPORTFORMAT_HCL,
		},
		Type:     options.ENUM_STRING,
		ViperKey: "export.exportFormat",
	}
}

func initServicesOption() {
	cobraParamName := "services"
	cobraValue := new(customtypes.MultiService)
	defaultValue := customtypes.NewMultiService()
	envVar := "PINGCTL_EXPORT_SERVICES"

	options.PlatformExportServiceOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "s",
			Usage:     fmt.Sprintf("Specifies service(s) to export. Accepts comma-separated string to delimit multiple services. Allowed: [%s]. Also configurable via environment variable %s", strings.Join(customtypes.MultiServiceValidValues(), ", "), envVar),
			Value:     cobraValue,
			DefValue:  strings.Join(customtypes.MultiServiceValidValues(), ", "),
		},
		Type:     options.ENUM_MULTI_SERVICE,
		ViperKey: "export.services",
	}
}

func initOutputDirectoryOption() {
	cobraParamName := "output-directory"
	cobraValue := new(customtypes.String)
	defaultValue := getDefaultExportDir()
	envVar := "PINGCTL_EXPORT_OUTPUT_DIRECTORY"

	options.PlatformExportOutputDirectoryOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "d",
			Usage:     fmt.Sprintf("Specifies output directory for export. Also configurable via environment variable %s", envVar),
			Value:     cobraValue,
			DefValue:  "$(pwd)/export",
		},
		Type:     options.ENUM_STRING,
		ViperKey: "export.outputDirectory",
	}
}

func initOverwriteOption() {
	cobraParamName := "overwrite"
	cobraValue := new(customtypes.Bool)
	defaultValue := customtypes.Bool(false)

	options.PlatformExportOverwriteOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "PINGCTL_EXPORT_OVERWRITE",
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "o",
			Usage:     "Overwrite existing generated exports in output directory.",
			Value:     cobraValue,
			DefValue:  "false",
		},
		Type:     options.ENUM_BOOL,
		ViperKey: "export.overwrite",
	}
}

func initPingOneWorkerEnvironmentIDOption() {
	cobraParamName := "pingone-worker-environment-id"
	cobraValue := new(customtypes.UUID)
	defaultValue := customtypes.UUID("")
	envVar := "PINGCTL_PINGONE_WORKER_ENVIRONMENT_ID"

	options.PlatformExportPingoneWorkerEnvironmentIDOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    fmt.Sprintf("The ID of the PingOne environment that contains the worker client used to authenticate.  Also configurable via environment variable %s", envVar),
			Value:    cobraValue,
			DefValue: "",
		},
		Type:     options.ENUM_UUID,
		ViperKey: "export.pingone.worker.environmentID",
	}
}

func initPingOneExportEnvironmentIDOption() {
	cobraParamName := "pingone-export-environment-id"
	cobraValue := new(customtypes.UUID)
	defaultValue := customtypes.UUID("")
	envVar := "PING_CTL_PINGONE_EXPORT_ENVIRONMENT_ID"

	options.PlatformExportPingoneExportEnvironmentIDOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    fmt.Sprintf("The ID of the PingOne environment to export. Also configurable via environment variable %s", envVar),
			Value:    cobraValue,
			DefValue: "",
		},
		Type:     options.ENUM_UUID,
		ViperKey: "export.pingone.export.environmentID",
	}
}

func initPingOneWorkerClientIDOption() {
	cobraParamName := "pingone-worker-client-id"
	cobraValue := new(customtypes.UUID)
	defaultValue := customtypes.UUID("")
	envVar := "PINGCTL_PINGONE_WORKER_CLIENT_ID"

	options.PlatformExportPingoneWorkerClientIDOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    fmt.Sprintf("The ID of the PingOne worker client used to authenticate.  Also configurable via environment variable %s", envVar),
			Value:    cobraValue,
			DefValue: "",
		},
		Type:     options.ENUM_UUID,
		ViperKey: "export.pingone.worker.clientID",
	}
}

func initPingOneWorkerClientSecretOption() {
	cobraParamName := "pingone-worker-client-secret"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")
	envVar := "PINGCTL_PINGONE_WORKER_CLIENT_SECRET"

	options.PlatformExportPingoneWorkerClientSecretOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    fmt.Sprintf("The PingOne worker client secret used to authenticate.  Also configurable via environment variable %s", envVar),
			Value:    cobraValue,
			DefValue: "",
		},
		Type:     options.ENUM_STRING,
		ViperKey: "export.pingone.worker.clientSecret",
	}
}

func initPingOneRegionOption() {
	cobraParamName := "pingone-region"
	cobraValue := new(customtypes.PingOneRegion)
	defaultValue := customtypes.PingOneRegion("")
	envVar := "PINGCTL_PINGONE_REGION"

	options.PlatformExportPingoneRegionOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    fmt.Sprintf("The region of the PingOne service(s). Allowed: %s.  Also configurable via environment variable %s", strings.Join(customtypes.PingOneRegionValidValues(), ", "), envVar),
			Value:    cobraValue,
			DefValue: "",
		},
		Type:     options.ENUM_STRING,
		ViperKey: "export.pingone.region",
	}
}

func initPingFederateHTTPSHostOption() {
	cobraParamName := "pingfederate-https-host"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")
	envVar := "PINGCTL_PINGFEDERATE_HTTPS_HOST"

	options.PlatformExportPingfederateHTTPSHostOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    fmt.Sprintf("The PingFederate HTTPS host used to communicate with PingFederate's API.  Also configurable via environment variable %s", envVar),
			Value:    cobraValue,
			DefValue: "",
		},
		Type:     options.ENUM_STRING,
		ViperKey: "export.pingfederate.httpsHost",
	}
}

func initPingFederateAdminAPIPathOption() {
	cobraParamName := "pingfederate-admin-api-path"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("/pf-admin-api/v1")
	envVar := "PINGCTL_PINGFEDERATE_ADMIN_API_PATH"

	options.PlatformExportPingfederateAdminAPIPathOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    fmt.Sprintf("The PingFederate API URL path used to communicate with PingFederate's API.  Also configurable via environment variable %s", envVar),
			Value:    cobraValue,
			DefValue: "/pf-admin-api/v1",
		},
		Type:     options.ENUM_STRING,
		ViperKey: "export.pingfederate.adminAPIPath",
	}
}

func initPingFederateXBypassExternalValidationHeaderOption() {
	cobraParamName := "pingfederate-x-bypass-external-validation-header"
	cobraValue := new(customtypes.Bool)
	defaultValue := customtypes.Bool(false)
	envVar := "PINGCTL_PINGFEDERATE_X_BYPASS_EXTERNAL_VALIDATION_HEADER"

	options.PlatformExportPingfederateXBypassExternalValidationHeaderOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    fmt.Sprintf("Header value in request for PingFederate. PingFederate's connection tests will be bypassed when set to true.  Also configurable via environment variable %s", envVar),
			Value:    cobraValue,
			DefValue: "false",
		},
		Type:     options.ENUM_BOOL,
		ViperKey: "export.pingfederate.xBypassExternalValidationHeader",
	}
}

func initPingFederateCACertificatePemFilesOption() {
	cobraParamName := "pingfederate-ca-certificate-pem-files"
	cobraValue := new(customtypes.StringSlice)
	defaultValue := customtypes.StringSlice{}
	envVar := "PINGCTL_PINGFEDERATE_CA_CERTIFICATE_PEM_FILES"

	options.PlatformExportPingfederateCACertificatePemFilesOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    fmt.Sprintf("Paths to files containing PEM-encoded certificates to be trusted as root CAs when connecting to the PingFederate server over HTTPS. Accepts comma-separated string to delimit multiple PEM files.  Also configurable via environment variable %s", envVar),
			Value:    cobraValue,
			DefValue: "[]",
		},
		Type:     options.ENUM_STRING_SLICE,
		ViperKey: "export.pingfederate.caCertificatePemFiles",
	}
}

func initPingFederateInsecureTrustAllTLSOption() {
	cobraParamName := "pingfederate-insecure-trust-all-tls"
	cobraValue := new(customtypes.Bool)
	defaultValue := customtypes.Bool(false)
	envVar := "PINGCTL_PINGFEDERATE_INSECURE_TRUST_ALL_TLS"

	options.PlatformExportPingfederateInsecureTrustAllTLSOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    fmt.Sprintf("Set to true to trust any certificate when connecting to the PingFederate server. This is insecure and should not be enabled outside of testing.  Also configurable via environment variable %s", envVar),
			Value:    cobraValue,
			DefValue: "false",
		},
		Type:     options.ENUM_BOOL,
		ViperKey: "export.pingfederate.insecureTrustAllTLS",
	}
}

func initPingFederateUsernameOption() {
	cobraParamName := "pingfederate-username"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")
	envVar := "PINGCTL_PINGFEDERATE_USERNAME"

	options.PlatformExportPingfederateUsernameOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    fmt.Sprintf("The PingFederate username used to authenticate.  Also configurable via environment variable %s", envVar),
			Value:    cobraValue,
			DefValue: "",
		},
		Type:     options.ENUM_STRING,
		ViperKey: "export.pingfederate.basicAuth.username",
	}
}

func initPingFederatePasswordOption() {
	cobraParamName := "pingfederate-password"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")
	envVar := "PINGCTL_PINGFEDERATE_PASSWORD"

	options.PlatformExportPingfederatePasswordOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    fmt.Sprintf("The PingFederate password used to authenticate.  Also configurable via environment variable %s", envVar),
			Value:    cobraValue,
			DefValue: "",
		},
		Type:     options.ENUM_STRING,
		ViperKey: "export.pingfederate.basicAuth.password",
	}
}

func initPingFederateAccessTokenOption() {
	cobraParamName := "pingfederate-access-token"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")
	envVar := "PINGCTL_PINGFEDERATE_ACCESS_TOKEN"

	options.PlatformExportPingfederateAccessTokenOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    fmt.Sprintf("The PingFederate access token used to authenticate.  Also configurable via environment variable %s", envVar),
			Value:    cobraValue,
			DefValue: "",
		},
		Type:     options.ENUM_STRING,
		ViperKey: "export.pingfederate.accessTokenAuth.accessToken",
	}
}

func initPingFederateClientIDOption() {
	cobraParamName := "pingfederate-client-id"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")
	envVar := "PINGCTL_PINGFEDERATE_CLIENT_ID"

	options.PlatformExportPingfederateClientIDOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    fmt.Sprintf("The PingFederate OAuth client ID used to authenticate.  Also configurable via environment variable %s", envVar),
			Value:    cobraValue,
			DefValue: "",
		},
		Type:     options.ENUM_STRING,
		ViperKey: "export.pingfederate.clientCredentialsAuth.clientID",
	}
}

func initPingFederateClientSecretOption() {
	cobraParamName := "pingfederate-client-secret"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")
	envVar := "PINGCTL_PINGFEDERATE_CLIENT_SECRET"

	options.PlatformExportPingfederateClientSecretOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    fmt.Sprintf("The PingFederate OAuth client secret used to authenticate.  Also configurable via environment variable %s", envVar),
			Value:    cobraValue,
			DefValue: "",
		},
		Type:     options.ENUM_STRING,
		ViperKey: "export.pingfederate.clientCredentialsAuth.clientSecret",
	}
}

func initPingFederateTokenURLOption() {
	cobraParamName := "pingfederate-token-url"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")
	envVar := "PINGCTL_PINGFEDERATE_TOKEN_URL"

	options.PlatformExportPingfederateTokenURLOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    fmt.Sprintf("The PingFederate OAuth token URL used to authenticate.  Also configurable via environment variable %s", envVar),
			Value:    cobraValue,
			DefValue: "",
		},
		Type:     options.ENUM_STRING,
		ViperKey: "export.pingfederate.clientCredentialsAuth.tokenURL",
	}
}

func initPingFederateScopesOption() {
	cobraParamName := "pingfederate-scopes"
	cobraValue := new(customtypes.StringSlice)
	defaultValue := customtypes.StringSlice{}
	envVar := "PINGCTL_PINGFEDERATE_SCOPES"

	options.PlatformExportPingfederateScopesOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    fmt.Sprintf("The PingFederate OAuth scopes used to authenticate. Accepts comma-separated string to delimit multiple scopes.  Also configurable via environment variable %s", envVar),
			Value:    cobraValue,
			DefValue: "[]",
		},
		Type:     options.ENUM_STRING_SLICE,
		ViperKey: "export.pingfederate.clientCredentialsAuth.scopes",
	}
}

func getDefaultExportDir() (defaultExportDir *customtypes.String) {
	l := logger.Get()
	pwd, err := os.Getwd()
	if err != nil {
		l.Err(err).Msg("Failed to determine current working directory")
		return nil
	}

	defaultExportDir = new(customtypes.String)

	err = defaultExportDir.Set(fmt.Sprintf("%s/export", pwd))
	if err != nil {
		l.Err(err).Msg("Failed to set default export directory")
		return nil
	}

	return defaultExportDir
}
