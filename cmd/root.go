// Copyright © 2018 Nik Ogura <nik@scribd.com>
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
	"github.com/nikogura/vault-authenticator/pkg/authenticator"
	"github.com/spf13/cobra"
	"os"
)

var verbose bool
var silent bool
var identifier string
var format string
var role string
var env string
var name string
var team string
var key string
var auth *authenticator.Authenticator

var rootCmd = &cobra.Command{
	Use:   "secrets",
	Short: "Client for Managed Secrets",
	Long:  "Client for Managed Secrets",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
	rootCmd.PersistentFlags().BoolVarP(&silent, "silent", "s", false, "Silent mode (disables auth prompts)")
	rootCmd.PersistentFlags().StringVarP(&identifier, "identifier", "i", "", "identifier (required for k8s usage).")
	rootCmd.PersistentFlags().StringVarP(&key, "key", "k", "", "key name.  -k <key> returns only the value of <key>.  (secrets are all multi valued)")
	rootCmd.PersistentFlags().StringVarP(&format, "format", "f", "", "format in which to output secrets.  Defaults to shell env. Other option is 'json'")
	rootCmd.PersistentFlags().StringVarP(&role, "role", "r", "", "role to auth to and pull secrets from.")
	rootCmd.PersistentFlags().StringVarP(&env, "environment", "e", "", "environment for secrets to pull. (generally admin only- in most cases your environment is determined by where you call `secrets`)")
	rootCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "secret name")
	rootCmd.PersistentFlags().StringVarP(&team, "team", "t", "", "team name")

	auth = authenticator.NewAuthenticator()
	auth.SetAddress("https://vault.example.com")
	auth.SetAuthMethods([]string{
		"iam",
		"k8s",
		"tls",
		"ldap",
	})
	auth.SetTlsClientCrtPath("/path/to/host.crt")
	auth.SetTlsClientKeyPath("/path/to/host.key")
}
