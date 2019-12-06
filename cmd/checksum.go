package cmd

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/scribd/vault-authenticator/pkg/authenticator"
	"github.com/spf13/cobra"
	"log"
)

// checksumCmd represents the fetch command
var checksumCmd = &cobra.Command{
	Use:   "checksum",
	Short: "Produce a sha256 checksum for secrets in Vault.",
	Long: `
Produce a sha256 checksum for secrets in Vault.

Reaches out to Vault, fetches secrets, and calculates the checksum over them.  Returns the hex value for the sum.

Used to detect if/when the secrets you're consuming have changed.
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

		// use only the secret named
		if name != "" {
			data = map[string]interface{}{name: data[name]}
		}

		jBytes, err := json.Marshal(data)
		if err != nil {
			log.Fatalf("Failed to marshal secret data into json")
		}

		sum := sha256.Sum256(jBytes)

		fmt.Printf("%x", sum)
	},
}

func init() {
	rootCmd.AddCommand(checksumCmd)
}
