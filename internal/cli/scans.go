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

	"github.com/spf13/cobra"

	"github.com/tacheshun/krank/internal/fetching"
	"github.com/tacheshun/krank/internal/storage"
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
		rmmData, err := service.FetchScans()
		if err != nil {
			log.Fatal(err)
		}
		var datas map[string]string
		for i, _ := range rmmData {
			datas, _, err = service.RunBasicScan(strconv.Itoa(rmmData[i].JobID))
		}
		if err != nil {
			log.Fatal(err)
		}

		jsonString, err := json.Marshal(datas)
		if err != nil {
			panic(err)
		}
		// HTTP Request to RMM Callback URL here
		req, err := http.NewRequest("POST", storage.NmapEndpointAck, bytes.NewBuffer(jsonString))
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
		fmt.Println("request Body:", bytes.NewBuffer(jsonString))
	}
}
