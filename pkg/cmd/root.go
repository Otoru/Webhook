package cmd

import (
	"io"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const WebhookDescription = `
Command line tool for simple implementation of webhooks

This tool aims to make a software developer's life easier.
Through it you should be able to perform simple validations such as:

- Receiving a webhook, printing its content on the terminal and sending a response.
- Creating more complex logic (such as validating the request body or decision tree).

All configuration is performed through the template files present in the workdir.
`

func CreateRootCommand(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "webhook",
		Short: "Command line tool for simple implementation of webhooks",
		Long:  WebhookDescription,
	}

	flags := cmd.Flags()
	flags.StringP("log-level", "l", "warn", "Set the application log-level")
	viper.BindPFlag("log-level", flags.Lookup("log-level"))

	cmd.AddCommand(
		CreateVerifyCommand(out),
		CreateVersionCommand(out),
	)

	return cmd
}
