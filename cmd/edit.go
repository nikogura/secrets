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
	"bufio"
	"fmt"
	"github.com/nikogura/vault-authenticator/pkg/authenticator"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit the secret at a particular path in Vault",
	Long: `
Edit the secret at a particular path in Vault.

Edit opens up your favorite editor (based on the value of the ENV var EDITOR), and lets you edit the secrets for a given role.

If no secrets exist, it will create a template from where you can start.

Secrets are stored encrypted in Hashicorp Vault at a particular path.   Generally secret paths start with 'secret/' .

`,
	Run: func(cmd *cobra.Command, args []string) {
		if team == "" {
			log.Fatal("Team name is required for edit.  (-t <name>)\n")
		}

		if env == "" {
			log.Fatal("Environment name is required for edit.  (-e <name>)\n")
		}

		if name == "" {
			if len(args) > 0 {
				name = args[0]

			}

			if name == "" {
				reader := bufio.NewReader(os.Stdin)

				fmt.Printf("Which secret would you like to edit?  ")

				input, err := reader.ReadString('\n')
				if err != nil {
					fmt.Printf("Error reading desired path: %s\n", err)
					os.Exit(1)
				}

				if input != "\n" {
					trimmed := strings.Trim(input, "\n")

					name = trimmed
				} else {
					log.Fatal("Cannot edit without a secret name.\n\n")
				}
			}
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

		err = authenticator.EditSecret(client, path)
		if err != nil {
			log.Fatalf("Error editing secret: %s", err)
		}

		fmt.Printf("Update complete.\n")
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

}
