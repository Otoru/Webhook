package cmd

import (
	"context"
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/dig"
	"go.uber.org/fx"
	"go.uber.org/multierr"

	"github.com/otoru/webhook/pkg/files"
)

const verifyDescription = `
Check if the working directory entered has valid files.

The command will go through each of the files and perform a recursive validation of the used file structure.

It is important to note that content validation itself is not performed.
So even with validation it is still possible to receive errors at runtime,
such as two listeners disputing the same connection port.
`

func CreateVerifyCommand(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "verify",
		Short:        "Check if the working directory entered has valid files",
		Long:         verifyDescription,
		Args:         cobra.NoArgs,
		SilenceUsage: false,
		RunE: func(cmd *cobra.Command, args []string) error {
			app := fx.New(
				fx.NopLogger,
				fx.Provide(files.GetYamlFiles),
				fx.Provide(files.GetDocuments),
				fx.Provide(files.GetSpecifications),
				fx.Invoke(func(specs []*files.Specification) error {
					var result error

					for _, item := range specs {
						result = multierr.Append(result, item.Validate())
					}

					if result == nil {
						fmt.Fprintln(out, "✨ Everything is ok!")
					}

					return result
				}),
			)

			ctx := context.Background()
			err := app.Start(ctx)

			return dig.RootCause(err)
		},
	}

	flags := cmd.Flags()

	flags.StringP("workdir", "w", "specs/", "Specifies the working directory to use")
	flags.Bool("short", false, "Print only the version number")
	viper.BindPFlag("workdir", flags.Lookup("workdir"))
	viper.BindPFlag("short", flags.Lookup("short"))

	return cmd
}
