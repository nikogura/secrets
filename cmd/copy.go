// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	"github.com/scribd/vault-authenticator/pkg/authenticator"
	"github.com/spf13/cobra"
	"log"
)

var toTeam string
var toName string
var toEnv string

// copyCmd represents the copy command
var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copy a secret from one path to another.",
	Long: `
Copy a secret from one path to another.

Does not delete original path.

Example: Copy
`,
	Run: func(cmd *cobra.Command, args []string) {
		if team == "" {
			log.Fatal("Team name is required for edit.  (-t <name>)\n")
		}

		if env == "" {
			log.Fatal("Environment name is required for edit.  (-t <name>)\n")
		}

		if toTeam == "" {
			toTeam = team
		}

		if toName == "" {
			log.Fatal("--toname is required for copy.  (-toname <name>)\n")
		}

		if toEnv == "" {
			log.Fatal("--toenv is required for copy.  (-toenv <name>)\n")
		}

		fromPath := fmt.Sprintf("%s/data/%s/%s", team, name, env)
		toPath := fmt.Sprintf("%s/data/%s/%s", toTeam, toName, toEnv)

		auth.SetVerbose(verbose)
		auth.SetPrompt(!silent)
		auth.SetIdentifier(identifier)
		auth.SetRole(role)

		client, err := auth.Auth()
		if err != nil {
			if err.Error() == authenticator.VAULT_AUTH_FAIL {
				log.Fatalf("failed to get vault token: %s cannot proceed.", err)
			}
		}

		err = authenticator.CopySecret(client, fromPath, toPath)
		if err != nil {
			log.Fatalf("Error copying secret from team %q name %q env %q to team %q name %q env %q: %s", team, name, env, toTeam, toName, toEnv, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(copyCmd)

	copyCmd.Flags().StringVarP(&toTeam, "toteam", "", "", "Team to copy secret to")
	copyCmd.Flags().StringVarP(&toName, "toname", "", "", "Name to copy secret to")
	copyCmd.Flags().StringVarP(&toEnv, "toenv", "", "", "Environment to copy secret to")
}
