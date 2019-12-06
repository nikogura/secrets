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

// moveCmd represents the move command
var moveCmd = &cobra.Command{
	Use:   "move",
	Short: "Move a secret from one path to another.",
	Long: `

Move a secret from one path to another.

Copies a secret from one path to another, and then deletes the original secret.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if team == "" {
			log.Fatal("Team name is required for move.  (-t <name>)\n")
		}

		if env == "" {
			log.Fatal("Environment name is required for move.  (-t <name>)\n")
		}

		if toTeam == "" {
			toTeam = team
		}

		if toName == "" {
			log.Fatal("--toname is required for move.  (-toname <name>)\n")
		}

		if toEnv == "" {
			log.Fatal("--toenv is required for move.  (-toenv <name>)\n")
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

		err = authenticator.MoveSecret(client, fromPath, toPath)
		if err != nil {
			log.Fatalf("Error moving secret from team %q name %q env %q to team %q name %q env %q: %s", team, name, env, toTeam, toName, toEnv, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(moveCmd)

	moveCmd.Flags().StringVarP(&toTeam, "toteam", "", "", "Team to copy secret to")
	moveCmd.Flags().StringVarP(&toName, "toname", "", "", "Name to copy secret to")
	moveCmd.Flags().StringVarP(&toEnv, "toenv", "", "", "Environment to copy secret to")
}
