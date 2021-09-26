/*
Copyright © 2021 techytoes

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

// getIPCmd represents the getIP command
var getIPCmd = &cobra.Command{
	Use:   "ip",
	Short: "Fetches IPv4 address of machine and also copies the same",
	Long: `
		Using this command fetches IPv4 address of the machine and also copies the same to clipboard.
		Example: > alfred ip
		Here you go -> 103.246.40.30
		Note: this is also copied to the clipboard.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		ip := getIPAddress()
		clipboard.WriteAll(ip)
		fmt.Println("Here you go ->", ip)
		fmt.Println("Note: this is also copied to the clipboard.")
	},
}

func init() {
	rootCmd.AddCommand(getIPCmd)
}

type Ip struct {
	IPAddress string `json:"query"`
	Status    string `json:"status"`
}

// function to fetch the IP address of the machine
func getIPAddress() string {
	url := "http://ip-api.com/json"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		panic(err)
	}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	ip := Ip{}
	if err := json.Unmarshal(body, &ip); err != nil {
		panic(err)
	}

	return ip.IPAddress
}
