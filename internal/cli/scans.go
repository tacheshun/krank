//Package cli ...
package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"github.com/tacheshun/krank/internal/fetching"

	"github.com/spf13/cobra"
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

		_, err := service.FetchScans()
		if err != nil {
			log.Fatal(err)
		}
		var datas map[string]string
		datas, _, err = service.RunBasicScan()
		if err != nil {
			log.Fatal(err)
		}
		jsonString, err := json.Marshal(datas)
		if err != nil {
			panic(err)
		}

		// HTTP Request to RMM Callback URL here
		req, err := http.NewRequest("POST", "http://localhost:8000/", bytes.NewBuffer(jsonString))
		if err != nil {
			panic(err)
		}

		req.Header.Set("X-Custom-Header", "fromKrank")
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("response Body:", string(body))
	}
}
