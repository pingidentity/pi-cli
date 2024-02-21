package platform_test

import (
	"bytes"
	"testing"

	"github.com/pingidentity/pingctl/cmd"
)

// Test Platform Export Command Executes without issue
func TestPlatformExportCmd_Execute(t *testing.T) {
	// Create the command
	rootCmd := cmd.NewRootCommand()

	// Redirect stdout to a buffer to capture the output
	var stdout bytes.Buffer
	rootCmd.SetOut(&stdout)
	rootCmd.SetErr(&stdout)

	rootCmd.SetArgs([]string{"platform", "export"})

	// Execute the command
	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("Export Command failed. Make sure to have PingOne env variables set if test is failing.\nErr: %q, Captured StdOut: %q", err, stdout.String())
	}
}
