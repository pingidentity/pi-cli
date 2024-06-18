package output

import (
	"encoding/json"
	"fmt"

	"github.com/fatih/color"
	"github.com/pingidentity/pingctl/internal/customtypes"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/pingidentity/pingctl/internal/viperconfig"
	"github.com/spf13/viper"
)

var (
	boldRed = color.New(color.FgRed).Add(color.Bold).SprintfFunc()
	cyan    = color.New(color.FgCyan).SprintfFunc()
	green   = color.New(color.FgGreen).SprintfFunc()
	red     = color.New(color.FgRed).SprintfFunc()
	white   = color.New(color.FgWhite).SprintfFunc()
	yellow  = color.New(color.FgYellow).SprintfFunc()
)

type CommandOutputResult string

type CommandOutput struct {
	Fields       map[string]interface{}
	Message      string
	ErrorMessage string
	FatalMessage string
	Result       CommandOutputResult
}

const (
	ENUMCOMMANDOUTPUTRESULT_NIL           CommandOutputResult = ""
	ENUMCOMMANDOUTPUTRESULT_SUCCESS       CommandOutputResult = "Success"
	ENUMCOMMANDOUTPUTRESULT_NOACTION_OK   CommandOutputResult = "No Action (OK)"
	ENUMCOMMANDOUTPUTRESULT_NOACTION_WARN CommandOutputResult = "No Action (Warning)"
	ENUMCOMMANDOUTPUTRESULT_FAILURE       CommandOutputResult = "Failure"
)

func Format(output CommandOutput) {
	colorizeOutput := viper.GetBool(viperconfig.ConfigOptions[viperconfig.RootColorParamName].ViperConfigKey)

	if !colorizeOutput {
		color.NoColor = true
	}

	// Get the output format from viper configuration
	// If output format is loaded from file, it is of type string
	// if output is loaded from parameter or "config set" it is of type common.OutputFormat
	outputFormat := viper.Get(viperconfig.ConfigOptions[viperconfig.RootOutputParamName].ViperConfigKey)
	var outputFormatString string
	switch format := outputFormat.(type) {
	case customtypes.OutputFormat:
		outputFormatString = format.String()
	case string:
		outputFormatString = format
	}

	switch outputFormatString {
	case customtypes.ENUM_OUTPUT_FORMAT_TEXT:
		formatText(output)
	case customtypes.ENUM_OUTPUT_FORMAT_JSON:
		formatJson(output)
	default:
		formatText(CommandOutput{
			Message: fmt.Sprintf("Output format %q is not recognized. Defaulting to \"text\" output", outputFormat),
			Result:  ENUMCOMMANDOUTPUTRESULT_NOACTION_WARN,
		})
		formatText(output)
	}
}

func formatText(output CommandOutput) {
	l := logger.Get()

	var resultFormat string
	var resultColor func(format string, a ...interface{}) string

	// Determine message color and format based on status
	switch output.Result {
	case ENUMCOMMANDOUTPUTRESULT_SUCCESS:
		resultFormat = "%s - %s"
		resultColor = green
	case ENUMCOMMANDOUTPUTRESULT_NOACTION_OK:
		resultFormat = "%s - %s"
		resultColor = green
	case ENUMCOMMANDOUTPUTRESULT_NOACTION_WARN:
		resultFormat = "%s - %s"
		resultColor = yellow
	case ENUMCOMMANDOUTPUTRESULT_FAILURE:
		resultFormat = "%s - %s"
		resultColor = red
	case ENUMCOMMANDOUTPUTRESULT_NIL:
		resultFormat = "%s%s"
		resultColor = white
	default:
		resultFormat = "%s%s"
		resultColor = white
	}

	// Supply the user a formatted message and a result status if any.
	fmt.Println(resultColor(resultFormat, output.Message, output.Result))
	l.Info().Msgf(resultColor(resultFormat, output.Message, output.Result))

	// Output and log any additional key/value pairs supplied to the user.
	if output.Fields != nil {
		fmt.Println(cyan("Additional Information:"))
		for k, v := range output.Fields {
			fmt.Println(cyan("%s: %s", k, v))
			l.Info().Msgf("%s: %s", k, v)
		}
	}

	// Inform the user of an error and log the error
	if output.ErrorMessage != "" {
		fmt.Println(red("Error: %s", output.ErrorMessage))
		l.Error().Msgf(output.ErrorMessage)
	}

	// Inform the user of a fatal error and log the fatal error. This exits the program.
	if output.FatalMessage != "" {
		fmt.Println(boldRed("Fatal: %s", output.FatalMessage))
		l.Fatal().Msgf(output.FatalMessage)
	}

}

func formatJson(output CommandOutput) {
	l := logger.Get()

	// Convert the CommandOutput struct to JSON
	jsonOut, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		l.Error().Err(err).Msgf("Failed to serialize output as JSON")
	}

	// Output the JSON as uncolored string
	fmt.Println(string(jsonOut))

	switch output.Result {
	case ENUMCOMMANDOUTPUTRESULT_NOACTION_WARN:
		l.Warn().Msgf(string(jsonOut))
	case ENUMCOMMANDOUTPUTRESULT_FAILURE:
		// Log the error if exists
		if output.ErrorMessage != "" {
			l.Error().Msgf(output.ErrorMessage)
		}

		// Log the fatal error if exists. This exits the program.
		if output.FatalMessage != "" {
			l.Fatal().Msgf(output.FatalMessage)
		}
	default: //ENUMCOMMANDOUTPUTRESULT_SUCCESS, ENUMCOMMANDOUTPUTRESULT_NIL, ENUMCOMMANDOUTPUTRESULT_NOACTION_OK
		l.Info().Msgf(string(jsonOut))
	}

}
