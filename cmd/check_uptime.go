package cmd

import (
	"github.com/inexio/thola/internal/request"
	"github.com/spf13/cobra"
)

func init() {
	addDeviceFlags(checkUptimeCMD)
	checkCMD.AddCommand(checkUptimeCMD)

	checkUptimeCMD.Flags().Float64("warning", 0, "warning threshold for uptime in seconds (warning if below)")
	checkUptimeCMD.Flags().Float64("critical", 0, "critical threshold for uptime in seconds (critical if below)")
}

var checkUptimeCMD = &cobra.Command{
	Use:   "uptime",
	Short: "Check the uptime of a device",
	Long: "Checks the uptime of a device in seconds.\n\n" +
		"The usage will be printed as performance data.",
	Run: func(cmd *cobra.Command, args []string) {
		r := request.CheckUptimeRequest{
			CheckDeviceRequest: getCheckDeviceRequest(args[0]),
			UptimeThreshold:    generateCheckThresholds(cmd, "warning", "", "critical", "", false),
		}
		handleRequest(&r)
	},
}
