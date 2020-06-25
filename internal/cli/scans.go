//Package cli.
package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	scanscli "github.com/tacheshun/krank/internal"
	"strconv"
)

// CobraFn function definion of run cobra command.
type CobraFn func(cmd *cobra.Command, args []string)

const idFlag = "id"

//InitScansCommand public method .
func InitScansCommand(repository scanscli.ScanRepo) *cobra.Command {
	scanCmd := &cobra.Command{
		Use:   "scans",
		Short: "Print data about scans",
		Run:   runScansFn(repository),
	}

	scanCmd.Flags().StringP(idFlag, "i", "", "id of the beer")

	return scanCmd
}

func runScansFn(repository scanscli.ScanRepo) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		scans, _ := repository.GetScans()

		id, _ := cmd.Flags().GetString(idFlag)

		if id != "" {
			i, _ := strconv.Atoi(id)
			for _, scan := range scans {
				if scan.ScanID == i {
					fmt.Println(scan)
					return
				}
			}
		} else {
			fmt.Println(scans)
		}
	}
}