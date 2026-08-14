package cmd

import (
	"github.com/spf13/cobra"

	"github.com/dynatrace-oss/dtctl/pkg/output"
	"github.com/dynatrace-oss/dtctl/pkg/resources/platform"
)

// describeEnvironmentCmd shows detailed environment information
var describeEnvironmentCmd = &cobra.Command{
	Use:     "environment",
	Aliases: []string{"env"},
	Short:   "Show details of the current environment",
	Long: `Show detailed information about the current Dynatrace environment.

Examples:
  # Describe the current environment
  dtctl describe environment
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		_, c, printer, err := Setup()
		if err != nil {
			return err
		}

		h := platform.NewHandler(c)
		info, err := h.GetEnvironment()
		if err != nil {
			return err
		}

		if outputFormat == "table" {
			const w = 12
			output.DescribeKV("ID:", w, "%s", info.EnvironmentID)
			output.DescribeKV("Type:", w, "%s", info.Type)
			output.DescribeKV("State:", w, "%s", info.State)
			output.DescribeKV("Created:", w, "%s", info.CreateTime.Format("2006-01-02"))
			output.DescribeKV("Block Time:", w, "%s", info.BlockTime.Format("2006-01-02"))
			return nil
		}

		enrichAgent(printer, "describe", "environment")
		return printer.Print(info)
	},
}

// describeLicenseCmd shows detailed license information
var describeLicenseCmd = &cobra.Command{
	Use:   "license",
	Short: "Show details of the environment license",
	Long: `Show detailed license information for the current Dynatrace environment.

Examples:
  # Describe the environment license
  dtctl describe license
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		_, c, printer, err := Setup()
		if err != nil {
			return err
		}

		h := platform.NewHandler(c)
		lic, err := h.GetLicense()
		if err != nil {
			return err
		}

		if outputFormat == "table" {
			const w = 24
			output.DescribeKV("Trial:", w, "%v", lic.Trial)
			output.DescribeKV("Platform Subscription:", w, "%v", lic.PlatformSubscription)
			return nil
		}

		enrichAgent(printer, "describe", "license")
		return printer.Print(lic)
	},
}
