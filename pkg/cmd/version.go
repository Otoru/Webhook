package cmd

import (
	"fmt"
	"io"

	"github.com/otoru/webhook/internal/version"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const versionDescription = `
Show the version for webhook.

This will print a representation the version of webhook.
The output will look something like this:

version.Info{Version:"v1.0.0", Go:"go1.13.10"}

- Version is the semantic version of the release.
- Go is the version of Go that was used to compile webhook.
`

func CreateVersionCommand(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the client version information",
		Long:  versionDescription,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			short := viper.GetBool("short")
			fmt.Fprintln(out, version.Get(short))
		},
	}

	flags := cmd.Flags()

	flags.Bool("short", false, "Print only the version number")
	viper.BindPFlag("short", flags.Lookup("short"))

	return cmd
}
