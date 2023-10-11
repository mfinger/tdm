package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"net"
	"os"
	"tdm/TasDevMgr"
)

func BackupCommand(cmd *cobra.Command, args []string) {
	device := TasDevMgr.Device{Address: net.ParseIP(args[0])}
	if err := device.Backup(args[1]); err != nil {
		fmt.Fprintf(os.Stderr, "[%s] Backup failed: %s\n", args[0], err.Error())
	}
	fmt.Printf("[%s] Configuration backed up to: %s\n", args[0], args[1])
}

func ProbeCommand(cmd *cobra.Command, args []string) {
	device := TasDevMgr.Device{Address: net.ParseIP(args[0])}
	if err := device.Probe(); err != nil {
		fmt.Fprintf(os.Stderr, "[%s] Probe failed: %s\n", args[0], err.Error())
	}
	device.Print()
}
