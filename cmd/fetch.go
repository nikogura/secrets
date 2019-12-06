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
	"encoding/json"
	"fmt"
	"github.com/scribd/vault-authenticator/pkg/authenticator"
	"github.com/spf13/cobra"
	"log"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch a single secret from Vault",
	Long: `
Fetch a secret from Vault.

Reaches out to Vault and fetches secrets.  If -p <path> is provided, fetches just that path (provided you can read it).

Otherwise it will fetch all secrets for the role provided and return them.  Roles are namespaced to prevent collisions.

Roles are requested with -r <roleNamespace-roleName>.  Either a role or a path is required.

Returns it as a string suitable for use in bash scripts e.g. FOO=$(dbt secrets fetch -r $namespace-$role), or as JSON if '-f json' is provided.

All secrets are multi-valued by default.  If you want only a single value, set -k <key name> to limit the output.

If you desire JUST the raw value of a single key, use '-k <key name> -f raw ...'.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			role = args[0]
		}

		auth.SetVerbose(verbose)
		auth.SetPrompt(!silent)
		auth.SetIdentifier(identifier)
		auth.SetRole(role)

		client, err := auth.Auth()
		if err != nil {
			log.Fatalf("failed to auth to Vault: %s.", err)
		}

		data, err := authenticator.SecretsForRole(client, role, env, verbose)
		if err != nil {
			log.Fatalf("Error getting secret for role %s: %s", role, err)
		}

		var fetchedSecretOutput string

		// use only the secret named
		if key != "" {
			data = map[string]interface{}{key: data[key]}
		}

		switch format {
		case "json":
			jsonBytes, err := json.Marshal(data)
			if err != nil {
				log.Fatalf("failed to marshal secret into json: %s", err)
			}
			fmt.Print(string(jsonBytes))

		case "raw":
			if key != "" {
				fmt.Printf("%s", data[key])
			}

		default:
			for k, v := range data {
				fetchedSecretOutput += fmt.Sprintf("%s: %s\n", k, v)
			}

			fmt.Printf(fetchedSecretOutput)
		}
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}
