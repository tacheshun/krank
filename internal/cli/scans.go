//Package cli ...
package cli

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/tacheshun/krank/internal/fetching"
)

// CobraFn function definion of run cobra command.
type CobraFn func(cmd *cobra.Command, args []string)

const idFlag = "id"

//InitScansCommand public method .
func InitScansCommand(service fetching.Service) *cobra.Command {
	scanCmd := &cobra.Command{
		Use:   "scans",
		Short: "Print data about scans",
		Run:   runScansFn(service),
	}

	scanCmd.Flags().StringP(idFlag, "i", "", "id of the scan")

	return scanCmd
}

func runScansFn(service fetching.Service) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString(idFlag)
		if id != "" {
			i, _ := strconv.Atoi(id)
			scan, err := service.FetchByID(i)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(scan)
			return
		}

		scans, err := service.FetchScans()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(scans)
		service.RunBasicScan()
	}
}