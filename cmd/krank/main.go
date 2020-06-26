// Application which will scan for vulnerabilities on a device.
package main

import (
	"github.com/spf13/cobra"

	scanscli "github.com/tacheshun/krank/internal"
	"github.com/tacheshun/krank/internal/cli"
	"github.com/tacheshun/krank/internal/fetching"
	"github.com/tacheshun/krank/internal/storage"
)

func main() {
	var repo scanscli.ScanRepo

	repo = storage.NewScanRepository()
	fetchingService := fetching.NewService(repo)

	rootCmd := &cobra.Command{Use: "scans-cli"}
	rootCmd.AddCommand(cli.InitScansCommand(fetchingService))
	_ = rootCmd.Execute()
}
