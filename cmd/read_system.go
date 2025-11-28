package cmd

import (
	"github.com/inexio/thola/internal/request"
	"github.com/spf13/cobra"
)

func init() {
	addDeviceFlags(readSystemCMD)
	readCMD.AddCommand(readSystemCMD)
}

var readSystemCMD = &cobra.Command{
	Use:   "system",
	Short: "Read out system information of a device",
	Long:  "Read out system information of a device like name, description, contact, location and uptime.",
	Run: func(cmd *cobra.Command, args []string) {
		request := request.ReadSystemRequest{
			ReadRequest: getReadRequest(args[0]),
		}
		handleRequest(&request)
	},
}
