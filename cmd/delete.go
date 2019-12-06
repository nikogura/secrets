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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a secret",
	Long: `
Delete a secret.

Deletes a secret at the path given.  What more can you say?
`,
	Run: func(cmd *cobra.Command, args []string) {
		if team == "" {
			log.Fatal("Team name is required for delete.  (-t <name>)\n")
		}

		if env == "" {
			log.Fatal("Environment name is required for delete.  (-e <name>)\n")
		}

		path := fmt.Sprintf("%s/data/%s/%s", team, name, env)

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

		err = authenticator.DeleteSecrets(client, path)
		if err != nil {
			log.Fatalf("Error deleting secret %q for team %q in environment %q: %s", name, team, env, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
