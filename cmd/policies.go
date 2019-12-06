package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

// policiesCmd represents the fetch command
var policiesCmd = &cobra.Command{
	Use:   "policies",
	Short: "Authenticate to Vault, and show granted policies",
	Long: `
Authenticate to Vault. and show granted policies.

Mostly used for debugging, this command auths to Vault and gets a token, then looks up what that token allows you to see.
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

		token, err := client.Auth().Token().LookupSelf()
		if err != nil { // will err if it's not
			log.Fatalf("failed looking up my own token: %s", err)
		}

		if token == nil {
			log.Fatalf("Failed to look up token.")
		}

		fmt.Printf("\nPolicies attached to this token:\n\n")
		policies, ok := token.Data["policies"].([]interface{})
		if ok {
			for _, policy := range policies {
				fmt.Printf("  %s\n", policy)
			}
		}

		fmt.Printf("\n\n")
	},
}

func init() {
	rootCmd.AddCommand(policiesCmd)
}
