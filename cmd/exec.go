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
	"github.com/nikogura/vault-authenticator/pkg/authenticator"
	"github.com/spf13/cobra"
	"log"
)

var execCleanEnv bool

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Exec a shell command with secret loaded into the command's environment.",
	Long: `
Exec a shell command with secret loaded into the command's environment.

Reaches out to Vault and fetches the secret at a path, loading it into the environment, then execev's the command.

Depending on the options used, will either merge the secrets with the default environment (default) or optionally will use a bare environment containing *only* the secrets from Vault.

`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("Exec can't function without a command to execute.\n\n")
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

		// use only the secret named
		if name != "" {
			data = map[string]interface{}{name: data[name]}
		}

		err = authenticator.Exec(args, data, execCleanEnv)
		if err != nil {
			log.Fatalf("Error calling %q: %s", args[0], err)
		}
	},
}

func init() {
	rootCmd.AddCommand(execCmd)

	execCmd.Flags().BoolVarP(&execCleanEnv, "clean", "c", false, "exec command in a clean ENV (secrets only, no parent ENV)")
}
