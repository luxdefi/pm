package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ava-labs/apm/apm"
)

func listRepositories() *cobra.Command {
	command := &cobra.Command{
		Use:   "list-repositories",
		Short: "list registered source repositories for avalanche plugins",
	}
	command.RunE = func(_ *cobra.Command, _ []string) error {

		credentials, err := getCredentials()
		if err != nil {
			return err
		}
		apm, err := apm.New(apm.Config{
			Directory: apmDir,
			Auth:      credentials,
		})
		if err != nil {
			return err
		}

		return apm.ListRepositories()
	}

	return command
}
