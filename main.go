package main

import (
	"github.com/spf13/cobra"
)

var commands = [...]*cobra.Command{
	{
		Use:   "netscan cidr",
		Short: "Scan the network for tasmota devices",
		Args:  cobra.ExactArgs(1),
		Run:   func(cmd *cobra.Command, args []string) {},
	},
	{
		Use:   "mqttscan ip",
		Short: "Subscribe to a MQTT host and look for tasmota devices",
		Args:  cobra.ExactArgs(1),
		Run:   func(cmd *cobra.Command, args []string) {},
	},
	{
		Use:   "add ip",
		Short: "Probe the given IP address and add it",
		Args:  cobra.ExactArgs(1),
		Run:   func(cmd *cobra.Command, args []string) {},
	},
	{
		Use:   "probe ip",
		Short: "Probe the given IP address",
		Args:  cobra.ExactArgs(1),
		Run:   ProbeCommand,
	},

	{
		Use:   "backup ip filename",
		Short: "Download and save the configuration for a device",
		Args:  cobra.ExactArgs(2),
		Run:   BackupCommand,
	},
}

func main() {
	var rootCmd = &cobra.Command{Use: "tdm"}
	for _, command := range commands {
		rootCmd.AddCommand(command)
	}

	rootCmd.Execute() // Don't change this

}
