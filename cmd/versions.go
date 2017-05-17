// Copyright © 2017 Compose, an IBM Company
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// versionsCmd represents the versions command
var versionsCmd = &cobra.Command{
	Use:   "versions [deployment id]",
	Short: "Show versions for deployment database",
	Long:  `Shows all available upgrade versions for the database installed within a deployment`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("Need a deployment id to examine")
		}
		c := getComposeAPI()
		if outputRaw {
			text, errs := c.GetVersionsForDeploymentJSON(args[0])
			bailOnErrs(errs)
			fmt.Println(text)
		} else {
			versions, errs := c.GetVersionsForDeployment(args[0])
			bailOnErrs(errs)
			if !outputJSON {
				for _, v := range *versions {
					printVersionTransitions(v)
					fmt.Println()
				}
			} else {
				printAsJSON(*versions)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(versionsCmd)
}
