package cmd

import (
	"context"
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/dig"
	"go.uber.org/fx"

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
		Use:   "verify",
		Short: "Check if the working directory entered has valid files",
		Long:  verifyDescription,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			app := fx.New(
				fx.NopLogger,
				fx.Provide(files.GetYamlFiles),
				fx.Invoke(func(files []string) {
					for _, file := range files {
						fmt.Fprintln(out, file)
					}
				}),
			)

			ctx := context.Background()
			app.Start(ctx)

			fmt.Println(dig.RootCause(app.Err()))

			return app.Err()
		},
	}

	flags := cmd.Flags()

	flags.StringP("workdir", "w", "specs/", "Specifies the working directory to use")
	viper.BindPFlag("workdir", flags.Lookup("workdir"))

	return cmd
}
