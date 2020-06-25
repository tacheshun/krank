// Application which greets you.
package main

import (
	"flag"
	"github.com/spf13/cobra"
	scanscli "github.com/tacheshun/krank/internal"
	"github.com/tacheshun/krank/internal/cli"
	"github.com/tacheshun/krank/internal/storage"
)

func main() {
	var repo scanscli.ScanRepo
	flag.Parse()

	repo = storage.NewScanRepository()

	rootCmd := &cobra.Command{Use: "scans-cli"}
	rootCmd.AddCommand(cli.InitScansCommand(repo))
	_ = rootCmd.Execute()
}
